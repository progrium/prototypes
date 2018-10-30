package ui

import (
	"strconv"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Dragger{})
}

type Dragger struct {
	vecty.Core

	Dragging bool
	Style    string
	mX, mY   int
	sX, sY   int
	X, Y     int
}

func (c *Dragger) OnMouseDown(e *vecty.Event) {
	c.Dragging = true
	c.mX = e.Get("pageX").Int()
	c.mY = e.Get("pageY").Int()
	c.sX = c.X
	c.sY = c.Y
}

func (c *Dragger) OnMouseMove(e *vecty.Event) {
	if !c.Dragging {
		return
	}
	c.X = c.sX + (e.Get("pageX").Int() - c.mX)
	c.Y = c.sY + (e.Get("pageY").Int() - c.mY)
	vecty.Rerender(c)
}

func (c *Dragger) OnMouseDone(e *vecty.Event) {
	c.Dragging = false
}

func (c *Dragger) Render() vecty.ComponentOrHTML {
	c.Style = InlineStyle(map[string]string{
		"width":      "100",
		"height":     "100",
		"left":       strconv.Itoa(c.X),
		"top":        strconv.Itoa(c.Y),
		"position":   "absolute",
		"background": "red",
	})
	return webui.Render(c)
}
