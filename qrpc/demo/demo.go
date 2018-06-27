package main

import (
	"fmt"
	"log"

	"github.com/progrium/prototypes/qrpc"
)

// TODO: API for more symmetrical use
// TODO: Maybe try another transport
// TODO: reflection stuff. register objects, not functions

const addr = "localhost:4242"

func main() {
	go qrpc.ListenAndServe(addr, qrpc.Destinations{
		{
			Path: "echo",
			Handler: func(args interface{}) (interface{}, error) {
				log.Println("echo called")
				m := args.(map[interface{}]interface{})
				m["Req"] = false
				m["Resp"] = true
				return m, nil
			},
		},
	})

	peer, err := qrpc.DialPeer(addr)
	if err != nil {
		panic(err)
	}
	defer peer.Close()
	var reply Demo
	args := Demo{
		Message: "Hello world 2",
		Req:     true,
	}
	fmt.Printf("%#v\n", args)
	err = peer.Call("echo", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", reply)
}

type Demo struct {
	Message string
	Req     bool
	Resp    bool
}
