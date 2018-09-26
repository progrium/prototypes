var qmux = require('./qmux.js');
var ws = require("nodejs-websocket");

class WSConn {
    constructor(conn, server) {
        this.server = server;
        this.conn = conn;
        this.buf = new Buffer(0);
        this.waiters = [];
        this.isClosed = false;
        this.conn.on("binary", stream => {
            stream.on("readable", () => {
                var buf = stream.read();
                if (buf !== undefined && this.buf !== undefined) {
                    this.buf = Buffer.concat([this.buf, buf], this.buf.length+buf.length);
                    if (this.waiters.length > 0) {
                        this.waiters.shift()();
                    }
                }
            });
        });
        this.conn.on("close", () => {
            this.close();
        });
        this.conn.on("error", (err) => {
            console.log("err", err)
        })
    }

    close() {
        if (this.isClosed) return Promise.resolve();
        return new Promise((resolve, reject) => {
            this.isClosed = true;
            this.buf = undefined;
            this.waiters.forEach(waiter => waiter());
            this.conn.close();
            resolve();
        });
    }

    write(buffer) {
        return new Promise((resolve, reject) => {
            if (!this.conn.sendBinary(buffer, () => {
                resolve(buffer.length)
            })) {
                reject("write failed");
            }
        });
    }

    read(len) {
        return new Promise((resolve, reject) => {
            var tryRead = () => {
                if (this.buf === undefined) {
                    resolve(undefined);
                    return;
                }
                if (this.buf.length >= len) {
                    var data = this.buf.slice(0, len);
                    this.buf = this.buf.slice(len);
                    resolve(data);
                    return;
                }
                this.waiters.push(tryRead);
            }
            tryRead();
        })
    }
}


var server = ws.createServer(async function (conn) {
    var session = new qmux.Session(new WSConn(conn, true));
    var ch = await session.open();
    console.log("|Server echoing on channel...");
    while(true) {
        var data = await ch.read(1);
        if (data === undefined) {
            console.log("|Server got EOF");
            break;
        }
        await ch.write(data);
    }
    server.close();
}).listen(8001);

var conn = ws.connect("ws://localhost:8001/", async function() {
    var session = new qmux.Session(new WSConn(conn, false));
    var ch = await session.accept();
    await ch.write(Buffer.from("Hello"));
    await ch.write(Buffer.from(" "));
    await ch.write(Buffer.from("world"));
    
    var data = await ch.read(11);
    console.log(data.toString('ascii'));

    await session.close();
});