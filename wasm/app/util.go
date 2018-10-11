package app

import (
	"bytes"

	"github.com/gowasm/vecty"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
	vvecty "github.com/progrium/prototypes/go-vtemplate/vecty"
)

func render(tmpl string, v interface{}) vecty.ComponentOrHTML {
	p := &vtemplate.Parser{
		Directives: vvecty.BuiltinDirectives(),
		Evaluator:  nil,
		CustomElements: map[string]vtemplate.CustomElement{
			"footer":   nil,
			"markdown": nil,
		},
	}
	n, err := p.Parse(bytes.NewBufferString(tmpl), v)
	if err != nil {
		return vecty.Text(err.Error())
	}
	return vvecty.Render(n, v, []interface{}{Footer{}, Markdown{}})
}
