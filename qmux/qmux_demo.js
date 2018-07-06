
var qmux = require('./qmux.js');
var ws = require("nodejs-websocket");

function wrapConn(conn) {
    return {
        recv: () => {
            return new Promise((resolve, reject) => {
                conn.on("binary", function (inStream) {
                    // Empty buffer for collecting binary data
                    var data = new Buffer(0);
                    // Read chunks of binary data and add to the buffer
                    inStream.on("readable", function () {
                        var newData = inStream.read();
                        if (newData)
                            data = Buffer.concat([data, newData], data.length+newData.length)
                    });
                    inStream.on("end", function () {
                        resolve(buffer.buffer);
                    });
                    inStream.on("error", function() {
                        reject();
                    });
                });
            });
        },
        send: (buffer) => {
            return new Promise((resolve, reject) => {
                if (!conn.sendBinary(new Uint8Array(buffer), () => {
                    resolve(buffer.length);
                })) {
                    reject();
                }
            });
        },
        close: () => {
            console.log(conn);
            throw new Exception();
            conn.close();
            return new Promise((resolve, reject) => {
                conn.once("close", () => {
                    resolve();
                })
                conn.once("error", () => {
                    reject();
                })
            });
        },
        closeWrite: () => {
            // TODO
        }
    }
}

var server = ws.createServer(async function (conn) {
    console.log("got connection!");
    var session = new qmux.Session(wrapConn(conn));
    var ch = await session.open();
    //console.log("channel open!");
}).listen(8001);

var client = ws.connect("ws://localhost:8001/", async function(conn) {
    console.log("connected!");
    var session = new qmux.Session(wrapConn(conn));
    var ch = await session.accept();
});