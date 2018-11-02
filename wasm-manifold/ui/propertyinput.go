package ui

import (
	"reflect"
	"strconv"

	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/event"
	"github.com/gowasm/vecty/prop"
	reflected "github.com/progrium/prototypes/go-reflected"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(PropertyInput{})
}

type PropertyInput struct {
	vecty.Core

	Value    reflected.Value `vecty:"prop"`
	Field    string          `vecty:"prop"`
	Object   reflected.Value `vecty:"prop"`
	OnChange func()          `vecty:"prop"`
}

func (c *PropertyInput) Render() vecty.ComponentOrHTML {
	switch c.Value.Kind() {
	case reflect.Bool:
		return elem.Input(vecty.Markup(
			prop.Type(prop.TypeCheckbox),
			prop.Checked(c.Value.Bool()),
			event.Input(func(e *vecty.Event) {
				c.Object.Set(c.Field, e.Target.Get("checked").Bool())
				if c.OnChange != nil {
					c.OnChange()
				}
			}),
		))
	case reflect.Int:
		return elem.Input(vecty.Markup(
			prop.Type(prop.TypeNumber),
			prop.Value(strconv.Itoa(c.Value.Int())),
			event.Input(func(e *vecty.Event) {
				c.Object.Set(c.Field, e.Target.Get("valueAsNumber").Int())
				if c.OnChange != nil {
					c.OnChange()
				}
			}),
		))
	case reflect.Float64:
		return elem.Input(vecty.Markup(
			prop.Type(prop.TypeNumber),
			prop.Value(strconv.FormatFloat(c.Value.Float(), 'f', -1, 64)),
			event.Input(func(e *vecty.Event) {
				c.Object.Set(c.Field, e.Target.Get("valueAsNumber").Float())
				if c.OnChange != nil {
					c.OnChange()
				}
			}),
		))
	default:
		return elem.Input(vecty.Markup(
			prop.Type(prop.TypeText),
			prop.Value(c.Value.String()),
			vecty.Style("border", "1px solid lightgray"),
			event.Input(func(e *vecty.Event) {
				c.Object.Set(c.Field, e.Target.Get("value").String())
				if c.OnChange != nil {
					c.OnChange()
				}
			}),
		))
	}
}
