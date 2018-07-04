
import * as msgpack from "msgpack-lite";

interface ISession {
	open(): Promise<IChannel>;
	accept(): Promise<IChannel>;
    close(): Promise<void>;
    // wait(): Promise<void>;
}

interface IChannel {
	recv(): Promise<Uint8Array>;
	send(buffer: Uint8Array): Promise<number>;
	close(): Promise<void>;
	closeWrite(): Promise<void>;
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

    async serve(session: ISession, ch: IChannel): Promise<void> {
        var buf = await ch.recv();
        var call = msgpack.decode(buf);
	    //call.parse()
        call.decode = async (): Promise<any> => {
            var buf = await ch.recv();
            return Promise.resolve(msgpack.decode(buf));
        };
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
    ch: IChannel;

    constructor(ch: IChannel, header: ResponseHeader) {
        this.ch = ch;
        this.header = header;
    }

    async return(v: any): Promise<void> {
        if (v instanceof Error) {
            this.header.Error = v.message;
            v = null;
        }
        await this.ch.send(msgpack.encode(this.header));
        await this.ch.send(msgpack.encode(v));
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
    session: ISession;
    api: API;

    constructor(session: ISession, api?: API) {
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
        await ch.send(msgpack.encode(new Call(path)));
        await ch.send(msgpack.encode(args));
        var buf = await ch.recv();
        let resp: ResponseHeader = msgpack.decode(buf);
        if (resp.Error !== null) {
            throw resp.Error
        }
        buf = await ch.recv();
        ch.close();
        return msgpack.decode(buf);
    }
}