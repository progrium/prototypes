var qrpc = require("./qrpc.js");
var libmux = require("libmux");

var session = null;
var listener = null;

process.on('exit', () => {
    process.exit();
});

(async () => {
    var api = new qrpc.API();
    api.handleFunc("echo", async (r, c) => {
        console.log("echoing!");
        r.return(await c.decode());
    });

    listener = await libmux.ListenTCP("localhost:4242");
    var server = new qrpc.Server();
    console.log("serving...");
    //server.serve(listener, api);

    console.log("connecting...");
    session = await libmux.DialTCP("localhost:4242");
    var client = new qrpc.Client(session);
    console.log(await client.call("echo", "Hello world!"));
    await session.close();
})().catch(async (err) => { 
    console.log(err);
    //process.exit(1);
});