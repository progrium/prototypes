"use strict";
exports.__esModule = true;
var ffi = require("ffi");
var ref = require("ref");
var ArrayType = require("ref-array");
var Struct = require("ref-struct");
var ByteArray = ArrayType(ref.types.uint8);
var GoString = Struct({
    p: "string",
    n: "longlong"
});
function goStr(str) {
    var s = new GoString();
    s["p"] = str;
    s["n"] = s["p"].length;
    return s;
}
var libmux = ffi.Library(__dirname + "/../libmux", {
    Error: ["int", ["int", ByteArray, "int"]],
    DialTCP: ["int", [GoString]],
    ListenTCP: ["int", [GoString]],
    DialWebsocket: ["int", [GoString]],
    ListenWebsocket: ["int", [GoString]],
    ListenerClose: ["int", ["int"]],
    ListenerAccept: ["int", ["int"]],
    SessionOpen: ["int", ["int"]],
    SessionAccept: ["int", ["int"]],
    SessionClose: ["int", ["int"]],
    ChannelRead: ["int", ["int", ByteArray, "int"]],
    ChannelWrite: ["int", ["int", ByteArray, "int"]],
    ChannelClose: ["int", ["int"]]
});
function lookupErr(id) {
    var buf = ByteArray(1 << 8);
    var n = libmux.Error(id * -1, buf, buf.length);
    return buf.buffer.slice(0, n).toString('ascii');
}
function handle(reject, name, retHandler) {
    return function (err, retcode) {
        if (err) {
            reject("ffi: " + err);
            return;
        }
        if (retcode < 0) {
            reject(name + "[" + (retcode * -1) + "]: " + lookupErr(retcode));
            return;
        }
        retHandler(retcode);
    };
}
function ListenTCP(addr) {
    return new Promise(function (resolve, reject) {
        libmux.ListenTCP.async(goStr(addr), handle(reject, "ListenTCP", function (retcode) {
            if (retcode === 0) {
                resolve();
                return;
            }
            resolve(new Listener(retcode));
        }));
    });
}
exports.ListenTCP = ListenTCP;
function DialTCP(addr) {
    return new Promise(function (resolve, reject) {
        libmux.DialTCP.async(goStr(addr), handle(reject, "DialTCP", function (retcode) {
            if (retcode === 0) {
                resolve();
                return;
            }
            resolve(new Session(retcode));
        }));
    });
}
exports.DialTCP = DialTCP;
function ListenWebsocket(addr) {
    return new Promise(function (resolve, reject) {
        libmux.ListenWebsocket.async(goStr(addr), handle(reject, "ListenWebsocket", function (retcode) {
            if (retcode === 0) {
                resolve();
                return;
            }
            resolve(new Listener(retcode));
        }));
    });
}
exports.ListenWebsocket = ListenWebsocket;
function DialWebsocket(addr) {
    return new Promise(function (resolve, reject) {
        libmux.DialWebsocket.async(goStr(addr), handle(reject, "DialWebsocket", function (retcode) {
            if (retcode === 0) {
                resolve();
                return;
            }
            resolve(new Session(retcode));
        }));
    });
}
exports.DialWebsocket = DialWebsocket;
var Listener = /** @class */ (function () {
    function Listener(id) {
        var _this = this;
        this.id = id;
        this.closed = false;
        process.once('SIGINT', function (code) {
            if (!_this.closed)
                libmux.ListenerClose(id);
        });
    }
    Listener.prototype.accept = function () {
        var _this = this;
        if (this.closed)
            return new Promise(function (r) { return r(); });
        return new Promise(function (resolve, reject) {
            libmux.ListenerAccept.async(_this.id, handle(reject, "ListenerAccept", function (retcode) {
                if (retcode === 0) {
                    resolve();
                    return;
                }
                resolve(new Session(retcode));
            }));
        });
    };
    Listener.prototype.close = function () {
        var _this = this;
        if (this.closed)
            return Promise.resolve();
        return new Promise(function (resolve, reject) {
            libmux.ListenerClose.async(_this.id, handle(reject, "ListenerClose", function () {
                _this.closed = true;
                resolve();
            }));
        });
    };
    return Listener;
}());
var Session = /** @class */ (function () {
    function Session(id) {
        this.id = id;
        this.closed = false;
    }
    Session.prototype.open = function () {
        var _this = this;
        if (this.closed)
            return new Promise(function (r) { return r(); });
        return new Promise(function (resolve, reject) {
            libmux.SessionOpen.async(_this.id, handle(reject, "SessionOpen", function (retcode) {
                if (retcode === 0) {
                    resolve();
                    return;
                }
                resolve(new Channel(retcode));
            }));
        });
    };
    Session.prototype.accept = function () {
        var _this = this;
        if (this.closed)
            return new Promise(function (r) { return r(); });
        return new Promise(function (resolve, reject) {
            libmux.SessionAccept.async(_this.id, handle(reject, "SessionAccept", function (retcode) {
                if (retcode === 0) {
                    resolve();
                    return;
                }
                resolve(new Channel(retcode));
            }));
        });
    };
    Session.prototype.close = function () {
        var _this = this;
        if (this.closed)
            return Promise.resolve();
        return new Promise(function (resolve, reject) {
            libmux.SessionClose.async(_this.id, handle(reject, "SessionClose", function () {
                _this.closed = true;
                resolve();
            }));
        });
    };
    return Session;
}());
var Channel = /** @class */ (function () {
    function Channel(id) {
        this.id = id;
        this.closed = false;
    }
    Channel.prototype.read = function (len) {
        var _this = this;
        if (this.closed)
            return new Promise(function (r) { return r(); });
        return new Promise(function (resolve, reject) {
            var buffer = ByteArray(len);
            libmux.ChannelRead.async(_this.id, buffer, buffer.length, handle(reject, "ChannelRead", function (retcode) {
                if (retcode === 0) {
                    _this.closed = true;
                    resolve();
                    return;
                }
                resolve(buffer.buffer.slice(0, retcode));
            }));
        });
    };
    Channel.prototype.write = function (buf) {
        var _this = this;
        if (this.closed)
            return new Promise(function (r) { return r(); });
        return new Promise(function (resolve, reject) {
            var buffer = ByteArray(buf);
            libmux.ChannelWrite.async(_this.id, buffer, buffer.length, handle(reject, "ChannelWrite", function (retcode) { return resolve(retcode); }));
        });
    };
    Channel.prototype.close = function () {
        var _this = this;
        if (this.closed)
            return Promise.resolve();
        return new Promise(function (resolve, reject) {
            libmux.ChannelClose.async(_this.id, handle(reject, "ChannelClose", function () {
                _this.closed = true;
                resolve();
            }));
        });
    };
    return Channel;
}());
