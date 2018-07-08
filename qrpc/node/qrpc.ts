import * as chan from "@nodeguy/channel";
import * as msgpack from "msgpack-lite";

interface Session {
	open(): Promise<Channel>;
	accept(): Promise<Channel>;
    close(): Promise<void>;
}

interface Listener {
    accept(): Promise<Session>;
    close(): Promise<void>;
}

interface Channel {
	read(len: number): Promise<Buffer>;
	write(buffer: Buffer): Promise<number>;
	close(): Promise<void>;
}

interface Codec {
    decode(): Promise<any>;
    encode(v: any): Promise<void>;
}

function errable(p: Promise<any>): Promise<any> {
    return p
        .then(ret => [ret, null])
        .catch(err => [null, err]);
}

function sleep(ms: number): Promise<void> {
    return new Promise(res => setTimeout(res, ms));
}

// only one codec per channel because of read loop!
class MsgpackCodec {
    channel: Channel;
    decoder: any;
    objChan: any;

    constructor(channel: Channel) {
        this.channel = channel;
        this.decoder = msgpack.createDecodeStream();
        var ch = chan();
        this.decoder.on("data", (obj) => {
            ch.push(obj);
        })
        this.objChan = ch;
        this.readLoop();
    }

    async readLoop() {
        while(true) {
            try {
                var buf = await this.channel.read(1 << 16);
                if (buf === undefined) {
                    return;
                }
                this.decoder.write(buf);
                console.log("codec readloop");
                console.log(buf);
                await sleep(500);
            } catch (e) {
                throw "codec readLoop: "+e;
            }
        }
    }

    async encode(v: any): Promise<void> {
        await this.channel.write(msgpack.encode(v));
        return Promise.resolve();
    }

    decode(): Promise<any> {
        return this.objChan.shift();
    }
}

export class Error {
    message: string;
    constructor(message: string) {
        this.message = message;
    }
}

export class API {
    handlers: { [key:string]:Handler; };

    constructor() {
        this.handlers = {};
    }

    handle(path: string, handler: Handler): void {
        this.handlers[path] = handler;
    }

    handleFunc(path: string, handler: (r: Responder, c: Call) => void): void {
        this.handle(path, {
            serveRPC: async (rr: Responder, cc: Call) => {
                await handler(rr, cc);
            }
        })
    }

    handler(path: string): Handler {
        for (var p in this.handlers) {
            if (this.handlers.hasOwnProperty(p)) {
                if (path.startsWith(p)) {
                    return this.handlers[p];
                }
            }
        }
    }

    async serveAPI(session: Session, ch: Channel): Promise<void> {
        var codec = new MsgpackCodec(ch);
        var cdata = await codec.decode();
        var call = new Call(cdata.Destination);
	    call.parse();
        call.decode = codec.decode;
        call.caller = new Client(session);
	    var header = new ResponseHeader();
        var resp = new responder(ch, codec, header);
        var handler = this.handler(call.Destination);
        if (!handler) {
            resp.return(new Error("handler does not exist for this destination"));
            return;
        }
        await handler.serveRPC(resp, call);
        return Promise.resolve();
    }
}

interface Responder {
    header: ResponseHeader;
    return(v: any): void;
}

class responder implements Responder {
    header: ResponseHeader;
    ch: Channel;
    codec: Codec;

    constructor(ch: Channel, codec: Codec, header: ResponseHeader) {
        this.ch = ch;
        this.codec = codec;
        this.header = header;
    }

    async return(v: any): Promise<void> {
        if (v instanceof Error) {
            this.header.Error = v.message;
            v = null;
        }
        await this.codec.encode(this.header);
        await this.codec.encode(v);
        return this.ch.close();
    }
}

class ResponseHeader {
    Error: string;
}

interface Caller {
    call(method: string, args: any): Promise<any>;
}

class Call {
	Destination: string;
	objectPath:  string;
	method:      string;
	caller:      Caller;
    decode:      () => Promise<any>;
    
    constructor(Destination: string) {
        this.Destination = Destination;
    }

    parse() {
        // TODO
    }
}

interface Handler {
	serveRPC(r: Responder, c: Call): void;
}

export class Client implements Caller {
    session: Session;
    api: API;

    constructor(session: Session, api?: API) {
        this.session = session;
        this.api = api;
    }

    async serveAPI(): Promise<void> {
        if (this.api === undefined) {
            this.api = new API();
        }
        while (true) {
            var ch = await this.session.accept();
            if (ch === undefined) {
                return;
            }
            this.api.serveAPI(this.session, ch)
            console.log("client serveAPI");
            await sleep(500);
        }
    }

    close(): Promise<void> {
        return this.session.close();
    }

    async call(path: string, args: any): Promise<any> {
        var ch = await this.session.open();
        var codec = new MsgpackCodec(ch);
        await codec.encode(new Call(path));
        await codec.encode(args);

        var resp: ResponseHeader = await codec.decode();
        if (resp.Error !== null) {
            throw resp.Error;
        }
        var ret = await codec.decode(); 
        await ch.close();
        return Promise.resolve(ret);
    }
}

export class Server {
    API: API;
    
    async serveAPI(sess: Session) {
        while (true) {
            var ch = await sess.accept();
            if (ch === undefined) {
                return;
            }
            this.API.serveAPI(sess, ch);
            console.log("server serveAPI");
            await sleep(500);
        }
    }

    async serve(l: Listener, api?: API) {
        if (!api) {
            this.API = api;
        }
        if (!this.API) {
            this.API = new API();
        }
        while (true) {
            var sess = await l.accept();
            if (sess === undefined) {
                return;
            }
            this.serveAPI(sess);
            console.log("server serve");
            await sleep(500);
        }
    }
}