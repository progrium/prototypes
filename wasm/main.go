package main

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/wasm/app"
)

func main() {
	c := make(chan struct{}, 0)
	vecty.SetTitle("Markdown Demo")
	vecty.RenderBody(&app.PageView{
		Input: `# Markdown Example

This is a live editor, try editing the Markdown on the right of the page.
`,
	})
	<-c
}
