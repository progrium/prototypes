package main

import (
	"log"

	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
)

const addr = "localhost:4242"

type Field struct {
	Type       string  `msgpack:"type"`
	Name       string  `msgpack:"name"`
	Value      string  `msgpack:"value"`
	Expression *string `msgpack:"expression"`
	Fields     []Field `msgpack:"fields"`
}

type Button struct {
	Name string `msgpack:"name"`
}

type Component struct {
	Name    string   `msgpack:"name"`
	Fields  []Field  `msgpack:"fields"`
	Buttons []Button `msgpack:"buttons"`
}

type Node struct {
	Active     bool        `msgpack:"active"`
	Components []Component `msgpack:"components"`
}

type Project struct {
	Name string `msgpack:"name"`
	Path string `msgpack:"path"`
}

type State struct {
	Projects       []Project       `msgpack:"projects"`
	CurrentProject string          `msgpack:"currentProject"`
	Components     []string        `msgpack:"components"`
	Hierarchy      []string        `msgpack:"hierarchy"`
	Nodes          map[string]Node `msgpack:"nodes"`
}

func main() {

	state := State{
		Projects: []Project{
			{Name: "project1", Path: "/Project1"},
			{Name: "project2", Path: "/Project2"},
		},
		CurrentProject: "project1",
		Hierarchy: []string{
			"/NodeA",
			"/NodeB",
			"/NodeC",
			"/NodeC/Child",
		},
		Components: []string{
			"foo.Component",
			"twilio.Client",
			"osx.Menu",
			"progrium.Jeff",
		},
		Nodes: map[string]Node{
			"/NodeA": Node{
				Active: true,
				Components: []Component{
					{
						Name: "foo.Component",
						Fields: []Field{
							{Type: "string", Name: "FooString", Value: "foobar"},
						},
					},
				},
			},
			"/NodeB": Node{
				Active:     true,
				Components: []Component{},
			},
			"/NodeC": Node{
				Active:     true,
				Components: []Component{},
			},
			"/NodeC/Child": Node{
				Active:     true,
				Components: []Component{},
			},
		},
	}

	// filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
	// 	if path == ".." {
	// 		return nil
	// 	}
	// 	if strings.HasPrefix(path, "../node_modules") {
	// 		return nil
	// 	}
	// 	state.Files = append(state.Files, strings.TrimPrefix(path, ".."))
	// 	return nil
	// })

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
	api.HandleFunc("selectProject", func(r qrpc.Responder, c *qrpc.Call) {
		var name string
		err := c.Decode(&name)
		if err != nil {
			r.Return(err)
			return
		}
		state.CurrentProject = name
		sendState()
		r.Return(nil)
	})
	// api.HandleFunc("readFile", func(r qrpc.Responder, c *qrpc.Call) {
	// 	var path string
	// 	err := c.Decode(&path)
	// 	if err != nil {
	// 		r.Return(err)
	// 		return
	// 	}
	// 	bytes, err := ioutil.ReadFile(".." + path)
	// 	if err != nil {
	// 		r.Return(err)
	// 		return
	// 	}
	// 	r.Return(string(bytes))
	// })

	// start server with api
	server := &qrpc.Server{}
	l, err := mux.ListenWebsocket(addr)
	if err != nil {
		panic(err)
	}
	log.Println("websocket server listening at", addr)
	log.Fatal(server.Serve(l, api))
}
