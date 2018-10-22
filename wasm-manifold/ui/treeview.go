package ui

import (
	"fmt"

	"github.com/gopherjs/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-manifold/manifold"
)

func init() {
	webui.Register(TreeView{})
}

type TreeView struct {
	vecty.Core

	root *manifold.Node

	selectedID string
	mounted    bool
}

func (c *TreeView) CreateNode(node *manifold.Node) {
	c.root.Append(node)
	vecty.Rerender(c)
	c.Save()
}

func (c *TreeView) Save() {
	nodes, err := manifold.Marshal(c.root)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(nodes))
	js.Global().Get("localStorage").Call("setItem", "tree_nodes", string(nodes))
}

func (c *TreeView) insertNode(node *manifold.Node, parent string, idx int) {
	p := c.root.FindID(parent)
	if p == nil {
		panic("parent id not found")
	}
	p.Insert(idx, node)
}

func (c *TreeView) deleteNode(id, parent string) {
	p := c.root.FindID(parent)
	if p == nil {
		panic("parent id not found " + parent)
	}
	p.Remove(id)
}

func (c *TreeView) Mount() {
	n, err := manifold.Unmarshal([]byte(js.Global().Get("localStorage").Call("getItem", "tree_nodes").String()))
	if err != nil {
		panic(err)
	}
	c.root = n
	var components []interface{}
	for _, n := range manifold.RegisteredComponents() {
		components = append(components, n)
	}
	js.Global().Set("components", components)
	js.Global().Call("$", "#jstree").Call("jstree", map[string]interface{}{
		"core": map[string]interface{}{
			"themes": map[string]interface{}{
				"dots": false,
			},
			"animation":      0,
			"check_callback": true,
		},
		"plugins": []interface{}{"contextmenu", "types", "dnd"},
		"types": map[string]interface{}{
			"default": map[string]interface{}{
				"icon": "file icon",
			},
		},
		"contextmenu": map[string]interface{}{
			"items": js.Global().Get("contextMenu"),
		},
	})
	js.Global().Set("addComponent", js.NewCallback(func(args []js.Value) {
		id := args[0].String()
		com := args[1].String()
		node := c.root.FindID(id)
		if node == nil {
			panic("node not found")
		}
		node.Components.Append(manifold.NewComponent(com))
	}))
	// js.Global().Call("$", "#jstree").Call("on", "create_node.jstree", js.NewCallback(func(args []js.Value) {
	// 	js.Global().Get("console").Call("log", args[0], args[1])
	// }))
	js.Global().Call("$", "#jstree").Call("on", "move_node.jstree", js.NewCallback(func(args []js.Value) {
		p := args[1].Get("parent").String()
		op := args[1].Get("old_parent").String()
		pos := args[1].Get("position").Int()
		id := args[1].Get("node").Get("id").String()
		node := c.root.FindID(id)
		if node == nil {
			panic("node not found " + id)
		}
		c.deleteNode(id, op)
		c.insertNode(node, p, pos)
		vecty.Rerender(c)
		c.Save()
	}))
	js.Global().Call("$", "#jstree").Call("on", "delete_node.jstree", js.NewCallback(func(args []js.Value) {
		id := args[1].Get("node").Get("id").String()
		parent := args[1].Get("node").Get("parent").String()
		c.deleteNode(id, parent)
		vecty.Rerender(c)
		c.Save()
	}))
	js.Global().Call("$", "#jstree").Call("on", "rename_node.jstree", js.NewCallback(func(args []js.Value) {
		id := args[1].Get("node").Get("id").String()
		node := c.root.FindID(id)
		if node == nil {
			panic("node not found " + id)
		}
		node.SetName(args[1].Get("node").Get("text").String())
		vecty.Rerender(c)
		c.Save()
	}))
	// c.Element.Node().Call("on", "activate_node.jstree",
	// 	c.selectedID = args[1].Get("node").Get("id").String()
	// 	if c.OnSelect != nil {
	// 		c.OnSelect()
	// 	}
	// }))
	c.Refresh()
	c.mounted = true
}

func (c *TreeView) Refresh() {
	js.Global().Call("$", "#jstree").Call("jstree", true).Get("settings").Get("core").Set("data", c.root.TreeNode()["children"])
	js.Global().Call("$", "#jstree").Call("jstree", true).Call("refresh")
}

func (c *TreeView) Render() vecty.ComponentOrHTML {
	if c.mounted {
		c.Refresh()
	}
	return nil
}

type TreeNode struct {
	ID       string
	Text     string
	Icon     string
	Opened   bool
	Disabled bool
	Selected bool
	Loaded   bool
	Children []*TreeNode
}

func (n *TreeNode) Map() map[string]interface{} {
	var nodes []interface{}
	for _, c := range n.Children {
		nodes = append(nodes, c.Map())
	}
	return map[string]interface{}{
		"id":   n.ID,
		"text": n.Text,
		"icon": n.Icon,
		"state": map[string]interface{}{
			"opened":   n.Opened,
			"disabled": n.Disabled,
			"selected": n.Selected,
			"loaded":   n.Loaded,
		},
		"children": nodes,
	}
}
