import * as ffi from "ffi";
import * as ref from "ref";

import * as ArrayType from "ref-array";
import * as Struct from "ref-struct";

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

var libmux = ffi.Library("../libmux", {
  Error: [GoString, ["int"]],

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

export function ListenTCP(addr: string): Promise<Listener> {
  return new Promise((resolve, reject) => {
    libmux.ListenTCP.async(goStr(addr), (err, ret) => {
      if (err) {
        reject(err);
        return;
      }
      if (ret < 0) {
        reject(libmux.Error(ret).p);
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

export function DialTCP(addr: string): Promise<Session> {
  return new Promise((resolve, reject) => {
    libmux.DialTCP.async(goStr(addr), (err, ret) => {
      if (err) {
        reject(err);
        return;
      }
      if (ret < 0) {
        reject(libmux.Error(ret).p);
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

class Listener {
  id: number;

  constructor(id: number) {
    this.id = id;
    process.once('SIGINT', function (code) {
      libmux.ListenerClose(id);
    });
  }

  Accept(): Promise<Session> {
    return new Promise((resolve, reject) => {
      libmux.ListenerAccept.async(this.id, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
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

  Close(): Promise<void> {
    return new Promise((resolve, reject) => {
      libmux.ListenerClose.async(this.id, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        resolve();
      });
    });
  }
}

class Session {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  Open(): Promise<Channel> {
    return new Promise((resolve, reject) => {
      libmux.SessionOpen.async(this.id, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        if (ret === 0) {
          resolve();
          return;
        }
        resolve(new Channel(ret));
      });
    });
  }

  Accept(): Promise<Channel> {
    return new Promise((resolve, reject) => {
      libmux.SessionAccept.async(this.id, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        if (ret === 0) {
          resolve();
          return;
        }
        resolve(new Channel(ret));
      });
    });
  }

  Close(): Promise<void> {
    return new Promise((resolve, reject) => {
      libmux.SessionClose.async(this.id, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        resolve();
      });
    });
  }
}

class Channel {
  id: number;

  constructor(id: number) {
    this.id = id;
  }

  Read(len: number): Promise<Buffer> {
    return new Promise((resolve, reject) => {
      var buffer = ByteArray(len);
      libmux.ChannelRead.async(this.id, buffer, buffer.length, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        resolve(buffer.buffer);
      })
    });
  }

  Write(buf: Buffer): Promise<number> {
    return new Promise((resolve, reject) => {
      var buffer = ByteArray(buf);
      libmux.ChannelWrite.async(this.id, buffer, buffer.length, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        resolve(ret);
      })
    });
  }

  Close(): Promise<void> {
    return new Promise((resolve, reject) => {
      libmux.ChannelClose.async(this.id, (err, ret) => {
        if (err) {
          reject(err);
          return;
        }
        if (ret < 0) {
          reject(libmux.Error(ret).p);
          return;
        }
        resolve();
      });
    });
  }
}