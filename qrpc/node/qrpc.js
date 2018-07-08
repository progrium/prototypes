"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
var chan = require("@nodeguy/channel");
var msgpack = require("msgpack-lite");
function errable(p) {
    return p
        .then(function (ret) { return [ret, null]; })["catch"](function (err) { return [null, err]; });
}
function sleep(ms) {
    return new Promise(function (res) { return setTimeout(res, ms); });
}
// only one codec per channel because of read loop!
var MsgpackCodec = /** @class */ (function () {
    function MsgpackCodec(channel) {
        this.channel = channel;
        this.decoder = msgpack.createDecodeStream();
        var ch = chan();
        this.decoder.on("data", function (obj) {
            ch.push(obj);
        });
        this.objChan = ch;
        this.readLoop();
    }
    MsgpackCodec.prototype.readLoop = function () {
        return __awaiter(this, void 0, void 0, function () {
            var buf, e_1;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (!true) return [3 /*break*/, 6];
                        _a.label = 1;
                    case 1:
                        _a.trys.push([1, 4, , 5]);
                        return [4 /*yield*/, this.channel.read(1 << 16)];
                    case 2:
                        buf = _a.sent();
                        if (buf === undefined) {
                            return [2 /*return*/];
                        }
                        this.decoder.write(buf);
                        console.log("codec readloop");
                        console.log(buf);
                        return [4 /*yield*/, sleep(500)];
                    case 3:
                        _a.sent();
                        return [3 /*break*/, 5];
                    case 4:
                        e_1 = _a.sent();
                        throw "codec readLoop: " + e_1;
                    case 5: return [3 /*break*/, 0];
                    case 6: return [2 /*return*/];
                }
            });
        });
    };
    MsgpackCodec.prototype.encode = function (v) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.channel.write(msgpack.encode(v))];
                    case 1:
                        _a.sent();
                        return [2 /*return*/, Promise.resolve()];
                }
            });
        });
    };
    MsgpackCodec.prototype.decode = function () {
        return this.objChan.shift();
    };
    return MsgpackCodec;
}());
var Error = /** @class */ (function () {
    function Error(message) {
        this.message = message;
    }
    return Error;
}());
exports.Error = Error;
var API = /** @class */ (function () {
    function API() {
        this.handlers = {};
    }
    API.prototype.handle = function (path, handler) {
        this.handlers[path] = handler;
    };
    API.prototype.handleFunc = function (path, handler) {
        var _this = this;
        this.handle(path, {
            serveRPC: function (rr, cc) { return __awaiter(_this, void 0, void 0, function () {
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, handler(rr, cc)];
                        case 1:
                            _a.sent();
                            return [2 /*return*/];
                    }
                });
            }); }
        });
    };
    API.prototype.handler = function (path) {
        for (var p in this.handlers) {
            if (this.handlers.hasOwnProperty(p)) {
                if (path.startsWith(p)) {
                    return this.handlers[p];
                }
            }
        }
    };
    API.prototype.serveAPI = function (session, ch) {
        return __awaiter(this, void 0, void 0, function () {
            var codec, cdata, call, header, resp, handler;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        codec = new MsgpackCodec(ch);
                        return [4 /*yield*/, codec.decode()];
                    case 1:
                        cdata = _a.sent();
                        call = new Call(cdata.Destination);
                        call.parse();
                        call.decode = codec.decode;
                        call.caller = new Client(session);
                        header = new ResponseHeader();
                        resp = new responder(ch, codec, header);
                        handler = this.handler(call.Destination);
                        if (!handler) {
                            resp["return"](new Error("handler does not exist for this destination"));
                            return [2 /*return*/];
                        }
                        return [4 /*yield*/, handler.serveRPC(resp, call)];
                    case 2:
                        _a.sent();
                        return [2 /*return*/, Promise.resolve()];
                }
            });
        });
    };
    return API;
}());
exports.API = API;
var responder = /** @class */ (function () {
    function responder(ch, codec, header) {
        this.ch = ch;
        this.codec = codec;
        this.header = header;
    }
    responder.prototype["return"] = function (v) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (v instanceof Error) {
                            this.header.Error = v.message;
                            v = null;
                        }
                        return [4 /*yield*/, this.codec.encode(this.header)];
                    case 1:
                        _a.sent();
                        return [4 /*yield*/, this.codec.encode(v)];
                    case 2:
                        _a.sent();
                        return [2 /*return*/, this.ch.close()];
                }
            });
        });
    };
    return responder;
}());
var ResponseHeader = /** @class */ (function () {
    function ResponseHeader() {
    }
    return ResponseHeader;
}());
var Call = /** @class */ (function () {
    function Call(Destination) {
        this.Destination = Destination;
    }
    Call.prototype.parse = function () {
        // TODO
    };
    return Call;
}());
var Client = /** @class */ (function () {
    function Client(session, api) {
        this.session = session;
        this.api = api;
    }
    Client.prototype.serveAPI = function () {
        return __awaiter(this, void 0, void 0, function () {
            var ch;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (this.api === undefined) {
                            this.api = new API();
                        }
                        _a.label = 1;
                    case 1:
                        if (!true) return [3 /*break*/, 4];
                        return [4 /*yield*/, this.session.accept()];
                    case 2:
                        ch = _a.sent();
                        if (ch === undefined) {
                            return [2 /*return*/];
                        }
                        this.api.serveAPI(this.session, ch);
                        console.log("client serveAPI");
                        return [4 /*yield*/, sleep(500)];
                    case 3:
                        _a.sent();
                        return [3 /*break*/, 1];
                    case 4: return [2 /*return*/];
                }
            });
        });
    };
    Client.prototype.close = function () {
        return this.session.close();
    };
    Client.prototype.call = function (path, args) {
        return __awaiter(this, void 0, void 0, function () {
            var ch, codec, resp, ret;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.session.open()];
                    case 1:
                        ch = _a.sent();
                        codec = new MsgpackCodec(ch);
                        return [4 /*yield*/, codec.encode(new Call(path))];
                    case 2:
                        _a.sent();
                        return [4 /*yield*/, codec.encode(args)];
                    case 3:
                        _a.sent();
                        return [4 /*yield*/, codec.decode()];
                    case 4:
                        resp = _a.sent();
                        if (resp.Error !== null) {
                            throw resp.Error;
                        }
                        return [4 /*yield*/, codec.decode()];
                    case 5:
                        ret = _a.sent();
                        return [4 /*yield*/, ch.close()];
                    case 6:
                        _a.sent();
                        return [2 /*return*/, Promise.resolve(ret)];
                }
            });
        });
    };
    return Client;
}());
exports.Client = Client;
var Server = /** @class */ (function () {
    function Server() {
    }
    Server.prototype.serveAPI = function (sess) {
        return __awaiter(this, void 0, void 0, function () {
            var ch;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (!true) return [3 /*break*/, 3];
                        return [4 /*yield*/, sess.accept()];
                    case 1:
                        ch = _a.sent();
                        if (ch === undefined) {
                            return [2 /*return*/];
                        }
                        this.API.serveAPI(sess, ch);
                        console.log("server serveAPI");
                        return [4 /*yield*/, sleep(500)];
                    case 2:
                        _a.sent();
                        return [3 /*break*/, 0];
                    case 3: return [2 /*return*/];
                }
            });
        });
    };
    Server.prototype.serve = function (l, api) {
        return __awaiter(this, void 0, void 0, function () {
            var sess;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (!api) {
                            this.API = api;
                        }
                        if (!this.API) {
                            this.API = new API();
                        }
                        _a.label = 1;
                    case 1:
                        if (!true) return [3 /*break*/, 4];
                        return [4 /*yield*/, l.accept()];
                    case 2:
                        sess = _a.sent();
                        if (sess === undefined) {
                            return [2 /*return*/];
                        }
                        this.serveAPI(sess);
                        console.log("server serve");
                        return [4 /*yield*/, sleep(500)];
                    case 3:
                        _a.sent();
                        return [3 /*break*/, 1];
                    case 4: return [2 /*return*/];
                }
            });
        });
    };
    return Server;
}());
exports.Server = Server;
