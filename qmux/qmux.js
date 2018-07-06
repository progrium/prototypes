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
            if (f = 1, y && (t = y[op[0] & 2 ? "return" : op[0] ? "throw" : "next"]) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [0, t.value];
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
var msgChannelOpen = 100;
var msgChannelOpenConfirm = 101;
var msgChannelOpenFailure = 102;
var msgChannelWindowAdjust = 103;
var msgChannelData = 104;
var msgChannelEOF = 105;
var msgChannelClose = 106;
var minPacketLength = 9;
var channelMaxPacket = 1 << 15;
var channelWindowSize = 64 * channelMaxPacket;
var Session = /** @class */ (function () {
    function Session(conn) {
        this.conn = conn;
        this.channels = [];
        this.incoming = chan();
        this.loop();
    }
    Session.prototype.readPacket = function () {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                return [2 /*return*/, Promise.resolve(new ArrayBuffer(8))];
            });
        });
    };
    Session.prototype.handleChannelOpen = function (packet) {
        return __awaiter(this, void 0, void 0, function () {
            var msg, c;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        msg = decode(packet);
                        if (!(msg.maxPacketSize < minPacketLength || msg.maxPacketSize > 1 << 31)) return [3 /*break*/, 2];
                        return [4 /*yield*/, this.conn.send(encode(msgChannelOpenFailure, {
                                peersID: msg.peersID
                            }))];
                    case 1:
                        _a.sent();
                        _a.label = 2;
                    case 2:
                        c = this.newChannel();
                        c.remoteId = msg.peersID;
                        c.maxRemotePayload = msg.maxPacketSize;
                        c.remoteWin = msg.peersWindow;
                        c.maxIncomingPayload = channelMaxPacket;
                        return [4 /*yield*/, this.incoming.push(c)];
                    case 3:
                        _a.sent();
                        return [4 /*yield*/, this.conn.send(encode(msgChannelOpenConfirm, {
                                peersID: c.remoteId,
                                myID: c.localId,
                                myWindow: c.myWindow,
                                maxPacketSize: c.maxIncomingPayload
                            }))];
                    case 4:
                        _a.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    Session.prototype.open = function () {
        return __awaiter(this, void 0, void 0, function () {
            var ch;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        ch = this.newChannel();
                        ch.maxIncomingPayload = channelMaxPacket;
                        return [4 /*yield*/, this.conn.send(encode(msgChannelOpen, {
                                peersWindow: ch.myWindow,
                                maxPacketSize: ch.maxIncomingPayload,
                                peersID: ch.localId
                            }))];
                    case 1:
                        _a.sent();
                        return [4 /*yield*/, ch.ready.shift()];
                    case 2:
                        if (_a.sent()) {
                            return [2 /*return*/, Promise.resolve(ch)];
                        }
                        throw "failed to open";
                }
            });
        });
    };
    Session.prototype.newChannel = function () {
        var ch = new Channel();
        ch.remoteWin = 0;
        ch.myWindow = channelWindowSize;
        ch.ready = chan();
        ch.readBuf = chan();
        ch.session = this;
        ch.localId = this.addCh(ch);
        return ch;
    };
    Session.prototype.loop = function () {
        return __awaiter(this, void 0, void 0, function () {
            var _this = this;
            return __generator(this, function (_a) {
                try {
                    setInterval(function () { return __awaiter(_this, void 0, void 0, function () {
                        var packet, data, id, ch;
                        return __generator(this, function (_a) {
                            switch (_a.label) {
                                case 0: return [4 /*yield*/, this.readPacket()];
                                case 1:
                                    packet = _a.sent();
                                    if (!(packet[0] == msgChannelOpen)) return [3 /*break*/, 3];
                                    return [4 /*yield*/, this.handleChannelOpen(packet)];
                                case 2:
                                    _a.sent();
                                    return [2 /*return*/];
                                case 3:
                                    data = new DataView(packet);
                                    id = data.getUint32(1);
                                    ch = this.getCh(id);
                                    if (ch === undefined) {
                                        throw "invalid channel (" + id + ") on op " + packet[0];
                                    }
                                    return [4 /*yield*/, ch.handlePacket(data)];
                                case 4:
                                    _a.sent();
                                    return [2 /*return*/];
                            }
                        });
                    }); }, 20);
                }
                finally { }
                return [2 /*return*/];
            });
        });
    };
    Session.prototype.getCh = function (id) {
        return this.channels[id];
    };
    Session.prototype.addCh = function (ch) {
        var _this = this;
        this.channels.forEach(function (v, i) {
            if (v === undefined) {
                _this.channels[i] = ch;
                return i;
            }
        });
        this.channels.push(ch);
        return this.channels.length - 1;
    };
    Session.prototype.rmCh = function (id) {
        this.channels[id] = undefined;
    };
    Session.prototype.accept = function () {
        return this.incoming.shift();
    };
    Session.prototype.close = function () {
        return this.conn.close();
    };
    return Session;
}());
exports.Session = Session;
var Channel = /** @class */ (function () {
    function Channel() {
    }
    Channel.prototype.ident = function () {
        return this.localId;
    };
    Channel.prototype.sendPacket = function (packet) {
        if (this.sentClose) {
            throw "EOF";
        }
        this.sentClose = (packet[0] === msgChannelClose);
        return this.session.conn.send(packet.buffer);
    };
    Channel.prototype.sendMessage = function (type, msg) {
        var data = new DataView(encode(type, msg));
        data.setUint32(1, this.remoteId);
        return this.sendPacket(new Uint8Array(data.buffer));
    };
    Channel.prototype.handlePacket = function (packet) {
        return __awaiter(this, void 0, void 0, function () {
            var fmsg, cmsg, amsg;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        if (packet.buffer[0] === msgChannelData) {
                            this.handleData(packet);
                            return [2 /*return*/];
                        }
                        if (!(packet.buffer[0] === msgChannelClose)) return [3 /*break*/, 3];
                        return [4 /*yield*/, this.sendMessage(msgChannelClose, {
                                peersID: this.remoteId
                            })];
                    case 1:
                        _a.sent();
                        this.session.rmCh(this.localId);
                        return [4 /*yield*/, this.handleClose()];
                    case 2:
                        _a.sent();
                        return [2 /*return*/];
                    case 3:
                        if (packet.buffer[0] === msgChannelEOF) {
                            // TODO
                            return [2 /*return*/];
                        }
                        if (!(packet.buffer[0] === msgChannelOpenFailure)) return [3 /*break*/, 5];
                        fmsg = decode(packet.buffer);
                        this.session.rmCh(fmsg.peersID);
                        return [4 /*yield*/, this.ready.push(false)];
                    case 4:
                        _a.sent();
                        return [2 /*return*/];
                    case 5:
                        if (!(packet.buffer[0] === msgChannelOpenConfirm)) return [3 /*break*/, 7];
                        cmsg = decode(packet.buffer);
                        if (cmsg.maxPacketSize < minPacketLength || cmsg.maxPacketSize > 1 << 31) {
                            throw "invalid max packet size";
                        }
                        this.remoteId = cmsg.myID;
                        this.maxRemotePayload = cmsg.maxPacketSize;
                        this.remoteWin += cmsg.myWindow;
                        return [4 /*yield*/, this.ready.push(true)];
                    case 6:
                        _a.sent();
                        return [2 /*return*/];
                    case 7:
                        if (packet.buffer[0] === msgChannelWindowAdjust) {
                            amsg = decode(packet.buffer);
                            this.remoteWin += amsg.additionalBytes;
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    Channel.prototype.handleData = function (packet) {
        return __awaiter(this, void 0, void 0, function () {
            var length, data;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        length = packet.getUint32(5);
                        if (length == 0) {
                            return [2 /*return*/];
                        }
                        if (length > this.maxIncomingPayload) {
                            throw "incoming packet exceeds maximum payload size";
                        }
                        data = packet.buffer.slice(9, length);
                        // TODO: check packet length
                        if (this.myWindow < length) {
                            throw "remot side wrote too much";
                        }
                        this.myWindow -= length;
                        return [4 /*yield*/, this.readBuf.push(data)];
                    case 1:
                        _a.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    Channel.prototype.adjustWindow = function (n) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                return [2 /*return*/];
            });
        });
    };
    Channel.prototype.recv = function () {
        return __awaiter(this, void 0, void 0, function () {
            var _a, _b, _c;
            return __generator(this, function (_d) {
                switch (_d.label) {
                    case 0:
                        _b = (_a = Promise).resolve;
                        _c = Uint8Array.bind;
                        return [4 /*yield*/, this.readBuf.shift()];
                    case 1: return [2 /*return*/, _b.apply(_a, [new (_c.apply(Uint8Array, [void 0, _d.sent()]))()])];
                }
            });
        });
    };
    Channel.prototype.send = function (buffer) {
        if (this.sentEOF) {
            return Promise.reject("EOF");
        }
        // TODO: use window
        var header = new DataView(new ArrayBuffer(9));
        header.setUint8(0, msgChannelData);
        header.setUint32(1, this.remoteId);
        header.setUint32(5, buffer.byteLength);
        var packet = new Uint8Array(9 + buffer.byteLength);
        packet.set(new Uint8Array(header.buffer), 0);
        packet.set(new Uint8Array(buffer), 9);
        return this.sendPacket(packet);
    };
    Channel.prototype.handleClose = function () {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.ready.close()];
                    case 1:
                        _a.sent();
                        this.sentClose = true;
                        return [2 /*return*/];
                }
            });
        });
    };
    Channel.prototype.close = function () {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.sendMessage(msgChannelClose, {
                            peersID: this.remoteId
                        })];
                    case 1:
                        _a.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    Channel.prototype.closeWrite = function () {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        this.sentEOF = true;
                        return [4 /*yield*/, this.sendMessage(msgChannelEOF, {
                                peersID: this.remoteId
                            })];
                    case 1:
                        _a.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    return Channel;
}());
exports.Channel = Channel;
function encode(type, obj) {
    switch (type) {
        case msgChannelClose:
            var data = new DataView(new ArrayBuffer(5));
            data.setUint8(0, type);
            data.setUint32(1, obj.peersID);
            return data.buffer;
        case msgChannelData:
            var datamsg = obj;
            var data = new DataView(new ArrayBuffer(9));
            data.setUint8(0, type);
            data.setUint32(1, datamsg.peersID);
            data.setUint32(5, datamsg.length);
            var buf = new Uint8Array(9 + datamsg.length);
            buf.set(new Uint8Array(data.buffer), 0);
            buf.set(datamsg.rest, 9);
            return buf.buffer;
        case msgChannelEOF:
            var data = new DataView(new ArrayBuffer(5));
            data.setUint8(0, type);
            data.setUint32(1, obj.peersID);
            return data.buffer;
        case msgChannelOpen:
            var data = new DataView(new ArrayBuffer(13));
            var openmsg = obj;
            data.setUint8(0, type);
            data.setUint32(1, openmsg.peersID);
            data.setUint32(5, openmsg.peersWindow);
            data.setUint32(9, openmsg.maxPacketSize);
            return data.buffer;
        case msgChannelOpenConfirm:
            var data = new DataView(new ArrayBuffer(17));
            var confirmmsg = obj;
            data.setUint8(0, type);
            data.setUint32(1, confirmmsg.peersID);
            data.setUint32(5, confirmmsg.myID);
            data.setUint32(9, confirmmsg.myWindow);
            data.setUint32(13, confirmmsg.maxPacketSize);
            return data.buffer;
        case msgChannelOpenFailure:
            var data = new DataView(new ArrayBuffer(5));
            data.setUint8(0, type);
            data.setUint32(1, obj.peersID);
            return data.buffer;
        case msgChannelWindowAdjust:
            var data = new DataView(new ArrayBuffer(9));
            var adjustmsg = obj;
            data.setUint8(0, type);
            data.setUint32(1, adjustmsg.peersID);
            data.setUint32(5, adjustmsg.additionalBytes);
            return data.buffer;
        default:
            throw "unknown type";
    }
}
function decode(packet) {
    switch (packet[0]) {
        case msgChannelClose:
            var data = new DataView(new ArrayBuffer(5));
            var closeMsg = {
                peersID: data.getUint32(1)
            };
            return closeMsg;
        case msgChannelData:
            var data = new DataView(new ArrayBuffer(9));
            var dataLength = data.getUint32(5);
            var dataMsg = {
                peersID: data.getUint32(1),
                length: dataLength,
                rest: new Uint8Array(dataLength)
            };
            dataMsg.rest.set(new Uint8Array(data.buffer.slice(9)));
            return dataMsg;
        case msgChannelEOF:
            var data = new DataView(new ArrayBuffer(5));
            var eofMsg = {
                peersID: data.getUint32(1)
            };
            return eofMsg;
        case msgChannelOpen:
            var data = new DataView(new ArrayBuffer(13));
            var openMsg = {
                peersID: data.getUint32(1),
                peersWindow: data.getUint32(5),
                maxPacketSize: data.getUint32(9)
            };
            return openMsg;
        case msgChannelOpenConfirm:
            var data = new DataView(new ArrayBuffer(17));
            var confirmMsg = {
                peersID: data.getUint32(1),
                myID: data.getUint32(5),
                myWindow: data.getUint32(9),
                maxPacketSize: data.getUint32(13)
            };
            return confirmMsg;
        case msgChannelOpenFailure:
            var data = new DataView(new ArrayBuffer(5));
            var failureMsg = {
                peersID: data.getUint32(1)
            };
            return failureMsg;
        case msgChannelWindowAdjust:
            var data = new DataView(new ArrayBuffer(9));
            var adjustMsg = {
                peersID: data.getUint32(1),
                additionalBytes: data.getUint32(5)
            };
            return adjustMsg;
        default:
            throw "unknown type";
    }
}
