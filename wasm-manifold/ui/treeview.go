package ui

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(TreeView{})
}

type TreeView struct {
	vecty.Core

	OnSelect func()
	OnChange func()
	Element  *vecty.HTML `vecty:"ref"`

	selectedID string
}

func (c *TreeView) CreateNode(obj interface{}) {
	c.Element.Node().Get("jstree").Call("create_node", js.Null(), obj)
}

func (c *TreeView) Mount() {
	c.Element.Node().Call("on", "activate_node.jstree", js.NewCallback(func(args []js.Value) {
		c.selectedID = args[1].Get("node").Get("id").String()
		if c.OnSelect != nil {
			c.OnSelect()
		}
	}))
	if c.OnChange == nil {
		return
	}
	for _, event := range []string{
		"create_node.jstree",
		"move_node.jstree",
		"delete_node.jstree",
		"rename_node.jstree",
	} {
		c.Element.Node().Call("on", event, js.NewCallback(func(args []js.Value) {
			c.OnChange()
		}))
	}
}

func (c *TreeView) SelectedNode() js.Value {
	return c.Element.Node().Get("jstree").Call("get_node", c.selectedID)
}

func (c *TreeView) Render() vecty.ComponentOrHTML {
	return webui.Render(c)
}
