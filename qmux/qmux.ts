
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
	rest:    string;
}

interface channelEOFMsg {
	peersID: number;
}

interface channelCloseMsg {
	peersID: number;
}

interface IConn {
	read(buffer: Uint8Array): Promise<number>;
	write(buffer: Uint8Array): Promise<number>;
	close(): Promise<void>;
	closeWrite(): Promise<void>;
}

interface ISession {
	open(): Promise<IChannel>;
	accept(): Promise<IChannel>;
    close(): Promise<void>;
    wait(): Promise<void>;
}

interface IChannel extends IConn {
	ident(): number
}

class Session {
	conn: IConn;
	channels: Array<IChannel>;

	constructor(conn: IConn) {
        this.conn = conn;
	}
	
	addCh(ch: IChannel): number {
		
	}

	rmCh(id: number) {

	}
}