package ui

import (
	"github.com/gowasm/vecty"
	reflected "github.com/progrium/prototypes/go-reflected"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(PropertySet{})
}

type PropertySet struct {
	vecty.Core

	Value interface{} `vecty:"prop"`

	Name   string
	Fields []string
}

func (c *PropertySet) Render() vecty.ComponentOrHTML {
	rf := reflected.ValueOf(c.Value)
	c.Name = rf.Type().Name()
	c.Fields = rf.Type().Fields()
	return webui.Render(c)
}
