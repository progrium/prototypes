//go:generate sh -c "GOARCH=wasm GOOS=js go build -o ../static/app.wasm ."
package main

import (
	"github.com/gowasm/vecty"
	webui "github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-dockable/assets"
	"github.com/progrium/prototypes/wasm-dockable/ui"
)

func main() {
	c := make(chan struct{}, 0)
	webui.FindTemplate = assets.FindTemplate
	vecty.SetTitle("Dockable Port")
	vecty.RenderBody(&ui.App{})
	<-c
}
