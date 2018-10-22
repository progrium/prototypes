//go:generate sh -c "GOARCH=wasm GOOS=js go build -o ../static/app.wasm ."
package main

import (
	"github.com/gowasm/vecty"
	webui "github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-manifold/assets"
	"github.com/progrium/prototypes/wasm-manifold/ui"

	_ "github.com/progrium/prototypes/wasm-manifold/manifold/dev"
)

func main() {
	c := make(chan struct{}, 0)
	webui.FindTemplate = assets.FindTemplate
	vecty.SetTitle("Manifold UI Prototype")
	vecty.RenderBody(&ui.App{})
	<-c
}
