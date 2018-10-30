package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(WindowTab{})
}

type WindowTab struct {
	vecty.Core

	KeyIndex int       `vecty:"prop"`
	Title    string    `vecty:"prop"`
	Index    int       `vecty:"prop"`
	Selected bool      `vecty:"prop"`
	OnClick  func(int) `vecty:"prop"`

	TabStyle string
	Hovering bool
}

func (c *WindowTab) Key() interface{} {
	return c.KeyIndex
}

func (c *WindowTab) HandleMouseOver(e *vecty.Event) {
	e.Call("preventDefault")
	c.Hovering = true
	vecty.Rerender(c)
}
func (c *WindowTab) HandleMouseOut(e *vecty.Event) {
	e.Call("preventDefault")
	c.Hovering = false
	vecty.Rerender(c)
}

func (c *WindowTab) HandleClick(e *vecty.Event) {
	c.OnClick(c.Index)
}

func (c *WindowTab) Render() vecty.ComponentOrHTML {
	var dynamicStyle map[string]string
	if c.Selected {
		dynamicStyle = map[string]string{
			"color":           "rgb(200,200,200)",
			"backgroundColor": "rgb(83,83,83)",
			"border":          "1px solid rgba(255,255,255,0.025)",
			"borderBottom":    "0",
		}
	} else {
		dynamicStyle = map[string]string{
			"color":           "grey",
			"backgroundColor": "transparent",
			"cursor":          "auto",
			"border":          "1px solid transparent",
		}
		if c.Hovering {
			dynamicStyle["backgroundColor"] = "rgba(83,83,83,0.5)"
			dynamicStyle["cursor"] = "pointer"
		}
	}
	c.TabStyle = InlineStyle(map[string]string{
		"position":     "relative",
		"height":       "calc(100% + 1px)",
		"padding":      "6px 10px",
		"boxSizing":    "border-box",
		"overflow":     "hidden",
		"whiteSpace":   "nowrap",
		"fontSize":     "9px",
		"fontFamily":   "verdana",
		"fontWeight":   "bold",
		"borderRadius": "3px 3px 0 0",
		"marginRight":  "1",
	}, dynamicStyle)
	return webui.Render(c)
}
