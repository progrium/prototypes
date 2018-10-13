package main

import (
	"github.com/gowasm/vecty"
	webui "github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-markdown/app"
	"github.com/progrium/prototypes/wasm-markdown/assets"
)

func main() {
	c := make(chan struct{}, 0)
	webui.FindTemplate = assets.FindTemplate
	vecty.SetTitle("Markdown Demo")
	vecty.RenderBody(&app.PageView{
		Input: `# Markdown Example

This is a live editor, try editing the Markdown on the right of the page.
`,
	})
	<-c
}
