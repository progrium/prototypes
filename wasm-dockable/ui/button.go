package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Button{})
}

type Button struct {
	vecty.Core

	OnClick func(*vecty.Event) `vecty:"prop"`

	Style string
}

func (c *Button) Render() vecty.ComponentOrHTML {
	c.Style = InlineStyle(map[string]string{
		"padding":  "8",
		"minWidth": "128",
	})
	return webui.Render(c)
}
