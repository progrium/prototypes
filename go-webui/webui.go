package webui

import (
	"fmt"
	"io"

	"github.com/gowasm/vecty"
	reflected "github.com/progrium/prototypes/go-reflected"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
	vectytemplate "github.com/progrium/prototypes/go-vtemplate/vecty"
)

var registry = make(map[string]reflected.Value)

type TemplateFinder func(int) (io.Reader, string, error)

var FindTemplate TemplateFinder

func Register(c interface{}) {
	v := reflected.ValueOf(c)
	registry[v.Type().Name()] = v
}

func registryElements() map[string]vtemplate.CustomElement {
	elements := make(map[string]vtemplate.CustomElement)
	for n, _ := range registry {
		// check if component implements CustomElement
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
	tmpl, filepath, err := FindTemplate(2)
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
		return vecty.Text(fmt.Sprintf("%s: %s", filepath, err.Error()))
	}
	return vectytemplate.Render(n, v, registryComponents())
}
