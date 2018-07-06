package main

import (
	"flag"
	"fmt"

	"github.com/progrium/prototypes/qrpc"
	"github.com/progrium/prototypes/qrpc/transport"
)

const addr = "localhost:4242"

func main() {
	sess, err := transport.DialTCP(addr)
	if err != nil {
		panic(err)
	}
	client := &qrpc.Client{Session: sess}

	flag.Parse()

	var resp string
	err = client.Call(flag.Arg(0), "Hello", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %#v\n", resp)
}
