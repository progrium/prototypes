package main

import (
	"fmt"
	"log"

	"github.com/progrium/prototypes/qrpc"
)

// TODO: Maybe try another transport
// TODO: reflection stuff. register objects, not functions

const addr = "localhost:4242"

type EchoMessage struct {
	Message string
	Req     bool
	Resp    bool
}

func simpleEcho(message string) (string, error) {
	log.Println("simple echo called")
	return message, nil
}

func main() {
	done := make(chan bool)
	api := qrpc.NewAPI()
	api.Handle("simple-echo", qrpc.ExportFunc(simpleEcho))
	api.HandleFunc("echo-client", func(r qrpc.Responder, c *qrpc.Call) {
		log.Println("echo-client called")
		var msg EchoMessage
		err := c.Decode(&msg)
		if err != nil {
			r.Return(err)
			return
		}
		msg.Req = false
		msg.Resp = true
		r.Return(msg)
	})
	api.HandleFunc("echo-server", func(r qrpc.Responder, c *qrpc.Call) {
		log.Println("echo-server called")
		var msg EchoMessage
		err := c.Decode(&msg)
		if err != nil {
			r.Return(err)
			return
		}
		msg.Req = false
		msg.Resp = true
		r.Return(msg)

		req := &EchoMessage{
			Message: "hello client",
			Req:     true,
		}
		fmt.Printf("req: %#v\n", req)
		var resp EchoMessage
		err = c.Caller.Call("echo-client", req, &resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("resp: %#v\n", resp)
		done <- true
	})

	server := &qrpc.Server{}
	go server.ListenAndServe(addr, api)

	client, err := qrpc.Dial(addr, api)
	if err != nil {
		panic(err)
	}
	go client.ServeAPI()
	req := &EchoMessage{
		Message: "hello server",
		Req:     true,
	}
	fmt.Printf("req: %#v\n", req)
	var resp EchoMessage
	err = client.Call("echo-server", req, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %#v\n", resp)
	<-done
}
