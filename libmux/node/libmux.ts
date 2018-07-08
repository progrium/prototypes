import * as ffi from "ffi";
import * as ref from "ref";

import * as ArrayType from "ref-array";
import * as Struct from "ref-struct";
import { PerformanceObserver } from "perf_hooks";

let ByteArray = ArrayType(ref.types.uint8);
let GoString = Struct({
  p: "string",
  n: "longlong"
});

function goStr(str: string): any {
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
  ChannelClose: ["int", ["int"]],
});

function lookupErr(id: number): string {
  var buf = ByteArray(1<<8);
  var n = libmux.Error(id*-1, buf, buf.length);
  return buf.buffer.slice(0,n).toString('ascii');
}

function handle(reject, name, retHandler) {
  return (err, retcode) => {
    if (err) {
      reject("ffi: "+err);
      return;
    }
    if (retcode < 0) {
      reject(name+"["+(retcode*-1)+"]: "+lookupErr(retcode));
      return;
    }
    retHandler(retcode);
  };
}

export function ListenTCP(addr: string): Promise<Listener> {
  return new Promise((resolve, reject) => {
    libmux.ListenTCP.async(goStr(addr), handle(reject, "ListenTCP", (retcode) => {
      if (retcode === 0) {
        resolve();
        return;
      }
      resolve(new Listener(retcode));
    }));
  });
}

export function DialTCP(addr: string): Promise<Session> {
  return new Promise((resolve, reject) => {
    libmux.DialTCP.async(goStr(addr), handle(reject, "DialTCP", (retcode) => {
      if (retcode === 0) {
        resolve();
        return;
      }
      resolve(new Session(retcode));
    }));
  });
}

export function ListenWebsocket(addr: string): Promise<Listener> {
  return new Promise((resolve, reject) => {
    libmux.ListenWebsocket.async(goStr(addr), handle(reject, "ListenWebsocket", (retcode) => {
      if (retcode === 0) {
        resolve();
        return;
      }
      resolve(new Listener(retcode));
    }));
  });
}

export function DialWebsocket(addr: string): Promise<Session> {
  return new Promise((resolve, reject) => {
    libmux.DialWebsocket.async(goStr(addr), handle(reject, "DialWebsocket", (retcode) => {
      if (retcode === 0) {
        resolve();
        return;
      }
      resolve(new Session(retcode));
    }));
  });
}

class Listener {
  id: number;
  closed: boolean;

  constructor(id: number) {
    this.id = id;
    this.closed = false;
    process.once('SIGINT', (code) => {
      if (!this.closed) 
        libmux.ListenerClose(id);
    });
  }

  accept(): Promise<Session> {
    if (this.closed) return new Promise(r => r());
    return new Promise((resolve, reject) => {
      libmux.ListenerAccept.async(this.id, handle(reject, "ListenerAccept", (retcode) => {
        if (retcode === 0) {
          resolve();
          return;
        }
        resolve(new Session(retcode));
      }));
    });
  }

  close(): Promise<void> {
    if (this.closed) return Promise.resolve();
    return new Promise((resolve, reject) => {
      libmux.ListenerClose.async(this.id, handle(reject, "ListenerClose", () => {
        this.closed = true;
        resolve();
      }));
    });
  }
}

class Session {
  id: number;
  closed: boolean;

  constructor(id: number) {
    this.id = id;
    this.closed = false;
  }

  open(): Promise<Channel> {
    if (this.closed) return new Promise(r => r());
    return new Promise((resolve, reject) => {
      libmux.SessionOpen.async(this.id, handle(reject, "SessionOpen", (retcode) => {
        if (retcode === 0) {
          resolve();
          return;
        }
        resolve(new Channel(retcode));
      }));
    });
  }

  accept(): Promise<Channel> {
    if (this.closed) return new Promise(r => r());
    return new Promise((resolve, reject) => {
      libmux.SessionAccept.async(this.id, handle(reject, "SessionAccept", (retcode) => {
        if (retcode === 0) {
          resolve();
          return;
        }
        resolve(new Channel(retcode));
      }));
    });
  }

  close(): Promise<void> {
    if (this.closed) return Promise.resolve();
    return new Promise((resolve, reject) => {
      libmux.SessionClose.async(this.id, handle(reject, "SessionClose", () => {
        this.closed = true;
        resolve();
      }));
    });
  }
}

class Channel {
  id: number;
  closed: boolean;

  constructor(id: number) {
    this.id = id;
    this.closed = false;
  }

  read(len: number): Promise<Buffer> {
    if (this.closed) return new Promise(r => r());
    return new Promise((resolve, reject) => {
      var buffer = ByteArray(len);
      libmux.ChannelRead.async(this.id, buffer, buffer.length, handle(reject, "ChannelRead", (retcode) => {
        if (retcode === 0) {
          this.closed = true;
          resolve();
          return;
        }
        resolve(buffer.buffer.slice(0, retcode));
      }));
    });
  }

  write(buf: Buffer): Promise<number> {
    if (this.closed) return new Promise(r => r());
    return new Promise((resolve, reject) => {
      var buffer = ByteArray(buf);
      libmux.ChannelWrite.async(this.id, buffer, buffer.length, handle(reject, "ChannelWrite", (retcode) => resolve(retcode)));
    });
  }

  close(): Promise<void> {
    if (this.closed) return Promise.resolve();
    return new Promise((resolve, reject) => {
      libmux.ChannelClose.async(this.id, handle(reject, "ChannelClose", () => {
        this.closed = true;
        resolve();
      }));
    });
  }
}