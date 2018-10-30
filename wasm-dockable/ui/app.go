package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(App{})
}

type App struct {
	vecty.Core
}

func (c *App) OnClick(e *vecty.Event) {
	vecty.Rerender(c)
}

func (c *App) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
