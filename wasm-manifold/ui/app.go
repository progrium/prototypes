package ui

import (
	"syscall/js"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-manifold/manifold"
)

func init() {
	webui.Register(App{})
}

type App struct {
	vecty.Core

	TreeView  *TreeView  `vecty:"ref"`
	Inspector *Inspector `vecty:"ref"`
}

func (c *App) OnReset(e *vecty.Event) {
	js.Global().Get("localStorage").Call("setItem", "tree_nodes", "{}")
	js.Global().Get("location").Call("reload")
}

func (c *App) OnAdd(e *vecty.Event) {
	ret := js.Global().Call("prompt", "New object")
	if ret == js.Null() {
		return
	}
	c.TreeView.CreateNode(manifold.NewNode(ret.String()))
}

func (c *App) OnSelect(n *manifold.Node) {
	c.Inspector.Node = n
	vecty.Rerender(c.Inspector)
}

func (c *App) Mount() {

}

func (c *App) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
