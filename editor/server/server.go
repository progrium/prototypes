package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
)

const addr = "localhost:4242"

type State struct {
	Message string
	Counter int
	Files   []string
}

func main() {

	state := State{
		Message: "Hello world",
		Counter: 1,
		Files:   []string{},
	}

	filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		if path == ".." {
			return nil
		}
		if strings.HasPrefix(path, "../node_modules") {
			return nil
		}
		state.Files = append(state.Files, strings.TrimPrefix(path, ".."))
		return nil
	})

	clients := make(map[qrpc.Caller]string)

	sendState := func() {
		for client, callback := range clients {
			err := client.Call(callback, state, nil)
			if err != nil {
				delete(clients, client)
				log.Println(err)
			}
		}
	}

	// define api
	api := qrpc.NewAPI()
	api.HandleFunc("subscribe", func(r qrpc.Responder, c *qrpc.Call) {
		clients[c.Caller] = "state"
		sendState()
	})
	api.HandleFunc("readFile", func(r qrpc.Responder, c *qrpc.Call) {
		var path string
		err := c.Decode(&path)
		if err != nil {
			r.Return(err)
			return
		}
		bytes, err := ioutil.ReadFile(".." + path)
		if err != nil {
			r.Return(err)
			return
		}
		r.Return(string(bytes))
	})

	// start server with api
	server := &qrpc.Server{}
	l, err := mux.ListenWebsocket(addr)
	if err != nil {
		panic(err)
	}
	log.Println("websocket server listening at", addr)
	log.Fatal(server.Serve(l, api))
}
