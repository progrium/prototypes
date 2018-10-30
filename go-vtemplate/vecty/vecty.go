package vecty

import (
	"reflect"
	"strconv"

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
	callback := b.Expression
	listener := func(e *vecty.Event) {
		rcvr.Get(callback).Invoke(e)
	}
	if rcvr.Type().HasMethod(callback) {
		listener = func(e *vecty.Event) {
			rcvr.Call(callback, e)
		}
	}
	b.Node.Attrs[b.Argument] = &vecty.EventListener{
		Name:     b.Argument,
		Listener: listener,
	}
	return nil
}

func Render(n *vtemplate.Node, v interface{}, c []interface{}) vecty.ComponentOrHTML {
	switch n.Type {
	case vtemplate.CustomNode:
		for _, proto := range c {
			comType := reflected.ValueOf(proto).Type()
			if comType.Name() == n.Name {
				refName := ""
				com := reflected.New(comType)
				for k, v := range n.Attrs {
					if k == "ref" {
						refName = v.(string)
						continue
					}
					for _, prop := range comType.Fields() {
						// TODO: only set for fields tagged as prop
						if k == prop {
							if comType.FieldType(prop).Kind() == reflect.Int {
								i, err := strconv.Atoi(v.(string))
								if err != nil {
									panic(err)
								}
								com.Set(prop, i)
							} else {
								com.Set(prop, v)
							}
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
				ch := com.Interface()
				if refName != "" {
					rv := reflected.ValueOf(v)
					refFields := rv.Type().FieldsTagged("vecty", "ref")
					for _, f := range refFields {
						if f == refName {
							rv.Set(refName, ch)
							break
						}
					}
				}
				return ch.(vecty.ComponentOrHTML)
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
		refName := ""
		var applyers []vecty.Applyer
		for key, val := range n.Attrs {
			if key == "ref" {
				refName = val.(string)
				continue
			}
			switch tval := val.(type) {
			case *vecty.EventListener:
				applyers = append(applyers, tval)
			default:
				if key == "key" {
					applyers = append(applyers, vecty.Key(val))
				} else {
					applyers = append(applyers, vecty.Attribute(key, val))
				}
			}
		}
		var children []vecty.ComponentOrHTML
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
		if n.Name == "fragment" {
			return vecty.List(children)
		}
		markupAndChildren := []vecty.MarkupOrChild{vecty.Markup(applyers...)}
		for _, child := range children {
			markupAndChildren = append(markupAndChildren, child)
		}
		el := vecty.Tag(n.Name, markupAndChildren...)
		if refName != "" {
			rv := reflected.ValueOf(v)
			refFields := rv.Type().FieldsTagged("vecty", "ref")
			for _, f := range refFields {
				if f == refName {
					rv.Set(refName, el)
					break
				}
			}
		}
		return el
	case vtemplate.TextNode:
		return vecty.Text(n.Text)
	default:
		return vecty.Text("")
	}
}
