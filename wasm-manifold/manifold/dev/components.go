package dev

import "github.com/progrium/prototypes/wasm-manifold/manifold"

func init() {
	manifold.RegisterComponent(DemoComponent{})
}

type DemoComponent struct {
	StringValue string
	IntValue    int
	BoolValue   bool
}
