package vecty

import (
	"strings"

	"github.com/gowasm/vecty"
	reflected "github.com/progrium/prototypes/go-reflected"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
	"github.com/progrium/prototypes/go-vtemplate/html"
)

func BuiltinDirectives() map[string]vtemplate.Directive {
	return map[string]vtemplate.Directive{
		"if":   html.IfDirective{},
		"bind": html.BindDirective{},
		"for":  html.ForDirective{},
		"on":   OnDirective{},
	}
}

type OnDirective struct{}

func (d OnDirective) Apply(b vtemplate.Binding) error {
	rcvr := b.Node.Data
	method := b.Expression
	b.Node.Attrs[b.Argument] = &vecty.EventListener{
		Name: b.Argument,
		Listener: func(e *vecty.Event) {
			rcvr.Call(method, e)
		},
	}
	return nil
}

func Render(n *vtemplate.Node, v interface{}, c []interface{}) vecty.ComponentOrHTML {
	switch n.Type {
	case vtemplate.CustomNode:
		for _, proto := range c {
			comType := reflected.ValueOf(proto).Type()
			comName := strings.ToLower(comType.Name())
			if comName == n.Name {
				com := reflected.New(comType)
				for k, v := range n.Attrs {
					for _, prop := range comType.Fields() {
						// TODO: only set for fields tagged as prop
						if k == strings.ToLower(prop) {
							com.Set(prop, v)
						}
					}
				}
				if len(n.Children) > 0 {
					slotFields := comType.FieldsTagged("vecty", "slot")
					if len(slotFields) > 0 {
						// TODO: named slots
						var slot []vecty.ComponentOrHTML
						for _, child := range n.Children {
							slot = append(slot, Render(child, v, c))
						}
						com.Set(slotFields[0], vecty.List(slot))
					}
				}
				ch := com.Interface().(vecty.ComponentOrHTML)
				return ch
			}
		}
		// no component found
		return vecty.Text("")
	case vtemplate.ElementNode:
		if n.Name == "slot" {
			rv := reflected.ValueOf(v)
			slotFields := rv.Type().FieldsTagged("vecty", "slot")
			if len(slotFields) > 0 {
				return rv.Get(slotFields[0]).Interface().(vecty.ComponentOrHTML)
			}
			return vecty.Text("")
		}
		var applyers []vecty.Applyer
		for key, val := range n.Attrs {
			switch tval := val.(type) {
			case *vecty.EventListener:
				applyers = append(applyers, tval)
			default:
				applyers = append(applyers, vecty.Attribute(key, val))
			}
		}
		children := []vecty.MarkupOrChild{vecty.Markup(applyers...)}
		for _, child := range n.Children {
			res := Render(child, v, c)
			switch r := res.(type) {
			case vecty.List:
				for _, rr := range r {
					children = append(children, rr)
				}
			default:
				children = append(children, res)
			}
		}
		return vecty.Tag(n.Name, children...)
	case vtemplate.TextNode:
		return vecty.Text(n.Text)
	default:
		return vecty.Text("")
	}
}
