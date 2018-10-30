package ui

import (
	"fmt"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Slider{})
}

type Slider struct {
	vecty.Core

	SliderStyle    string
	ContainerStyle string

	Value string
}

func (c *Slider) OnChange(e *vecty.Event) {
	fmt.Println(e.Get("target").Get("value").String())
}

func (c *Slider) Render() vecty.ComponentOrHTML {
	c.SliderStyle = InlineStyle(map[string]string{
		"width": "100%",
	})
	c.ContainerStyle = InlineStyle(map[string]string{
		"position":      "relative",
		"padding":       "1",
		"minWidth":      "128",
		"minHeight":     "128",
		"flexGrow":      "1",
		"display":       "flex",
		"flexDirection": "column",
	})
	return webui.Render(c)
}
