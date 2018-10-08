package vecty

import (
	"strings"

	"github.com/gowasm/vecty"
	reflected "github.com/progrium/prototypes/go-reflected"
)

func NodeToVecty(n *Node, v interface{}, r ComponentRegistry) vecty.ComponentOrHTML {
	switch n.Type {
	case ElementNode:
		if r != nil {
			for com := range r.Components() {
				comType := reflected.ValueOf(com).Type()
				comName := comType.Name()
				if comName == n.Data {
					c := reflected.New(comType)
					for k, v := range n.Attrs {
						for _, prop := range comType.Fields() {
							if k == strings.ToLower(prop) {
								c.Set(prop, v)
							}
						}
					}
					ch := c.Interface().(vecty.ComponentOrHTML)
					return ch
				}
			}
		}
		var applyers []vecty.Applyer
		for key, val := range n.Attrs {
			// still need to do this because its vecty specific (why directive is commented out above)
			if strings.HasPrefix(key, "v-on:") {
				method := val
				key := strings.Replace(key, "v-on:", "", 1)
				rcvr := reflected.ValueOf(v)
				applyers = append(applyers, &vecty.EventListener{
					Name: key,
					Listener: func(e *vecty.Event) {
						rcvr.Call(method, e)
					},
				})
				continue
			}
			applyers = append(applyers, vecty.Attribute(key, val))
		}
		children := []vecty.MarkupOrChild{vecty.Markup(applyers...)}
		for _, c := range n.Children {
			children = append(children, NodeToVecty(c, v, r))
		}
		return vecty.Tag(n.Data, children...)
	case TextNode:
		return vecty.Text(n.Data)
	default:
		return vecty.Text("")
	}
}
