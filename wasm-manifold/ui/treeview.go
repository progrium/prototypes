package ui

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gopherjs/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(TreeView{})
}

type TreeView struct {
	vecty.Core

	nodes   []*TreeNode
	nodeIDs map[string]*TreeNode

	selectedID string
	mounted    bool
}

func (c *TreeView) CreateNode(node TreeNode) {
	if node.ID == "" {
		node.ID = fmt.Sprintf("n%d", time.Now().Unix())
	}
	node.Loaded = true
	c.nodes = append(c.nodes, &node)
	c.nodeIDs[node.ID] = &node
	vecty.Rerender(c)
	c.Save()
}

func (c *TreeView) Save() {
	nodes, err := json.Marshal(c.nodes)
	if err != nil {
		panic(err)
	}
	nodeIDs, err := json.Marshal(c.nodeIDs)
	if err != nil {
		panic(err)
	}
	js.Global().Get("localStorage").Call("setItem", "tree_nodes", string(nodes))
	js.Global().Get("localStorage").Call("setItem", "tree_nodeIDs", string(nodeIDs))
}

func (c *TreeView) insertNode(node *TreeNode, parent string, idx int) {
	if parent == "#" {
		c.nodes = append(c.nodes[:idx], append([]*TreeNode{node}, c.nodes[idx:]...)...)
	} else {
		c.nodeIDs[parent].Children = append(c.nodeIDs[parent].Children[:idx], append([]*TreeNode{node}, c.nodeIDs[parent].Children[idx:]...)...)
	}
	c.nodeIDs[node.ID] = node
}

func (c *TreeView) deleteNode(id, parent string) {
	var children []*TreeNode
	if parent == "#" {
		children = c.nodes
	} else {
		children = c.nodeIDs[parent].Children
	}
	for idx := 0; idx < len(children); idx++ {
		if children[idx].ID == id {
			if parent == "#" {
				c.nodes = append(children[:idx], children[idx+1:]...)
			} else {
				c.nodeIDs[parent].Children = append(children[:idx], children[idx+1:]...)
			}
			break
		}
	}
	delete(c.nodeIDs, id)
}

func (c *TreeView) Mount() {
	err := json.Unmarshal([]byte(js.Global().Get("localStorage").Call("getItem", "tree_nodes").String()), &(c.nodes))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(js.Global().Get("localStorage").Call("getItem", "tree_nodeIDs").String()), &(c.nodeIDs))
	if err != nil {
		panic(err)
	}
	if c.nodeIDs == nil {
		c.nodeIDs = make(map[string]*TreeNode)
	}
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
			"demo": map[string]interface{}{
				"icon": "braille icon",
			},
		},
	})
	// js.Global().Call("$", "#jstree").Call("on", "create_node.jstree", js.NewCallback(func(args []js.Value) {
	// 	js.Global().Get("console").Call("log", args[0], args[1])
	// }))
	js.Global().Call("$", "#jstree").Call("on", "move_node.jstree", js.NewCallback(func(args []js.Value) {
		p := args[1].Get("parent").String()
		op := args[1].Get("old_parent").String()
		pos := args[1].Get("position").Int()
		id := args[1].Get("node").Get("id").String()
		node := c.nodeIDs[id]
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
		c.nodeIDs[id].Text = args[1].Get("node").Get("text").String()
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
	var nodes []interface{}
	for _, n := range c.nodes {
		nodes = append(nodes, n.Map())
	}
	js.Global().Call("$", "#jstree").Call("jstree", true).Get("settings").Get("core").Set("data", nodes)
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
