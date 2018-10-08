package html

import (
	"fmt"
	"io"

	reflected "github.com/progrium/prototypes/go-reflected"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
)

func BuiltinDirectives() map[string]vtemplate.Directive {
	return map[string]vtemplate.Directive{
		"if":   IfDirective{},
		"bind": BindDirective{},
		"for":  ForDirective{},
	}
}

type IfDirective struct{}

func (d IfDirective) Apply(b vtemplate.Binding) error {
	if b.Value == reflected.Undefined() || !b.Value.Bool() {
		b.Node.Type = vtemplate.NullNode
	}
	return nil
}

type BindDirective struct{}

func (d BindDirective) Apply(b vtemplate.Binding) error {
	if b.Value == reflected.Undefined() {
		return nil
	}
	b.Node.Attrs[b.Argument] = b.Value.Interface()
	return nil
}

type ForDirective struct{}

func (d ForDirective) Apply(b vtemplate.Binding) error {
	if b.IterVar == "" {
		return fmt.Errorf("v-for: invalid expression")
	}
	// create reflected map and copy data into it.
	// this lets us add values to it
	nv := reflected.ValueOf(map[string]interface{}{})
	for _, key := range b.Node.Data.Props() {
		nv.Set(key, b.Node.Data.Get(key).Interface())
	}
	b.Node.Children = nil
	for _, v := range b.Value.Iter() {
		nv.Set(b.IterVar, v.Interface())
		for c := b.Node.Html.FirstChild; c != nil; c = c.NextSibling {
			cn, err := b.Parser.ParseNode(c, nv)
			if err != nil {
				return err
			}
			if cn != nil {
				b.Node.Children = append(b.Node.Children, cn)
			}
		}
	}
	return nil
}

func Render(w io.Writer, n *vtemplate.Node) error {
	switch {
	case n.Type == vtemplate.ElementNode || n.Type == vtemplate.CustomNode:
		attrs := ""
		for k, v := range n.Attrs {
			attrs = fmt.Sprintf("%s %s=\"%s\"", attrs, k, v)
		}
		fmt.Fprintf(w, "<%s%s>\n", n.Name, attrs)
		indenter := newIndentWriter(w, []byte("  "))
		for _, child := range n.Children {
			if err := Render(indenter, child); err != nil {
				return err
			}
		}
		fmt.Fprintf(w, "</%s>\n", n.Name)
	case n.Type == vtemplate.TextNode:
		fmt.Fprintln(w, n.Text)
	default:
		// nothing?
	}
	return nil
}

// Writer indents each line of its input.
type indentWriter struct {
	w   io.Writer
	bol bool
	pre [][]byte
	sel int
	off int
}

// newIndentWriter makes a new write filter that indents the input
// lines. Each line is prefixed in order with the corresponding
// element of pre. If there are more lines than elements, the last
// element of pre is repeated for each subsequent line.
func newIndentWriter(w io.Writer, pre ...[]byte) io.Writer {
	return &indentWriter{
		w:   w,
		pre: pre,
		bol: true,
	}
}

// The only errors returned are from the underlying indentWriter.
func (w *indentWriter) Write(p []byte) (n int, err error) {
	for _, c := range p {
		if w.bol {
			var i int
			i, err = w.w.Write(w.pre[w.sel][w.off:])
			w.off += i
			if err != nil {
				return n, err
			}
		}
		_, err = w.w.Write([]byte{c})
		if err != nil {
			return n, err
		}
		n++
		w.bol = c == '\n'
		if w.bol {
			w.off = 0
			if w.sel < len(w.pre)-1 {
				w.sel++
			}
		}
	}
	return n, nil
}
