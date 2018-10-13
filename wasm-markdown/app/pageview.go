package app

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(PageView{})
}

type PageView struct {
	vecty.Core

	Input string
}

func (p *PageView) OnTextAreaChange(e *vecty.Event) {
	// When input is typed into the textarea, update the local
	// component state and rerender.
	p.Input = e.Target.Get("value").String()
	vecty.Rerender(p)
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return webui.Render(p)
}
