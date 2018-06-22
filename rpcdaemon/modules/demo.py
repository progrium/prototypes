#!/usr/local/bin/python3 -u

import duplex
import duplex.ws4py

rpc = duplex.RPC("json", async=False)

def echo(ch):
    obj, _ = ch.recv()
    print("echo: {0}".format(obj))
    ch.send(obj)

rpc.register("echo", echo)

print("connecting to 8000...")
conn = duplex.ws4py.Client("ws://localhost:8000")
peer = rpc.handshake(conn)
print(peer.call("register", ["echo"]))
peer.route()

