import * as chan from "@nodeguy/channel";

const msgChannelOpen = 100;
const msgChannelOpenConfirm = 101;
const msgChannelOpenFailure = 102;
const msgChannelWindowAdjust = 103;
const msgChannelData = 104;
const msgChannelEOF = 105;
const msgChannelClose = 106;

const minPacketLength = 9;
const channelMaxPacket = 1 << 15;
const channelWindowSize = 64 * channelMaxPacket;

interface channelOpenMsg {
	peersID:       number;
	peersWindow:   number;
	maxPacketSize: number;
}

interface channelOpenConfirmMsg {
	peersID:       number;
	myID:          number;
	myWindow:      number;
	maxPacketSize: number;
}

interface channelOpenFailureMsg {
	peersID: number;
}

interface channelWindowAdjustMsg {
	peersID:         number;
	additionalBytes: number;
}

interface channelDataMsg {
	peersID: number;
	length:  number;
	rest:    Uint8Array;
}

interface channelEOFMsg {
	peersID: number;
}

interface channelCloseMsg {
	peersID: number;
}

interface IConn {
	recv(): Promise<Uint8Array>;
	send(buffer: ArrayBuffer): Promise<number>;
	close(): Promise<void>;
	closeWrite(): Promise<void>;
}

interface ISession {
	open(): Promise<IChannel>;
	accept(): Promise<IChannel>;
    close(): Promise<void>;
    //wait(): Promise<void>;
}

interface IChannel extends IConn {
	ident(): number
}

export class Session implements ISession {
	conn: IConn;
	channels: Array<Channel>;
	incoming: chan;

	constructor(conn: IConn) {
		this.conn = conn;
		this.channels = [];
		this.incoming = chan();
		this.loop();
	}

	async readPacket(): Promise<ArrayBuffer> {
		var sizes = {
			msgChannelOpen:         12,
			msgChannelOpenConfirm:  16,
			msgChannelOpenFailure:  4,
			msgChannelWindowAdjust: 8,
			msgChannelData:         8,
			msgChannelEOF:          4,
			msgChannelClose:        4,
		}
		var packet = await this.conn.recv();
		// if (packet[0] == msgChannelData) {
		// 	dataSize := binary.BigEndian.Uint32(rest[4:8])
		// 	data := make([]byte, dataSize)
		// 	_, err := c.Read(data)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	packet = append(packet, data...)
		// }
		return Promise.resolve(packet.buffer);
	}

	async handleChannelOpen(packet: ArrayBuffer) {
		var msg: channelOpenMsg = decode(packet);
		if (msg.maxPacketSize < minPacketLength || msg.maxPacketSize > 1<<31) {
			await this.conn.send(encode(msgChannelOpenFailure, {
				peersID: msg.peersID
			}));
		}
		var c = this.newChannel();
		c.remoteId = msg.peersID;
		c.maxRemotePayload = msg.maxPacketSize;
		c.remoteWin = msg.peersWindow;
		c.maxIncomingPayload = channelMaxPacket;
		await this.incoming.push(c);
		await this.conn.send(encode(msgChannelOpenConfirm, {
			peersID: c.remoteId,
			myID: c.localId,
			myWindow: c.myWindow,
			maxPacketSize: c.maxIncomingPayload
		}));
	}

	async open(): Promise<IChannel> {
		var ch = this.newChannel();
		ch.maxIncomingPayload = channelMaxPacket;
		await this.conn.send(encode(msgChannelOpen, {
			peersWindow: ch.myWindow,
			maxPacketSize: ch.maxIncomingPayload,
			peersID: ch.localId
		}));
		if (await ch.ready.shift()) {
			return Promise.resolve(ch);
		}
		throw "failed to open";
	}

	newChannel(): Channel {
		var ch = new Channel();
		ch.remoteWin = 0;
		ch.myWindow = channelWindowSize;
		ch.ready = chan();
		ch.readBuf = chan();
		ch.session = this;
		ch.localId = this.addCh(ch);
		return ch;
	}

	async loop() {
		try {
			setInterval(async () => {
				var packet = await this.readPacket();
				if (packet[0] == msgChannelOpen) {
					await this.handleChannelOpen(packet);
					return;
				}
				var data = new DataView(packet);
				var id = data.getUint32(1);
				var ch = this.getCh(id);
				if (ch === undefined) {
					throw "invalid channel ("+id+") on op "+packet[0];
				}
				await ch.handlePacket(data);
			}, 20);
		} finally {}
		// catch {
		// 	this.channels.forEach(async (ch) => {
		// 		await ch.close();
		// 	})
		// 	this.channels = [];
		// 	await this.conn.close();
		// }
	}

	getCh(id: number): Channel {
		return this.channels[id];
	}
	
	addCh(ch: Channel): number {
		this.channels.forEach((v,i) => {
			if (v === undefined) {
				this.channels[i] = ch;
				return i;
			}
		});
		this.channels.push(ch);
		return this.channels.length-1;
	}

	rmCh(id: number) {
		this.channels[id] = undefined;
	}

	accept(): Promise<IChannel> {
		return this.incoming.shift();
	}

	close(): Promise<void> {
		return this.conn.close();
	}
}

export class Channel {
	localId: number;
	remoteId: number;
	maxIncomingPayload: number;
	maxRemotePayload: number;
	session: Session;
	ready: chan;
	sentEOF: boolean;
	sentClose: boolean;
	remoteWin: number;
	myWindow: number;
	readBuf: chan;

	ident(): number {
		return this.localId;
	}

	sendPacket(packet: Uint8Array): Promise<number> {
		if (this.sentClose) {
			throw "EOF";
		}
		this.sentClose = (packet[0] === msgChannelClose);
		return this.session.conn.send(packet.buffer);
	}

	sendMessage(type: number, msg: any): Promise<number> {
		var data = new DataView(encode(type, msg));
		data.setUint32(1, this.remoteId);
		return this.sendPacket(new Uint8Array(data.buffer));
	}

	async handlePacket(packet: DataView) {
		if (packet.buffer[0] === msgChannelData) {
			this.handleData(packet);
			return; 
		}
		if (packet.buffer[0] === msgChannelClose) {
			await this.sendMessage(msgChannelClose, {
				peersID: this.remoteId
			});
			this.session.rmCh(this.localId);
			await this.handleClose();
			return;
		}
		if (packet.buffer[0] === msgChannelEOF) {
			// TODO
			return;
		}
		if (packet.buffer[0] === msgChannelOpenFailure) {
			var fmsg:channelOpenFailureMsg = decode(packet.buffer);
			this.session.rmCh(fmsg.peersID);
			await this.ready.push(false);
			return;
		}
		if (packet.buffer[0] === msgChannelOpenConfirm) {
			var cmsg:channelOpenConfirmMsg = decode(packet.buffer);
			if (cmsg.maxPacketSize < minPacketLength || cmsg.maxPacketSize > 1<<31) {
				throw "invalid max packet size";
			}
			this.remoteId = cmsg.myID;
			this.maxRemotePayload = cmsg.maxPacketSize;
			this.remoteWin += cmsg.myWindow;
			await this.ready.push(true);
			return;
		}
		if (packet.buffer[0] === msgChannelWindowAdjust) {
			var amsg:channelWindowAdjustMsg = decode(packet.buffer);
			this.remoteWin += amsg.additionalBytes;
		}
	}

	async handleData(packet: DataView) {
		var length = packet.getUint32(5);
		if (length == 0) {
			return;
		}
		if (length > this.maxIncomingPayload) {
			throw "incoming packet exceeds maximum payload size";
		}
		var data = packet.buffer.slice(9, length);
		// TODO: check packet length
		if (this.myWindow < length) {
			throw "remot side wrote too much";
		}
		this.myWindow -= length;
		await this.readBuf.push(data);
	}

	async adjustWindow(n: number) {
		// TODO
	}

	async recv(): Promise<Uint8Array> {
		return Promise.resolve(new Uint8Array(await this.readBuf.shift()));
	}

	send(buffer: ArrayBuffer): Promise<number> {
		if (this.sentEOF) {
			return Promise.reject("EOF");
		}
		// TODO: use window
		var header = new DataView(new ArrayBuffer(9));
		header.setUint8(0, msgChannelData);
		header.setUint32(1, this.remoteId);
		header.setUint32(5, buffer.byteLength);
		var packet = new Uint8Array(9+buffer.byteLength);
		packet.set(new Uint8Array(header.buffer), 0);
		packet.set(new Uint8Array(buffer), 9);
		return this.sendPacket(packet);
	}

	async handleClose(): Promise<void> {
		await this.ready.close();
		this.sentClose = true;
	}

	async close() {
		await this.sendMessage(msgChannelClose, {
			peersID: this.remoteId
		});
	}

	async closeWrite() {
		this.sentEOF = true;
		await this.sendMessage(msgChannelEOF, {
			peersID: this.remoteId
		});
	}
}

function encode(type: number, obj: any): ArrayBuffer {
	switch (type) {
		case msgChannelClose:
			var data = new DataView(new ArrayBuffer(5));
			data.setUint8(0, type);
			data.setUint32(1, (<channelCloseMsg>obj).peersID);
			return data.buffer;
		case msgChannelData:
			var datamsg = <channelDataMsg>obj;
			var data = new DataView(new ArrayBuffer(9));
			data.setUint8(0, type);
			data.setUint32(1, datamsg.peersID);
			data.setUint32(5, datamsg.length);
			var buf = new Uint8Array(9+datamsg.length);
			buf.set(new Uint8Array(data.buffer), 0);
			buf.set(datamsg.rest, 9);
			return buf.buffer;
		case msgChannelEOF:
			var data = new DataView(new ArrayBuffer(5));
			data.setUint8(0, type);
			data.setUint32(1, (<channelEOFMsg>obj).peersID);
			return data.buffer;
		case msgChannelOpen:
			var data = new DataView(new ArrayBuffer(13));
			var openmsg = <channelOpenMsg>obj;
			data.setUint8(0, type);
			data.setUint32(1, openmsg.peersID);
			data.setUint32(5, openmsg.peersWindow);
			data.setUint32(9, openmsg.maxPacketSize);
			return data.buffer;
		case msgChannelOpenConfirm:
			var data = new DataView(new ArrayBuffer(17));
			var confirmmsg = <channelOpenConfirmMsg>obj;
			data.setUint8(0, type);
			data.setUint32(1, confirmmsg.peersID);
			data.setUint32(5, confirmmsg.myID);
			data.setUint32(9, confirmmsg.myWindow);
			data.setUint32(13, confirmmsg.maxPacketSize);
			return data.buffer;
		case msgChannelOpenFailure:
			var data = new DataView(new ArrayBuffer(5));
			data.setUint8(0, type);
			data.setUint32(1, (<channelOpenFailureMsg>obj).peersID);
			return data.buffer;
		case msgChannelWindowAdjust:
			var data = new DataView(new ArrayBuffer(9));
			var adjustmsg = <channelWindowAdjustMsg>obj;
			data.setUint8(0, type);
			data.setUint32(1, adjustmsg.peersID);
			data.setUint32(5, adjustmsg.additionalBytes);
			return data.buffer;
		default:
			throw "unknown type";
	}
}

function decode(packet: ArrayBuffer): any {
	switch (packet[0]) {
		case msgChannelClose:
			var data = new DataView(new ArrayBuffer(5));
			var closeMsg:channelCloseMsg = {
				peersID: data.getUint32(1)
			};
			return closeMsg;
		case msgChannelData:
			var data = new DataView(new ArrayBuffer(9));
			var dataLength = data.getUint32(5);
			var dataMsg:channelDataMsg = {
				peersID: data.getUint32(1),
				length: dataLength,
				rest: new Uint8Array(dataLength),
			};
			dataMsg.rest.set(new Uint8Array(data.buffer.slice(9)));
			return dataMsg;
		case msgChannelEOF:
			var data = new DataView(new ArrayBuffer(5));
			var eofMsg:channelEOFMsg = {
				peersID: data.getUint32(1),
			};
			return eofMsg;
		case msgChannelOpen:
			var data = new DataView(new ArrayBuffer(13));
			var openMsg:channelOpenMsg = {
				peersID: data.getUint32(1),
				peersWindow: data.getUint32(5),
				maxPacketSize: data.getUint32(9),
			};
			return openMsg;
		case msgChannelOpenConfirm:
			var data = new DataView(new ArrayBuffer(17));
			var confirmMsg:channelOpenConfirmMsg = {
				peersID: data.getUint32(1),
				myID: data.getUint32(5),
				myWindow: data.getUint32(9),
				maxPacketSize: data.getUint32(13),
			};
			return confirmMsg;
		case msgChannelOpenFailure:
			var data = new DataView(new ArrayBuffer(5));
			var failureMsg:channelOpenFailureMsg = {
				peersID: data.getUint32(1),
			};
			return failureMsg;
		case msgChannelWindowAdjust:
			var data = new DataView(new ArrayBuffer(9));
			var adjustMsg:channelWindowAdjustMsg = {
				peersID: data.getUint32(1),
				additionalBytes: data.getUint32(5),
			};
			return adjustMsg;
		default:
			throw "unknown type";
	}
}