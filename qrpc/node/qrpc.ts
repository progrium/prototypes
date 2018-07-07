import * as chan from "@nodeguy/channel";
import * as msgpack from "msgpack-lite";

interface Session {
	open(): Promise<Channel>;
	accept(): Promise<Channel>;
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
                    this.decoder.end();
                    return;
                }
                this.decoder.write(buf);
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
    handlers: { [key:string]:Handler; }

    handle(path: string, handler: Handler): void {
        this.handlers[path] = handler;
    }

    handleFunc(path: string, handler: (r: Responder, c: Call) => void): void {
        this.handle(path, {
            serveRPC: (rr: Responder, cc: Call) => {
                handler(rr, cc);
            }
        })
    }

    async serve(session: Session, ch: Channel): Promise<void> {
        var codec = new MsgpackCodec(ch);
        var call = await codec.decode();
	    //call.parse()
        call.decode = codec.decode;
        call.Caller = new Client(session);
	    var header = new ResponseHeader();
        var resp = new responder(ch, header);
        if (!this.handlers.hasOwnProperty(call.Destination)) {
            resp.return(new Error("handler does not exist for this destination"));
        }
	    this.handlers[call.Destination].serveRPC(resp, call);
        return Promise.resolve(undefined);
    }
}

interface Responder {
    header: ResponseHeader;
    return(v: any): void;
}

class responder implements Responder {
    header: ResponseHeader;
    ch: Channel;

    constructor(ch: Channel, header: ResponseHeader) {
        this.ch = ch;
        this.header = header;
    }

    async return(v: any): Promise<void> {
        if (v instanceof Error) {
            this.header.Error = v.message;
            v = null;
        }
        var codec = new MsgpackCodec(this.ch);
        await codec.encode(this.header);
        await codec.encode(v);
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
            this.api.serve(this.session, ch)
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