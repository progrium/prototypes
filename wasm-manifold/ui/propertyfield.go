package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(PropertyField{})
}

type PropertyField struct {
	vecty.Core

	Name  string
	Value string
}

func (c *PropertyField) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
