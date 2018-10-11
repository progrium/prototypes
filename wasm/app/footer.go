package app

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/wasm/pkg/webui"
)

func init() {
	webui.Register(Footer{})
}

type Footer struct {
	vecty.Core

	Copyright string     `vecty:"prop"`
	Children  vecty.List `vecty:"slot"`
}

func (m *Footer) Render() vecty.ComponentOrHTML {
	return webui.Render(m)
}
