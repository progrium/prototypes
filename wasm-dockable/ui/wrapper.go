package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Wrapper{})
}

type Wrapper struct {
	vecty.Core

	Children vecty.List `vecty:"slot"`
}

func (c *Wrapper) Render() vecty.ComponentOrHTML {
	children := c.Children[:]
	var newChildren vecty.List
	for _, child := range children {
		newChildren = append(newChildren, &WrapperWrapped{
			Children: vecty.List{child},
		})
	}
	c.Children = newChildren
	return webui.Render(c)
}
