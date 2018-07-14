"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const electron_1 = require("electron");
const qrpc = require("qrpc");
const libmux = require("libmux");
const api_1 = require("./api");
let listener;
electron_1.app.on("ready", () => __awaiter(this, void 0, void 0, function* () {
    var api = new qrpc.API();
    api.handle("dock", qrpc.Export(api_1.default.app.dock));
    listener = yield libmux.ListenWebsocket("localhost:4242");
    var server = new qrpc.Server();
    console.log("serving...");
    console.log(api_1.default.app.dock);
    yield server.serve(listener, api);
}));
// Quit when all windows are closed.
electron_1.app.on("window-all-closed", () => {
    // On OS X it is common for applications and their menu bar
    // to stay active until the user quits explicitly with Cmd + Q
    if (process.platform !== "darwin") {
        electron_1.app.quit();
    }
});
electron_1.app.on("before-quit", () => {
});
//# sourceMappingURL=main.js.map