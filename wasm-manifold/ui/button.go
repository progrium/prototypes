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
	Slot    vecty.List         `vecty:"slot"`
}

func (c *Button) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
