package ui

import (
	"syscall/js"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(PanelWrapper{})
}

type PanelWrapper struct {
	vecty.Core

	PanelID        int                 `vecty:"prop"`
	KeyValue       string              `vecty:"prop"`
	IsFirst        bool                `vecty:"prop"`
	IsLast         bool                `vecty:"prop"`
	Style          string              `vecty:"prop"`
	Resize         string              `vecty:"prop"`
	OnWindowResize func(int, int, int) `vecty:"prop"`

	ResizeStyle string
	ResizeObj   *vecty.HTML `vecty:"ref"`
	Element     *vecty.HTML `vecty:"ref"`
	Children    vecty.List  `vecty:"slot"`
}

func (c *PanelWrapper) Key() interface{} {
	return c.KeyValue
}

func (c *PanelWrapper) Mount() {
	if c.Resize == "stretch" {
		c.ResizeObj.Node().Call("addEventListener", "load", js.NewCallback(func(args []js.Value) {
			c.ResizeObj.Node().Get("contentDocument").Get("defaultView").Call("addEventListener", "resize", js.NewCallback(func(args []js.Value) {
				c.calculateStretchWidth()
			}))
		}))
		c.ResizeObj.Node().Set("data", "about:blank")
		js.Global().Call("setTimeout", js.NewCallback(func(args []js.Value) {
			js.Global().Get("window").Call("requestAnimationFrame", js.NewCallback(func(args []js.Value) {
				c.calculateStretchWidth()
			}))
		}), 0)
	}
}

func (c *PanelWrapper) calculateStretchWidth() {
	if c.OnWindowResize != nil {
		rect := c.Element.Node().Call("getBoundingClientRect")
		c.OnWindowResize(c.PanelID, rect.Get("width").Int(), rect.Get("height").Int())
	}
}

func (c *PanelWrapper) Render() vecty.ComponentOrHTML {
	c.ResizeStyle = InlineStyle(map[string]string{
		"position": "absolute",
		"top":      "0",
		"left":     "0",
		"width":    "100%",
		"height":   "100%",
		"zIndex":   "-1",
		"opacity":  "0",
	})
	return webui.Render(c)
}
