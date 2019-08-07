package workspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/progrium/prototypes/editor/server/manifold"
	"github.com/spf13/afero"
)

type Node struct {
	Children   []string
	Active     bool
	Name       string
	ID         string
	Components []*manifold.Component
}

func saveNode(parentPath string, node *manifold.Node, fs afero.Fs) error {
	n := &Node{
		Active:     node.Active,
		Name:       node.Name,
		ID:         node.ID,
		Components: node.Components,
	}
	for _, child := range node.Children {
		n.Children = append(n.Children, child.Name)
	}
	buf, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	if err := afero.WriteFile(fs, "node.json", buf, 0644); err != nil {
		return err
	}
	node.Dir = path.Join(parentPath, node.Name)
	for _, child := range node.Children {
		if err := fs.MkdirAll(child.Name, 0755); err != nil {
			return err
		}
		if err := saveNode(node.Dir, child, afero.NewBasePathFs(fs, child.Name)); err != nil {
			return err
		}
	}
	return nil
}

func loadNode(parentPath string, nodeDir afero.Fs) (*manifold.Node, error) {
	b, err := afero.ReadFile(nodeDir, "node.json")
	if err != nil {
		return nil, err
	}
	var n Node
	err = json.Unmarshal(b, &n)
	// TODO: Handle missing components
	if err != nil {
		return nil, err
	}
	node := &manifold.Node{
		ID:         n.ID,
		Name:       n.Name,
		Active:     n.Active,
		Components: n.Components,
		Dir:        path.Join(parentPath, n.Name),
	}
	for _, childName := range n.Children {
		child, err := loadNode(node.Dir, afero.NewBasePathFs(nodeDir, childName))
		if err != nil {
			return nil, err
		}
		node.Append(child)
	}
	node.TempInflate()
	return node, nil
}

func SaveHierarchy(root *manifold.Node) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := path.Join(wd, "hierarchy")
	os.MkdirAll(dir, 0755)
	fs := afero.NewBasePathFs(afero.NewOsFs(), dir)
	return saveNode(dir, root, fs)
}

func LoadHierarchy() (*manifold.Node, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir := path.Join(wd, "hierarchy")
	fs := afero.NewBasePathFs(afero.NewOsFs(), dir)
	if ok, err := afero.Exists(fs, "node.json"); !ok || err != nil {
		return NewHierarchy(), nil
	}
	return loadNode(dir, fs)
}

func NewHierarchy() *manifold.Node {
	n := manifold.NewNode("")
	n.Append(manifold.NewNode("NodeA"))
	n.Append(manifold.NewNode("NodeB"))
	return n
}

func DelegateSource(id string) ([]byte, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir := path.Join(wd, "workspace/delegates")
	fs := afero.NewBasePathFs(afero.NewOsFs(), dir)
	if ok, _ := afero.Exists(fs, path.Join(id, "delegate.go")); !ok {
		return nil, fmt.Errorf("no delegate found")
	}
	return afero.ReadFile(fs, path.Join(id, "delegate.go"))
}

func WriteDelegate(id string, contents []byte) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := path.Join(wd, "workspace/delegates")
	fs := afero.NewBasePathFs(afero.NewOsFs(), dir)
	return afero.WriteFile(fs, path.Join(id, "delegate.go"), contents, 0644)
}

func CreateDelegate(node *manifold.Node) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := path.Join(wd, "workspace/delegates")
	fs := afero.NewBasePathFs(afero.NewOsFs(), dir)

	fs.MkdirAll(node.ID, 0755)
	if ok, _ := afero.Exists(fs, path.Join(node.ID, "delegate.go")); !ok {
		delegate := fmt.Sprintf(`package node

import "github.com/progrium/prototypes/editor/server/manifold"

func init() {
	manifold.RegisterDelegate(&Delegate{}, "%s")
}

type Delegate struct{}
`, node.ID)
		if err := afero.WriteFile(fs, path.Join(node.ID, "delegate.go"), []byte(delegate), 0644); err != nil {
			return err
		}
	}

	var imports []string
	ids := append(manifold.RegisteredDelegates(), node.ID)
	for _, id := range ids {
		imports = append(imports, fmt.Sprintf(` _ "github.com/progrium/prototypes/editor/server/workspace/delegates/%s"`, id))
	}
	delegates := fmt.Sprintf(`package delegates
import (
%s
)
`, strings.Join(imports, "\n"))
	err = afero.WriteFile(fs, "delegates.go", []byte(delegates), 0644)
	if err != nil {
		return err
	}

	return nil
}
