package ui

import (
	"strconv"
	"syscall/js"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Divider{})
}

type Divider struct {
	vecty.Core

	Direction    string                  `vecty:"prop"`
	KeyValue     string                  `vecty:"prop"`
	HandleResize func(int, int, int) int `vecty:"prop"`
	DividerWidth int                     `vecty:"prop" default:"1"`
	HandleBleed  int                     `vecty:"prop" default:"4"`
	BorderColor  string                  `vecty:"prop"`
	PanelID      int                     `vecty:"prop"`
	ShowHandles  bool                    `vecty:"prop"`

	DividerStyle string
	HandleStyle  string
	ClassName    string
	Dragging     bool
	InitX        int
	InitY        int

	onMouseMove *js.Callback
	onMouseUp   *js.Callback
}

func (c *Divider) Key() interface{} {
	return c.KeyValue
}

func (c *Divider) OnMouseDown(e *vecty.Event) {
	if e.Get("button").Int() != 0 {
		return
	}
	if c.onMouseMove == nil {
		mm := js.NewCallback(c.OnMouseMove)
		c.onMouseMove = &mm
		mu := js.NewCallback(c.OnMouseUp)
		c.onMouseUp = &mu
	}
	js.Global().Get("document").Call("addEventListener", "mousemove", *c.onMouseMove, map[string]interface{}{"once": true})
	js.Global().Get("document").Call("addEventListener", "mouseup", *c.onMouseUp, map[string]interface{}{"once": true})
	c.Dragging = true
	c.InitX = e.Get("pageX").Int()
	c.InitY = e.Get("pageY").Int()
	e.Call("stopPropagation")
	e.Call("preventDefault")
}

func (c *Divider) OnMouseUp(args []js.Value) {
	c.Dragging = false
	args[0].Call("stopPropagation")
	args[0].Call("preventDefault")
}

func (c *Divider) OnMouseMove(args []js.Value) {
	if !c.Dragging {
		return
	}
	js.Global().Get("document").Call("addEventListener", "mousemove", *c.onMouseMove, map[string]interface{}{"once": true})
	pageX := args[0].Get("pageX").Int()
	pageY := args[0].Get("pageY").Int()
	deltaX := pageX - c.InitX
	deltaY := pageY - c.InitY
	flowX := 0
	if c.Direction == "row" {
		flowX = 1
	}
	flowY := 0
	if c.Direction == "column" {
		flowY = 1
	}
	flowDelta := deltaX*flowX + deltaY*flowY
	resultDelta := c.HandleResize(c.PanelID, deltaX, deltaY)
	if (resultDelta + flowDelta) != 0 {
		if resultDelta == flowDelta {
			c.InitX = pageX
			c.InitY = pageY
		} else {
			c.InitX = pageX + resultDelta*flowX
			c.InitY = pageY + resultDelta*flowY
		}

	}
	args[0].Call("stopPropagation")
	args[0].Call("preventDefault")
	vecty.Rerender(c)
}

func (c *Divider) handleWidth() int {
	return c.DividerWidth + c.HandleBleed*2
}

func (c *Divider) handleOffset() int {
	return c.DividerWidth/2 - c.handleWidth()/2
}

func (c *Divider) Render() vecty.ComponentOrHTML {
	baseDividerStyle := map[string]string{
		"flexGrow":        "0",
		"position":        "relative",
		"backgroundColor": c.BorderColor,
	}
	baseHandleStyle := map[string]string{
		"position":        "absolute",
		"zIndex":          "100",
		"backgroundColor": "auto",
	}
	if c.ShowHandles {
		baseHandleStyle["backgroundColor"] = "rgba(0,128,255,0.25)"
	}
	switch c.Direction {
	case "row":
		c.DividerStyle = InlineStyle(map[string]string{
			"width":     strconv.Itoa(c.DividerWidth),
			"minWidth":  strconv.Itoa(c.DividerWidth),
			"maxWdith":  strconv.Itoa(c.DividerWidth),
			"height":    "auto",
			"minHeight": "auto",
			"maxHeight": "auto",
		}, baseDividerStyle)
		c.HandleStyle = InlineStyle(map[string]string{
			"width":  strconv.Itoa(c.handleWidth()),
			"left":   strconv.Itoa(c.handleOffset()),
			"cursor": "col-resize",
			"height": "100%",
			"top":    "0",
		}, baseHandleStyle)
	case "column":
		c.DividerStyle = InlineStyle(map[string]string{
			"width":     "auto",
			"minWidth":  "auto",
			"maxWdith":  "auto",
			"height":    strconv.Itoa(c.DividerWidth),
			"minHeight": strconv.Itoa(c.DividerWidth),
			"maxHeight": strconv.Itoa(c.DividerWidth),
		}, baseDividerStyle)
		c.HandleStyle = InlineStyle(map[string]string{
			"width":  "100%",
			"left":   "0",
			"cursor": "row-resize",
			"height": strconv.Itoa(c.handleWidth()),
			"top":    strconv.Itoa(c.handleOffset()),
		}, baseHandleStyle)
	default:
		panic("invalid value for direction property: " + c.Direction)
	}
	c.ClassName = "divider"
	if c.Dragging {
		c.ClassName = "divider dragging"
	}
	return webui.Render(c)
}
