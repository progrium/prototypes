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
    var n = libmux.Error(id, buf, 1 << 8);
    return buf.buffer.slice(0, n).toString('ascii');
}
function ListenTCP(addr) {
    return new Promise(function (resolve, reject) {
        libmux.ListenTCP.async(goStr(addr), function (err, ret) {
            if (err) {
                reject(err);
                return;
            }
            if (ret < 0) {
                reject(lookupErr(ret));
                return;
            }
            if (ret === 0) {
                resolve();
                return;
            }
            resolve(new Listener(ret));
        });
    });
}
exports.ListenTCP = ListenTCP;
function DialTCP(addr) {
    return new Promise(function (resolve, reject) {
        libmux.DialTCP.async(goStr(addr), function (err, ret) {
            if (err) {
                reject(err);
                return;
            }
            if (ret < 0) {
                reject(lookupErr(ret));
                return;
            }
            if (ret === 0) {
                resolve();
                return;
            }
            resolve(new Session(ret));
        });
    });
}
exports.DialTCP = DialTCP;
var Listener = /** @class */ (function () {
    function Listener(id) {
        this.id = id;
        process.once('SIGINT', function (code) {
            libmux.ListenerClose(id);
        });
    }
    Listener.prototype.accept = function () {
        var _this = this;
        return new Promise(function (resolve, reject) {
            libmux.ListenerAccept.async(_this.id, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                if (ret === 0) {
                    resolve();
                    return;
                }
                resolve(new Session(ret));
            });
        });
    };
    Listener.prototype.close = function () {
        var _this = this;
        return new Promise(function (resolve, reject) {
            libmux.ListenerClose.async(_this.id, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                resolve();
            });
        });
    };
    return Listener;
}());
var Session = /** @class */ (function () {
    function Session(id) {
        this.id = id;
    }
    Session.prototype.open = function () {
        var _this = this;
        return new Promise(function (resolve, reject) {
            libmux.SessionOpen.async(_this.id, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                if (ret === 0) {
                    resolve();
                    return;
                }
                resolve(new Channel(ret));
            });
        });
    };
    Session.prototype.accept = function () {
        var _this = this;
        return new Promise(function (resolve, reject) {
            libmux.SessionAccept.async(_this.id, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                if (ret === 0) {
                    resolve();
                    return;
                }
                resolve(new Channel(ret));
            });
        });
    };
    Session.prototype.close = function () {
        var _this = this;
        return new Promise(function (resolve, reject) {
            libmux.SessionClose.async(_this.id, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                resolve();
            });
        });
    };
    return Session;
}());
var Channel = /** @class */ (function () {
    function Channel(id) {
        this.id = id;
    }
    Channel.prototype.read = function (len) {
        var _this = this;
        return new Promise(function (resolve, reject) {
            var buffer = ByteArray(len);
            libmux.ChannelRead.async(_this.id, buffer, buffer.length, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject("ERR" + lookupErr(ret));
                    return;
                }
                if (ret === 0) {
                    resolve();
                    return;
                }
                resolve(buffer.buffer.slice(0, ret));
            });
        });
    };
    Channel.prototype.write = function (buf) {
        var _this = this;
        return new Promise(function (resolve, reject) {
            var buffer = ByteArray(buf);
            libmux.ChannelWrite.async(_this.id, buffer, buffer.length, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                resolve(ret);
            });
        });
    };
    Channel.prototype.close = function () {
        var _this = this;
        return new Promise(function (resolve, reject) {
            libmux.ChannelClose.async(_this.id, function (err, ret) {
                if (err) {
                    reject(err);
                    return;
                }
                if (ret < 0) {
                    reject(lookupErr(ret));
                    return;
                }
                resolve();
            });
        });
    };
    return Channel;
}());
