package main

import (
	"fmt"
	"log"

	"github.com/progrium/prototypes/qrpc"
	"github.com/progrium/prototypes/qrpc/transport"
)

const busAddr = "localhost:4242"

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// define api
	api := qrpc.NewAPI()
	api.HandleFunc("echo", func(r qrpc.Responder, c *qrpc.Call) {
		var msg string
		err := c.Decode(&msg)
		if err != nil {
			r.Return(err)
			return
		}
		log.Println("got echo")
		r.Return(msg)
	})

	// connect backend to bus
	sess, err := transport.DialTCP(busAddr)
	if err != nil {
		panic(err)
	}
	backend := &qrpc.Client{Session: sess, API: api}
	err = backend.Call("register", []string{"echo"}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("serving...")
	backend.ServeAPI()
}
