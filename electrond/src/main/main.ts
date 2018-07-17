import * as electron from "electron";
import { app, globalShortcut } from "electron";
import * as qrpc from "qrpc";
import * as libmux from "libmux";

import * as rpc from "./rpc";

let listener: any;

function sleep(ms: number): Promise<void> {
  return new Promise(res => setTimeout(res, ms));
}

app.on("ready", async () => {
  var api = new qrpc.API();
  rpc.register(api);
  listener = await libmux.ListenWebsocket("localhost:4242");
  var server = new qrpc.Server();
  console.log("serving..."); 
  //loop()
  await server.serve(listener, api);
});

async function loop() {
  while (true) {
    console.log("ping")
    await sleep(3000)
  }
}

// Quit when all windows are closed.
app.on("window-all-closed", () => {
  // On OS X it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("before-quit", () => {

});
