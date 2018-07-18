import * as electron from "electron";
import * as qrpc from "qrpc";
import * as util from "./util";
export function register(api: qrpc.API, om: qrpc.ObjectManager) {
    api.handleFunc("app.quit", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.quit as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.focus", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.focus as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.hide", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.hide as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.show", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.show as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getAppPath", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getAppPath as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getPath", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["name"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getPath as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getFileIcon", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var callbackHandle: any = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["path", "options", "callback"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getFileIcon as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getVersion", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getVersion as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getLocale", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getLocale as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getAppMetrics", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getAppMetrics as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.setBadgeCount", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["count"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.setBadgeCount as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.getBadgeCount", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.getBadgeCount as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.bounce", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.bounce as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.cancelBounce", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["id"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.cancelBounce as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.downloadFinished", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["filePath"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.downloadFinished as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.setBadge", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["text"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.setBadge as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.getBadge", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.getBadge as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.hide", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.hide as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.show", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.show as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("app.dock.isVisible", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.app.dock.isVisible as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.readText", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.readText as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.writeText", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["text", "type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.writeText as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.readHTML", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.readHTML as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.writeHTML", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["markup", "type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.writeHTML as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.readImage", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.readImage as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.writeImage", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["image", "type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.writeImage as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.readRTF", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.readRTF as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.writeRTF", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["text", "type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.writeRTF as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.readBookmark", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.readBookmark as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.writeBookmark", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["title", "url", "type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.writeBookmark as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.clear", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.clear as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("clipboard.availableFormats", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["type"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.clipboard.availableFormats as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("desktopCapturer.getSources", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var callbackHandle: any = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["options", "callback"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.desktopCapturer.getSources as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("dialog.showOpenDialog", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["filePaths", "bookmarks"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.dialog.showOpenDialog as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("dialog.showSaveDialog", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["filename", "bookmark"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.dialog.showSaveDialog as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("dialog.showMessageBox", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["response", "checkboxChecked"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.dialog.showMessageBox as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("dialog.showErrorBox", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["title", "content"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.dialog.showErrorBox as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("globalShortcut.register", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var callbackHandle: any = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["accelerator", "callback"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.globalShortcut.register as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("globalShortcut.isRegistered", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["accelerator"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.globalShortcut.isRegistered as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("globalShortcut.unregister", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["accelerator"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.globalShortcut.unregister as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("process.getCPUUsage", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (process.getCPUUsage as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("process.getHeapStatistics", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (process.getHeapStatistics as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("process.getProcessMemoryInfo", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (process.getProcessMemoryInfo as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("process.getSystemMemoryInfo", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (process.getSystemMemoryInfo as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("protocol.registerStandardSchemes", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["schemes", "options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.protocol.registerStandardSchemes as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("protocol.registerFileProtocol", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var handlerHandle: any = obj["handler"];
        obj["handler"] = () => {
            try {
                c.caller.call(handlerHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["scheme", "handler"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.protocol.registerFileProtocol as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("protocol.registerStringProtocol", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var handlerHandle: any = obj["handler"];
        obj["handler"] = () => {
            try {
                c.caller.call(handlerHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["scheme", "handler"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.protocol.registerStringProtocol as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("protocol.registerHttpProtocol", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var handlerHandle: any = obj["handler"];
        obj["handler"] = () => {
            try {
                c.caller.call(handlerHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["scheme", "handler"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.protocol.registerHttpProtocol as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("protocol.unregisterProtocol", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["scheme"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.protocol.unregisterProtocol as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("protocol.isProtocolHandled", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["scheme"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        var cbArgs: any = ["error"];
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.protocol.isProtocolHandled as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.getCursorScreenPoint", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.getCursorScreenPoint as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.getPrimaryDisplay", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.getPrimaryDisplay as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.getAllDisplays", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.getAllDisplays as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.getDisplayNearestPoint", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["point"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.getDisplayNearestPoint as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.getDisplayMatching", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["rect"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.getDisplayMatching as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.screenToDipPoint", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["point"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.screenToDipPoint as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("screen.dipToScreenPoint", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["point"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.screen.dipToScreenPoint as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.showItemInFolder", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["fullPath"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.showItemInFolder as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.openItem", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["fullPath"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.openItem as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.openExternal", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var callbackHandle: any = obj["callback"];
        obj["callback"] = () => {
            try {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            }
            catch (e) {
                console.log("callback to missing session");
            }
        }
        var args: any = ["url", "options", "callback"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.openExternal as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.moveItemToTrash", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["fullPath"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.moveItemToTrash as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.beep", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.beep as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.writeShortcutLink", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["shortcutPath", "operation", "options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.writeShortcutLink as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("shell.readShortcutLink", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["shortcutPath"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.shell.readShortcutLink as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("nativeImage.createEmpty", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.nativeImage.createEmpty as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("nativeImage.createFromPath", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["path"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj: any = (electron.nativeImage.createFromPath as any)(...args);
			var ret: any = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("nativeImage.createFromBuffer", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["buffer", "options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.nativeImage.createFromBuffer as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("nativeImage.createFromDataURL", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["dataURL"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.nativeImage.createFromDataURL as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("nativeImage.createFromNamedImage", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["imageName", "hslShift"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.nativeImage.createFromNamedImage as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Menu.setApplicationMenu", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["menu"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.Menu.setApplicationMenu as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Menu.getApplicationMenu", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.Menu.getApplicationMenu as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Menu.sendActionToFirstResponder", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["action"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.Menu.sendActionToFirstResponder as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Menu.buildFromTemplate", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["template"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        if (args[0][0]["click"]) {
            var callbackHandle: any = args[0][0]["click"];
            args[0][0]["click"] = () => {
                c.caller.call(callbackHandle.ObjectPath + "/__call__", null);
            };
        }
        console.log(c.Destination, obj);
        try {
            var newObj: any = (electron.Menu.buildFromTemplate as any)(...args);
		    var ret: any = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Menu.make", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj: any = new (electron.Menu as any)(...args);
            newObj.serveRPC = async (r: qrpc.Responder, c: qrpc.Call) => {
                var handlers = {}
                handlers["popup"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["options"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["closePopup"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["browserWindow"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["append"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["menuItem"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["getMenuItemById"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["id"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["insert"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["pos", "menuItem"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers[c.method](r, c);
            }
            var ret: any = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("MenuItem.make", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj: any = new (electron.MenuItem as any)(...args);
            newObj.serveRPC = async (r: qrpc.Responder, c: qrpc.Call) => {
                var handlers = {}
                handlers[c.method](r, c);
            }
            var ret: any = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Tray.make", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["image"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        args[0] = om.object(args[0].ObjectPath).value;
        console.log(c.Destination, obj);
        try {
            var newObj: any = new (electron.Tray as any)(...args);
            newObj.serveRPC = async (r: qrpc.Responder, c: qrpc.Call) => {
                var handlers = {}
                handlers["destroy"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = [].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setImage"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["image"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setPressedImage"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["image"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setToolTip"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["toolTip"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setTitle"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["title"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setHighlightMode"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["mode"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setIgnoreDoubleClickEvents"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["ignore"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["getIgnoreDoubleClickEvents"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = [].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["displayBalloon"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["options"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["popUpContextMenu"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["menu", "position"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["setContextMenu"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = ["menu"].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    args[0] = om.object(args[0].ObjectPath).value;
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["getBounds"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = [].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["isDestroyed"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = [].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers[c.method](r, c);
            }
            var ret: any = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Notification.isSupported", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = [].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var ret: any = (electron.Notification.isSupported as any)(...args);
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
    api.handleFunc("Notification.make", async (r: qrpc.Responder, c: qrpc.Call) => {
        var obj: any = await c.decode();
        var args: any = ["options"].map((param: string): any => {
            return util.argX((obj || {})[param]);
        });
        console.log(c.Destination, obj);
        try {
            var newObj: any = new (electron.Notification as any)(...args);
            newObj.serveRPC = async (r: qrpc.Responder, c: qrpc.Call) => {
                var handlers = {}
                handlers["show"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = [].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers["close"] = async (r: qrpc.Responder, c: qrpc.Call) => {
                    var obj: any = await c.decode();
                    var args: any = [].map((param: string): any => {
                        return util.argX((obj || {})[param]);
                    });
                    console.log(c.Destination, obj);
                    try {
                        var objRef = om.object(c.objectPath);
                        var ret: any = objRef.value[c.method](...args);
                        r.return(ret);
                    }
                    catch (e) {
                        console.log(e.stack);
                        r.return(e);
                    }
                }
                handlers[c.method](r, c);
            }
            var ret: any = om.register(newObj).handle();
            r.return(ret);
        }
        catch (e) {
            console.log(e.stack);
            r.return(e);
        }
    });
}
