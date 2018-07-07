var qrpc = require("./qrpc.js");
var libmux = require("libmux");

(async () => {
    var session = await libmux.DialTCP("localhost:4242");
    var client = new qrpc.Client(session);
    console.log(await client.call("echo", "Hello world?"));
    await session.close();
})();