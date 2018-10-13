package ui

import (
	"fmt"
	"syscall/js"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(App{})
}

type App struct {
	vecty.Core

	TreeView *TreeView `vecty:"ref"`
}

func (c *App) OnReset(e *vecty.Event) {
	js.Global().Get("localStorage").Call("setItem", "jstree", "[]")
	js.Global().Get("location").Call("reload")
}

func (c *App) OnAdd(e *vecty.Event) {
	var name = js.Global().Call("prompt", "New object").String()
	c.TreeView.CreateNode(map[string]interface{}{
		"text": name,
		"obj": map[string]interface{}{
			"_": map[string]interface{}{
				"name": name,
			},
		},
	})
}

func (c *App) OnSelect() {
	fmt.Println(c.TreeView.SelectedNode().Get("text").String())
}

func (c *App) OnChange() {
	fmt.Println("CHANGE")
}

func (c *App) Mount() {

}

func (c *App) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
