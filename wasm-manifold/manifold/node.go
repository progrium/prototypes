package manifold

import (
	"encoding/json"
	"fmt"
	"time"
)

type Node struct {
	Components
	parent   *Node
	children []*Node
	active   bool
	name     string
	id       string
}

func Marshal(n *Node) ([]byte, error) {
	return json.Marshal(n.export())
}

func Unmarshal(b []byte) (*Node, error) {
	var node map[string]interface{}
	err := json.Unmarshal(b, &node)
	node["id"] = "#"
	return load(node), err
}

func load(data map[string]interface{}) *Node {
	n := &Node{
		id:     data["id"].(string),
		name:   data["name"].(string),
		active: data["active"].(bool),
	}
	if data["children"] == nil {
		return n
	}
	for _, child := range data["children"].([]interface{}) {
		c := load(child.(map[string]interface{}))
		c.parent = n
		n.children = append(n.children, c)
	}
	return n
}

func NewNode(name string) *Node {
	return &Node{
		name:   name,
		active: true,
		id:     fmt.Sprintf("n%d", time.Now().Unix()), // TODO: replace me
	}
}

func (n *Node) export() map[string]interface{} {
	var nodes []interface{}
	for _, c := range n.children {
		nodes = append(nodes, c.export())
	}
	return map[string]interface{}{
		"id":         n.id,
		"name":       n.name,
		"active":     n.active,
		"children":   nodes,
		"components": n.Components.components,
	}
}

func (n *Node) TreeNode() map[string]interface{} {
	var nodes []interface{}
	for _, c := range n.children {
		nodes = append(nodes, c.TreeNode())
	}
	return map[string]interface{}{
		"id":       n.id,
		"text":     n.name,
		"children": nodes,
	}
}

func (n *Node) Find(name string) *Node {
	for _, child := range n.children {
		if child.name == name {
			return child
		}
	}
	for _, child := range n.children {
		if res := child.Find(name); res != nil {
			return res
		}
	}
	return nil
}

func (n *Node) FindID(id string) *Node {
	if n.id == id {
		return n
	}
	for _, child := range n.children {
		if child.id == id {
			return child
		}
	}
	for _, child := range n.children {
		if res := child.FindID(id); res != nil {
			return res
		}
	}
	return nil
}

func (n *Node) RemoveAt(idx int) *Node {
	node := n.children[idx]
	n.children = append(n.children[:idx], n.children[idx+1:]...)
	return node
}

func (n *Node) Remove(id string) *Node {
	for idx, child := range n.children {
		if child.id == id {
			return n.RemoveAt(idx)
		}
	}
	return nil
}

func (n *Node) Insert(idx int, node *Node) {
	node.parent = n
	n.children = append(n.children[:idx], append([]*Node{node}, n.children[idx:]...)...)
}

func (n *Node) Append(node *Node) {
	node.parent = n
	n.children = append(n.children, node)
}

func (n *Node) SetActive(active bool) {
	n.active = active
}

func (n *Node) Active() bool {
	return n.active
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) SetName(name string) {
	n.name = name
}

func (n *Node) Children() []*Node {
	return n.children
}

func (n *Node) Save() error {
	return nil
}
