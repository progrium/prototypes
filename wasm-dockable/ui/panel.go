package ui

import (
	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(Panel{})
}

type WindowData struct {
	PanelWidth
}

type PanelWidth struct {
	Size    int
	Resize  string
	MinSize int
	Snap    []int
}

type Panel struct {
	vecty.Core

	Style       string
	PanelWidths []PanelWidth
	Spacing     int
	Expanded    bool
	Windows     []WindowData
}

func (c *Panel) Render() vecty.ComponentOrHTML {
	c.PanelWidths = nil
	for _, win := range c.Windows {
		width := PanelWidth{
			Size:    ((len(c.Windows) - 1) * 29) + 13,
			Resize:  "fixed",
			MinSize: win.MinSize,
		}
		if c.Expanded {
			width.Size = win.Size
			width.Resize = win.Resize
		}
		c.PanelWidths = append(c.PanelWidths, width)
	}
	c.Style = InlineStyle(map[string]string{
		"flexGrow":      "1",
		"display":       "flex",
		"flexDirection": "column",
		"minWidth":      "0",
	})
	return webui.Render(c)
}
