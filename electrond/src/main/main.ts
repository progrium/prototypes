import { app } from "electron";
import * as qrpc from "qrpc";
import * as libmux from "libmux";

import apimodule from "./api";

let listener: any;

app.on("ready", async () => {
  var api = new qrpc.API();
  api.handle("dock", qrpc.Export(apimodule.app.dock));
  listener = await libmux.ListenWebsocket("localhost:4242");
  var server = new qrpc.Server();
  console.log("serving..."); 
  console.log(apimodule.app.dock);
  await server.serve(listener, api);
});

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
