package manifold

import (
	"fmt"
	"time"
)

type Node struct {
	ComponentSet

	parent   *Node
	Children []*Node
	Active   bool
	Name     string
	ID       string
}

func NewNode(name string) *Node {
	return &Node{
		Name:   name,
		Active: true,
		ID:     fmt.Sprintf("n%d", time.Now().Unix()), // TODO: replace me
	}
}

// func (n *Node) export() map[string]interface{} {
// 	var nodes []interface{}
// 	for _, c := range n.children {
// 		nodes = append(nodes, c.export())
// 	}
// 	return map[string]interface{}{
// 		"id":         n.id,
// 		"name":       n.name,
// 		"active":     n.active,
// 		"children":   nodes,
// 		"components": n.Components.components,
// 	}
// }

func (n *Node) TreeNode() map[string]interface{} {
	var nodes []interface{}
	for _, c := range n.Children {
		nodes = append(nodes, c.TreeNode())
	}
	return map[string]interface{}{
		"id":       n.ID,
		"text":     n.Name,
		"children": nodes,
	}
}

func (n *Node) Find(name string) *Node {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}
	for _, child := range n.Children {
		if res := child.Find(name); res != nil {
			return res
		}
	}
	return nil
}

func (n *Node) FindID(id string) *Node {
	if n.ID == id {
		return n
	}
	for _, child := range n.Children {
		if child.ID == id {
			return child
		}
	}
	for _, child := range n.Children {
		if res := child.FindID(id); res != nil {
			return res
		}
	}
	return nil
}

func (n *Node) RemoveAt(idx int) *Node {
	node := n.Children[idx]
	n.Children = append(n.Children[:idx], n.Children[idx+1:]...)
	return node
}

func (n *Node) Remove(id string) *Node {
	for idx, child := range n.Children {
		if child.ID == id {
			return n.RemoveAt(idx)
		}
	}
	return nil
}

func (n *Node) Insert(idx int, node *Node) {
	node.parent = n
	n.Children = append(n.Children[:idx], append([]*Node{node}, n.Children[idx:]...)...)
}

func (n *Node) Append(node *Node) {
	node.parent = n
	n.Children = append(n.Children, node)
}

func (n *Node) Save() error {
	return nil
}
