package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-manifold/manifold"
)

func init() {
	webui.Register(Inspector{})
}

type Inspector struct {
	vecty.Core

	Root         func() *manifold.Node `vecty:"prop"`
	Node         *manifold.Node        `vecty:"prop"`
	OnChange     func()                `vecty:"prop"`
	OnNodeChange func()                `vecty:"prop"`
}

func (c *Inspector) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
