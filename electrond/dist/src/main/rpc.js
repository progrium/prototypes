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
const electron = require("electron");
const util = require("./util");
function register(api, om) {
    api.handleFunc("app.quit", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.quit(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.focus", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.focus(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.hide", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.hide(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.show", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.show(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getAppPath", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getAppPath(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getPath", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["name"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getPath(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getFileIcon", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var callbackHandle = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["path", "options", "callback"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getFileIcon(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getVersion", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getVersion(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getLocale", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getLocale(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getAppMetrics", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getAppMetrics(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.setBadgeCount", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["count"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.setBadgeCount(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.getBadgeCount", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.getBadgeCount(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.bounce", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.bounce(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.cancelBounce", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["id"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.cancelBounce(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.downloadFinished", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["filePath"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.downloadFinished(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.setBadge", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["text"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.setBadge(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.getBadge", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.getBadge(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.hide", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.hide(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.show", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.show(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("app.dock.isVisible", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.app.dock.isVisible(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.readText", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.readText(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.writeText", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["text", "type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.writeText(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.readHTML", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.readHTML(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.writeHTML", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["markup", "type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.writeHTML(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.readImage", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.readImage(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.writeImage", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["image", "type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.writeImage(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.readRTF", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.readRTF(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.writeRTF", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["text", "type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.writeRTF(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.readBookmark", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.readBookmark(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.writeBookmark", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["title", "url", "type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.writeBookmark(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.clear", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.clear(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("clipboard.availableFormats", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["type"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.clipboard.availableFormats(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("desktopCapturer.getSources", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var callbackHandle = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["options", "callback"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.desktopCapturer.getSources(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("dialog.showOpenDialog", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["filePaths", "bookmarks"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.dialog.showOpenDialog(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("dialog.showSaveDialog", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["filename", "bookmark"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.dialog.showSaveDialog(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("dialog.showMessageBox", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["response", "checkboxChecked"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.dialog.showMessageBox(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("dialog.showErrorBox", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["title", "content"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.dialog.showErrorBox(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("globalShortcut.register", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var callbackHandle = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["accelerator", "callback"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.globalShortcut.register(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("globalShortcut.isRegistered", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["accelerator"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.globalShortcut.isRegistered(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("globalShortcut.unregister", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["accelerator"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.globalShortcut.unregister(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("process.getCPUUsage", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = process.getCPUUsage(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("process.getHeapStatistics", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = process.getHeapStatistics(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("process.getProcessMemoryInfo", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = process.getProcessMemoryInfo(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("process.getSystemMemoryInfo", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = process.getSystemMemoryInfo(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("protocol.registerStandardSchemes", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["schemes", "options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.protocol.registerStandardSchemes(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("protocol.registerFileProtocol", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var handlerHandle = obj["handler"];
        obj["handler"] = () => {
            try {
                c.caller.call(handlerHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["scheme", "handler"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.protocol.registerFileProtocol(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("protocol.registerStringProtocol", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var handlerHandle = obj["handler"];
        obj["handler"] = () => {
            try {
                c.caller.call(handlerHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["scheme", "handler"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.protocol.registerStringProtocol(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("protocol.registerHttpProtocol", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var handlerHandle = obj["handler"];
        obj["handler"] = () => {
            try {
                c.caller.call(handlerHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["scheme", "handler"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.protocol.registerHttpProtocol(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("protocol.unregisterProtocol", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["scheme"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.protocol.unregisterProtocol(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("protocol.isProtocolHandled", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["scheme"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret = electron.protocol.isProtocolHandled(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.getCursorScreenPoint", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.getCursorScreenPoint(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.getPrimaryDisplay", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.getPrimaryDisplay(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.getAllDisplays", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.getAllDisplays(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.getDisplayNearestPoint", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["point"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.getDisplayNearestPoint(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.getDisplayMatching", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["rect"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.getDisplayMatching(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.screenToDipPoint", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["point"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.screenToDipPoint(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("screen.dipToScreenPoint", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["point"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.screen.dipToScreenPoint(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.showItemInFolder", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["fullPath"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.showItemInFolder(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.openItem", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["fullPath"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.openItem(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.openExternal", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var callbackHandle = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        };
        var args = ["url", "options", "callback"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.openExternal(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.moveItemToTrash", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["fullPath"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.moveItemToTrash(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.beep", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.beep(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.writeShortcutLink", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["shortcutPath", "operation", "options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.writeShortcutLink(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("shell.readShortcutLink", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["shortcutPath"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.shell.readShortcutLink(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("nativeImage.createEmpty", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.nativeImage.createEmpty(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("nativeImage.createFromPath", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["path"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj = electron.nativeImage.createFromPath(...args);
            var ret = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("nativeImage.createFromBuffer", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["buffer", "options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.nativeImage.createFromBuffer(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("nativeImage.createFromDataURL", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["dataURL"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.nativeImage.createFromDataURL(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("nativeImage.createFromNamedImage", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["imageName", "hslShift"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.nativeImage.createFromNamedImage(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Menu.setApplicationMenu", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["menu"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.Menu.setApplicationMenu(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Menu.getApplicationMenu", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.Menu.getApplicationMenu(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Menu.sendActionToFirstResponder", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["action"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.Menu.sendActionToFirstResponder(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Menu.buildFromTemplate", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["template"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        if (args[0][0]["click"]) {
            var callbackHandle = args[0][0]["click"];
            args[0][0]["click"] = () => {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            };
        }
        console.log(c.Destination, obj);
        try {
            var newObj = electron.Menu.buildFromTemplate(...args);
            var ret = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Menu.make", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj = new electron.Menu(...args);
            newObj.serveRPC = (r, c) => __awaiter(this, void 0, void 0, function* () {
                var handlers = {};
                handlers["popup"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["options"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["closePopup"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["browserWindow"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["append"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["menuItem"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["getMenuItemById"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["id"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["insert"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["pos", "menuItem"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers[c.method](r, c);
            });
            var ret = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("MenuItem.make", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj = new electron.MenuItem(...args);
            newObj.serveRPC = (r, c) => __awaiter(this, void 0, void 0, function* () {
                var handlers = {};
                handlers[c.method](r, c);
            });
            var ret = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Tray.make", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["image"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        args[0] = om.object(args[0].ObjectPath).value;
        console.log(c.Destination, obj);
        try {
            var newObj = new electron.Tray(...args);
            newObj.serveRPC = (r, c) => __awaiter(this, void 0, void 0, function* () {
                var handlers = {};
                handlers["destroy"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = [].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setImage"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["image"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setPressedImage"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["image"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setToolTip"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["toolTip"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setTitle"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["title"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setHighlightMode"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["mode"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setIgnoreDoubleClickEvents"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["ignore"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["getIgnoreDoubleClickEvents"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = [].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["displayBalloon"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["options"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["popUpContextMenu"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["menu", "position"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["setContextMenu"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = ["menu"].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    args[0] = om.object(args[0].ObjectPath).value;
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["getBounds"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = [].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["isDestroyed"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = [].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers[c.method](r, c);
            });
            var ret = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Notification.isSupported", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret = electron.Notification.isSupported(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
    api.handleFunc("Notification.make", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = ["options"].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj = new electron.Notification(...args);
            newObj.serveRPC = (r, c) => __awaiter(this, void 0, void 0, function* () {
                var handlers = {};
                handlers["show"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = [].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers["close"] = (r, c) => __awaiter(this, void 0, void 0, function* () {
                    var obj = yield c.decode();
                    var args = [].map((param) => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                });
                handlers[c.method](r, c);
            });
            var ret = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    }));
}
exports.register = register;
//# sourceMappingURL=rpc.js.map