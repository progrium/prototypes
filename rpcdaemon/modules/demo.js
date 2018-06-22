#!/usr/local/bin/node

var ws = require("nodejs-websocket")
var duplex = require("duplex").duplex

var rpc = new duplex.RPC(duplex.JSON)

rpc.register("upper", function(ch) {
  ch.onrecv = function(err, obj) {
    console.log("upper: "+obj);
    ch.send(obj.toUpperCase())
  }
})


console.log("connecting to 8000...")
const conn = ws.connect("ws://localhost:8000/");
conn.on("connect", function() {
    const wrapped = duplex.wrap["nodejs-websocket"](conn);
    // conn.keepalive = function() {
    //     if (conn.readyState === conn.OPEN) {
    //     conn.send("");
    //     return setTimeout(conn.keepalive, 15000);
    //     }
    // };
    rpc.handshake(wrapped, function(peer) {
        
        peer.call("register", ["upper"], function(err) {
            if (err !== null) {
                console.log(err)
            }
        })

    });
});