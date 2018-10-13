package manifold

type Node struct {
	Parent     *Node
	Components []interface{}
	Children   []*Node
}
