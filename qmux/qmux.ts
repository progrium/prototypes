
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

interface Session {
	open(): Promise<Channel>;
	accept(): Promise<Channel>;
    close();
    wait(): Promise<void>;
}

// https://nodejs.org/api/fs.html#fs_fs_promises_api 
interface Channel {
	// Read reads up to len(data) bytes from the channel.
	Read(data []byte) (int, error)

	// Write writes len(data) bytes to the channel.
	Write(data []byte) (int, error)

	// Close signals end of channel use. No data may be sent after this
	// call.
	Close() error

	// CloseWrite signals the end of sending in-band
	// data. Requests may still be sent, and the other side may
	// still send data
	CloseWrite() error

	ID() uint32
}