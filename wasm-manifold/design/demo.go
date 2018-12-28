package main

import (
	"fmt"

	"github.com/progrium/prototypes/wasm-manifold/manifold"
)

// data bindings

func main() {
	var root = manifold.NewNode("DemoApp", func(n *manifold.Node) {
		n.Children = []*manifold.Node{
			manifold.NewNode("Server", func(n *manifold.Node) error {
				var handler web.Handler
				if !n.Reference("/DemoApp/Handler", &handler) {
					return fmt.Errorf("no handler at /DemoApp/Handler")
				} 
				n.Components = manifold.Components(
					&web.Component{
						ListenAddr: "localhost:8080",
						Handlers: []web.Handler{
							handler,
						},
					},
				)
			}),
			manifold.NewNode("Handler", func(n *manifold.Node) error {
				var s = struct{
					Web *web.Component `bind:"/DemoApp/Server"`
				}{}
				n.LoadRefs(&s)
				n.AppendComponent(&Handler{})
			})
		}
	})

	var root = &manifold.Node{
		Name: "DemoApp",
		Children: []*manifold.Node{
			{
				Name: "Server",
				OnAwake: func(n *manifold.Node) error {
					var handler web.Handler
					if !n.Reference("/DemoApp/Handler", &handler) {
						return fmt.Errorf("no handler at /DemoApp/Handler")
					}
					n.AppendComponent(&web.Component{
						ListenAddr: "localhost:8080",
						Handlers: []web.Handler{
							handler,
						},
					})
				},
			},
			{
				Name:       "Handler",
				Components: manifold.Components(),
			},
		},
	}

	println(root.Name)
}
