import * as qmux from "./qmux";
import * as WebSocket from "ws";

interface IListener {
    accept(): Promise<qmux.Session>;
    close(): Promise<void>;
}

interface IConn {
	read(len: number): Promise<Buffer>;
	write(buffer: Buffer): Promise<number>;
	close(): Promise<void>;
}

class queue {
	q: Array<any>
	waiters: Array<Function>
	closed: boolean

	constructor() {
		this.q = [];
		this.waiters = [];
		this.closed = false;
	}

	push(obj: any) {
		if (this.closed) throw "closed queue";
		if (this.waiters.length > 0) {
			this.waiters.shift()(obj);
			return;
		}
		this.q.push(obj);
	}

	shift(): Promise<any> {
		if (this.closed) return;
        return new Promise(resolve => {
            if (this.q.length > 0) {
                resolve(this.q.shift());
                return;
            }
            this.waiters.push(resolve);
        })
	}
	
	close() {
		if (this.closed) return;
		this.closed = true;
		this.waiters.forEach(waiter => {
			waiter(undefined);
		});
	}
}

export function DialWebsocket(addr: string): Promise<IConn> {
    return new Promise((resolve, reject) => {
        var socket = new WebSocket(addr);
        socket.on("open", () => {
            resolve(new Conn(socket));
        });
        socket.on("error", (err) => {
            reject(err);
        });
    })
}

export async function ListenWebsocket(port: number, host?: string): Promise<IListener> {
    var listener = new Listener();
    await listener.listen(port, host);
    return listener;
}

export class Listener implements IListener {
    server: WebSocket.Server
    pending: queue

	constructor() {
        this.pending = new queue();
    }
    
    listen(port: number, host?: string): Promise<void> {
        return new Promise(resolve => {
            this.server = new WebSocket.Server({
                host: host,
                port: port
            }, resolve);
            this.server.on("connection", (c) => {
                this.pending.push(new qmux.Session(new Conn(c)));
            });
            this.server.on("error", (err) => {
                throw err;
            });
            this.server.on("close", () => {
                this.pending.close();
            });
        });
    }

    accept(): Promise<qmux.Session> {
        return this.pending.shift();
    }
    
    close(): Promise<void> {
        return new Promise(resolve => {
            this.server.close(() => resolve());
        });
    }
}

export class Conn implements IConn {
    socket: WebSocket
    error: any

    constructor(socket: WebSocket) {
        this.socket = socket;
        this.socket.on('error', (err) => this.error = err);
    }
 
    read(len: number): Promise<Buffer> {
        const stream = this.socket;

        return new Promise((resolve, reject) => {
            if (this.error) {
                const err = this.error
                delete this.error
                return reject(err)
            }

            if (!stream.readable || stream.destroyed) {
                return resolve()
            }

            const readableHandler = () => {
                let chunk = stream.read(len);

                if (chunk != null) {
                    removeListeners();
                    resolve(chunk);
                }
            }

            const closeHandler = () => {
                removeListeners();
                resolve();
            }

            const endHandler = () => {
                removeListeners();
                resolve();
            }

            const errorHandler = (err) => {
                delete this.error;
                removeListeners();
                reject(err);
            }

            const removeListeners = () => {
                stream.removeListener('close', closeHandler);
                stream.removeListener('error', errorHandler);
                stream.removeListener('end', endHandler);
                stream.removeListener('readable', readableHandler);
            }

            stream.on('close', closeHandler);
            stream.on('end', endHandler);
            stream.on('error', errorHandler);
            stream.on('readable', readableHandler);

            readableHandler();
        });
    }

	write(buffer: Buffer): Promise<number> {
        return new Promise(resolve => {
            this.socket.write(buffer, () => resolve(buffer.byteLength));
        });
    }

	close(): Promise<void> {
        return new Promise(resolve => {
            this.socket.destroy();
            resolve();
        });
    }
}