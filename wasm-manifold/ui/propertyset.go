package ui

import (
	"github.com/gowasm/vecty"
	reflected "github.com/progrium/prototypes/go-reflected"
	"github.com/progrium/prototypes/go-webui"
	"github.com/progrium/prototypes/wasm-manifold/manifold"
)

func init() {
	webui.Register(PropertySet{})
}

type fieldView struct {
	Name  string
	Value reflected.Value
}

type PropertySet struct {
	vecty.Core

	Root     func() *manifold.Node `vecty:"prop"`
	Value    interface{}           `vecty:"prop"`
	OnChange func()                `vecty:"prop"`

	Object reflected.Value
	Name   string
	Fields []fieldView
}

func (c *PropertySet) Render() vecty.ComponentOrHTML {
	c.Object = reflected.ValueOf(c.Value)
	c.Name = c.Object.Type().Name()
	c.Fields = nil
	hidden := c.Object.Type().FieldsTagged("inspector", "hide")
	for _, field := range c.Object.Type().Fields() {
		for _, n := range hidden {
			if n == field {
				goto skip
			}
		}
		c.Fields = append(c.Fields, fieldView{
			Name:  field,
			Value: c.Object.Get(field),
		})
	skip:
	}
	return webui.Render(c)
}
