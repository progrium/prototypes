package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Window{})
}

type Tab struct {
	Title string
}

type Window struct {
	vecty.Core

	WindowStyle   string
	TitlebarStyle string
	CloseboxStyle string
	ContentStyle  string

	Tabs []Tab

	Selected int
}

func (c *Window) HandleTabSelect(idx int) {
	c.Selected = idx
	vecty.Rerender(c)
}

func (c *Window) Render() vecty.ComponentOrHTML {
	c.Tabs = []Tab{{Title: "Foo"}, {Title: "Bar"}}
	c.WindowStyle = InlineStyle(map[string]string{
		"position":      "relative",
		"overflow":      "hidden",
		"display":       "flex",
		"flexDirection": "column",
		"flexGrow":      "1",
		"borderRadius":  "3",
		"border":        "1px solid rgba(0,0,0,0.1)",
	})
	c.ContentStyle = InlineStyle(map[string]string{
		"backgroundColor": "rgb(83,83,83)",
		"color":           "rgb(180,180,180)",
		"fontSize":        "8pt",
		"flexGrow":        "1",
		"display":         "flex",
		"flexDirection":   "column",
		"border":          "1px solid rgba(255,255,255,0.025)",
	})
	c.TitlebarStyle = InlineStyle(map[string]string{
		"flexShrink":      "0",
		"position":        "relative",
		"height":          "24",
		"minHeight":       "22",
		"boxSizing":       "border-box",
		"display":         "flex",
		"backgroundColor": "rgb(66,66,66)",
	})
	c.CloseboxStyle = InlineStyle(map[string]string{
		"boxSizing":          "border-box",
		"color":              "grey",
		"width":              "12",
		"margin":             "0 6px",
		"height":             "100%",
		"textAlign":          "center",
		"fontSize":           "10px",
		"flexShrink":         "0",
		"flexGrow":           "0",
		"backgroundImage":    "url(/icons/hamburger.svg)",
		"backgroundSize":     "contain",
		"backgroundRepeat":   "no-repeat",
		"backgroundPosition": "center",
		"opacity":            "0.5",
	})
	return webui.Render(c)
}
