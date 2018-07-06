const libmux = require("./libmux");

(async () => {
  console.log("s|listening...")
  var listener = await libmux.ListenTCP("localhost:8383");
  console.log("s|waiting...")
  var session = await listener.Accept();
  console.log("s|opening...")
  var ch = await session.Open();
  console.log("s|writing...")
  await ch.Write(Buffer.from("Hello world"));
  console.log("s|done")
})();

(async () => {
  console.log("c|dialing...")
  var session = await libmux.DialTCP("localhost:8383");
  console.log("c|waiting...")
  var ch = await session.Accept();
  console.log("c|reading...")
  var buf = await ch.Read(11);
  console.log(buf.toString('ascii'));
  console.log("c|closing...")
  await session.Close();
})();