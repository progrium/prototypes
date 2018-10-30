package ui

import (
	"reflect"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(WrapperWrapped{})
}

func copy(vv interface{}) interface{} {
	v := reflect.ValueOf(vv)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("copy: value must be pointer to struct, found " + reflect.TypeOf(vv).String())
	}
	cpy := reflect.New(v.Elem().Type())
	cpy.Elem().Set(v.Elem())
	return cpy.Interface()
}

type WrapperWrapped struct {
	vecty.Core
	Children vecty.List `vecty:"slot"`
}

func (c *WrapperWrapped) Render() vecty.ComponentOrHTML {
	// var newList vecty.List
	// for _, child := range c.Children {
	// 	newList = append(newList, copy(child).(vecty.ComponentOrHTML))
	// }
	// c.Children = newList
	return webui.Render(c)
}
