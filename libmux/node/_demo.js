const libmux = require("./libmux");

(async () => {  
  console.log("s|listening...")
  var listener = await libmux.ListenTCP("localhost:8383");

  (async () => {  
    console.log("s|waiting...")
    var session = await listener.accept();
    console.log("s|opening...")
    var ch = await session.open();
    console.log("s|writing...")
    await ch.write(Buffer.from("Hello world"));
    console.log("s|done")
  })();

  (async () => {
    console.log("c|dialing...")
    var session = await libmux.DialTCP("localhost:8383");
    console.log("c|waiting...")
    var ch = await session.accept();
    console.log("c|reading...")
    var buf = await ch.read(11);
    console.log(buf.toString('ascii'));
    console.log("c|closing...")
    await session.close();
  })();

})();