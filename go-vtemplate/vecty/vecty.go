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

func renderComponent(com reflected.Value, n *vtemplate.Node, v interface{}, c []interface{}) vecty.ComponentOrHTML {
	refName := ""
	for k, attr := range n.Attrs {
		if k == "ref" {
			refName = attr.(string)
			continue
		}
		for _, prop := range com.Type().Fields() {
			// TODO: only set for fields tagged as prop
			if k == prop {
				if com.Type().FieldType(prop).Kind() == reflect.Int {
					i, err := strconv.Atoi(attr.(string))
					if err != nil {
						panic(err)
					}
					com.Set(prop, i)
				} else {
					com.Set(prop, attr)
				}
			}
		}
	}
	if len(n.Children) > 0 {
		slotFields := com.Type().FieldsTagged("vecty", "slot")
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

func parseAttrs(attrs map[string]interface{}) (string, []vecty.Applyer) {
	refName := ""
	var applyers []vecty.Applyer
	for key, val := range attrs {
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
	return refName, applyers
}

func Render(n *vtemplate.Node, v interface{}, c []interface{}) vecty.ComponentOrHTML {
	switch n.Type {
	case vtemplate.CustomNode:
		for _, proto := range c {
			comType := reflected.ValueOf(proto).Type()
			if comType.Name() == n.Name {
				com := reflected.New(comType)
				return renderComponent(com, n, v, c)
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
		if n.Name == "node" {
			switch is := n.Attrs["is"].(type) {
			case *vecty.HTML:
				delete(n.Attrs, "is")
				_, applyers := parseAttrs(n.Attrs)
				for _, applyer := range applyers {
					applyer.Apply(is)
				}
				return is
			default:
				delete(n.Attrs, "is")
				return renderComponent(reflected.ValueOf(is), n, v, c)
			}
		}
		refName, applyers := parseAttrs(n.Attrs)
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
