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
function register(api) {
    api.handleFunc("app.quit", (r, c) => __awaiter(this, void 0, void 0, function* () {
        var obj = yield c.decode();
        var args = [].map((param) => {
            return util.argX((obj || {})[param]);
        });
        console.log("app.quit", obj);
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
        console.log("app.focus", obj);
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
        console.log("app.hide", obj);
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
        console.log("app.show", obj);
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
        console.log("app.getAppPath", obj);
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
        console.log("app.getPath", obj);
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
        console.log("app.getFileIcon", obj);
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
        console.log("app.getVersion", obj);
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
        console.log("app.getLocale", obj);
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
        console.log("app.getAppMetrics", obj);
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
        console.log("app.setBadgeCount", obj);
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
        console.log("app.getBadgeCount", obj);
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
        console.log("app.dock.bounce", obj);
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
        console.log("app.dock.cancelBounce", obj);
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
        console.log("app.dock.downloadFinished", obj);
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
        console.log("app.dock.setBadge", obj);
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
        console.log("app.dock.getBadge", obj);
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
        console.log("app.dock.hide", obj);
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
        console.log("app.dock.show", obj);
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
        console.log("app.dock.isVisible", obj);
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
        console.log("clipboard.readText", obj);
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
        console.log("clipboard.writeText", obj);
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
        console.log("clipboard.readHTML", obj);
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
        console.log("clipboard.writeHTML", obj);
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
        console.log("clipboard.readImage", obj);
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
        console.log("clipboard.writeImage", obj);
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
        console.log("clipboard.readRTF", obj);
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
        console.log("clipboard.writeRTF", obj);
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
        console.log("clipboard.readBookmark", obj);
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
        console.log("clipboard.writeBookmark", obj);
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
        console.log("clipboard.clear", obj);
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
        console.log("clipboard.availableFormats", obj);
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
        console.log("desktopCapturer.getSources", obj);
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
        console.log("dialog.showOpenDialog", obj);
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
        console.log("dialog.showSaveDialog", obj);
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
        console.log("dialog.showMessageBox", obj);
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
        console.log("dialog.showErrorBox", obj);
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
        console.log("globalShortcut.register", obj);
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
        console.log("globalShortcut.isRegistered", obj);
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
        console.log("globalShortcut.unregister", obj);
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
        console.log("process.getCPUUsage", obj);
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
        console.log("process.getHeapStatistics", obj);
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
        console.log("process.getProcessMemoryInfo", obj);
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
        console.log("process.getSystemMemoryInfo", obj);
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
        console.log("protocol.registerStandardSchemes", obj);
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
        console.log("protocol.registerFileProtocol", obj);
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
        console.log("protocol.registerStringProtocol", obj);
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
        console.log("protocol.registerHttpProtocol", obj);
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
        console.log("protocol.unregisterProtocol", obj);
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
        console.log("protocol.isProtocolHandled", obj);
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
        console.log("screen.getCursorScreenPoint", obj);
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
        console.log("screen.getPrimaryDisplay", obj);
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
        console.log("screen.getAllDisplays", obj);
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
        console.log("screen.getDisplayNearestPoint", obj);
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
        console.log("screen.getDisplayMatching", obj);
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
        console.log("screen.screenToDipPoint", obj);
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
        console.log("screen.dipToScreenPoint", obj);
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
        console.log("shell.showItemInFolder", obj);
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
        console.log("shell.openItem", obj);
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
        console.log("shell.openExternal", obj);
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
        console.log("shell.moveItemToTrash", obj);
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
        console.log("shell.beep", obj);
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
        console.log("shell.writeShortcutLink", obj);
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
        console.log("shell.readShortcutLink", obj);
        try {
            var ret = electron.shell.readShortcutLink(...args);
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