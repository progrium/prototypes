package webui

import (
	"strings"

	"github.com/gowasm/vecty"
	reflected "github.com/progrium/prototypes/go-reflected"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
	vectytemplate "github.com/progrium/prototypes/go-vtemplate/vecty"
	"github.com/progrium/prototypes/wasm/assets"
)

var registry = make(map[string]reflected.Value)

func Register(c interface{}) {
	v := reflected.ValueOf(c)
	registry[strings.ToLower(v.Type().Name())] = v
}

func registryElements() map[string]vtemplate.CustomElement {
	elements := make(map[string]vtemplate.CustomElement)
	for n, _ := range registry {
		elements[n] = nil
	}
	return elements
}

func registryComponents() []interface{} {
	var c []interface{}
	for _, v := range registry {
		c = append(c, v.Interface())
	}
	return c
}

func Render(v interface{}) vecty.ComponentOrHTML {
	tmpl, err := assets.FindTemplate(2)
	if err != nil {
		return vecty.Text(err.Error())
	}
	p := &vtemplate.Parser{
		Directives:     vectytemplate.BuiltinDirectives(),
		Evaluator:      nil,
		CustomElements: registryElements(),
	}
	n, err := p.Parse(tmpl, v)
	if err != nil {
		return vecty.Text(err.Error())
	}
	return vectytemplate.Render(n, v, registryComponents())
}
