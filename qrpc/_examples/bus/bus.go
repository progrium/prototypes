package main

import (
	"fmt"
	"log"

	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
	"github.com/progrium/prototypes/qrpc/bus"
)

const busAddr = "localhost:4242"

func main() {
	// start the bus
	server := &qrpc.Server{}
	l, err := mux.ListenTCP(busAddr)
	if err != nil {
		panic(err)
	}
	go func() {
		log.Fatal(server.Serve(l, bus.NewBus()))
	}()

	// make and connect a backend
	api := qrpc.NewAPI()
	handler, err := qrpc.ExportFunc(func() string {
		return "Hello world"
	})
	if err != nil {
		panic(err)
	}
	api.Handle("hello", handler)
	backendSess, err := mux.DialTCP(busAddr)
	if err != nil {
		panic(err)
	}
	backend := &qrpc.Client{Session: backendSess, API: api}
	go backend.ServeAPI()
	err = backend.Call("register", []string{"hello"}, nil)
	if err != nil {
		panic(err)
	}

	// make and connect a frontend
	frontendSess, err := mux.DialTCP(busAddr)
	if err != nil {
		panic(err)
	}
	frontend := &qrpc.Client{Session: frontendSess}
	var resp string
	err = frontend.Call("hello", nil, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %#v\n", resp)
}
