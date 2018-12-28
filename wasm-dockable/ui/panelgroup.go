package ui

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gowasm/vecty"
	"github.com/progrium/prototypes/go-webui"
)

func init() {
	webui.Register(PanelGroup{})
}

type panelDef struct {
	Size      int
	Resize    string
	MinSize   int
	FixedSize int
	Snap      []int
}

type PanelGroup struct {
	vecty.Core

	Spacing     int              `vecty:"prop" default:"1"`
	Direction   string           `vecty:"prop" default:"row"`
	PanelWidths []PanelWidth     `vecty:"prop"`
	PanelColor  string           `vecty:"prop"`
	BorderColor string           `vecty:"prop"`
	ShowHandles bool             `vecty:"prop"`
	OnUpdate    func([]panelDef) `vecty:"prop"`

	Panels       []panelDef
	Children     vecty.List `vecty:"slot"`
	OrigChildren vecty.List
	Element      *vecty.HTML `vecty:"ref"`

	ContainerStyle string
}

func (c *PanelGroup) SkipRender(prev vecty.Component) (skip bool) {
	skip = false
	if c.Panels == nil {
		c.loadPanels()
	}
	if len(c.PanelWidths) == 0 {
		return
	}
	prevCom := prev.(*PanelGroup)
	if len(c.Panels) != len(prevCom.Panels) {
		c.loadPanels()
		return
	}
	for i, panel := range c.PanelWidths {
		if panel.Size != prevCom.PanelWidths[i].Size ||
			panel.MinSize != prevCom.PanelWidths[i].MinSize ||
			panel.Resize != prevCom.PanelWidths[i].Resize {
			c.loadPanels()
			break
		}
	}
	return
}

func (c *PanelGroup) loadPanels() {
	if len(c.Children) == 0 {
		return
	}
	defaultSize := 256
	defaultMinSize := 48
	defaultResize := "stretch"
	stretchIncluded := false
	c.Panels = nil
	for i, _ := range c.Children {
		panel := panelDef{
			Size:    defaultSize,
			Resize:  defaultResize,
			MinSize: defaultMinSize,
		}
		if i < len(c.PanelWidths) {
			panel.Size = c.PanelWidths[i].Size
			panel.Resize = c.PanelWidths[i].Resize
			panel.MinSize = c.PanelWidths[i].MinSize
			panel.Snap = c.PanelWidths[i].Snap
		}
		c.Panels = append(c.Panels, panel)
		if c.Panels[i].Resize == "stretch" {
			stretchIncluded = true
		}
		if i == len(c.Children)-1 && !stretchIncluded {
			c.Panels[i].Resize = "stretch"
		}
	}
}

func (c *PanelGroup) sizeDir(cap bool) string {
	dir := "width"
	if c.Direction == "column" {
		dir = "height"
	}
	if cap {
		dir = strings.Title(dir)
	}
	return dir
}

func (c *PanelGroup) panelMaxSize(idx int, panels []panelDef) int {
	if panels[idx].Resize == "fixed" {
		if panels[idx].FixedSize == 0 {
			panels[idx].FixedSize = panels[idx].Size
		}
		return panels[idx].FixedSize
	}
	return 0
}

func (c *PanelGroup) panelMinSize(idx int, panels []panelDef) int {
	if panels[idx].Resize == "fixed" {
		if panels[idx].FixedSize == 0 {
			panels[idx].FixedSize = panels[idx].Size
		}
		return panels[idx].FixedSize
	}
	return panels[idx].MinSize
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *PanelGroup) handleResize(idx int, deltaX, deltaY int) (delta int) {
	tmpPanels := c.Panels[:]
	delta = deltaY
	if c.Direction == "row" {
		delta = deltaX
	}
	delta = c.resizePanel(idx, delta, tmpPanels)
	c.Panels = tmpPanels
	if c.OnUpdate != nil {
		c.OnUpdate(tmpPanels)
	}
	vecty.Rerender(c)
	return
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func (c *PanelGroup) resizePanel(idx int, delta int, panels []panelDef) int {
	masterSize := 0
	for i := 0; i < len(panels); i++ {
		masterSize += panels[i].Size
	}
	rect := c.Element.Node().Call("getBoundingClientRect")
	baseSize := rect.Get("width").Int()
	if c.Direction == "column" {
		baseSize = rect.Get("height").Int()
	}
	size := baseSize - c.Spacing*(len(c.Children)-1)
	if masterSize != size {
		fmt.Println("ERROR: sizes don't match?", masterSize, size)
		panels[idx].Size += size - masterSize
	}

	result := delta

	panels[idx].Size += delta
	panels[idx+1].Size -= delta

	minSize := c.panelMinSize(idx, panels)
	maxSize := c.panelMaxSize(idx, panels)

	if panels[idx].Size < minSize {
		delta = minSize - panels[idx].Size
		if idx == 0 {
			result = c.resizePanel(idx, delta, panels)
		} else {
			result = c.resizePanel(idx-1, -delta, panels)
		}
	}

	if maxSize != 0 && panels[idx].Size > maxSize {
		delta = panels[idx].Size - maxSize
		if idx == 0 {
			result = c.resizePanel(idx, -delta, panels)
		} else {
			result = c.resizePanel(idx-1, delta, panels)
		}
	}

	minSize = c.panelMinSize(idx+1, panels)
	maxSize = c.panelMaxSize(idx+1, panels)

	if panels[idx+1].Size < minSize {
		delta = minSize - panels[idx+1].Size
		if idx+1 == len(panels)-1 {
			result = c.resizePanel(idx, -delta, panels)
		} else {
			result = c.resizePanel(idx+1, delta, panels)
		}
	}

	if maxSize != 0 && panels[idx+1].Size > maxSize {
		delta = panels[idx+1].Size - maxSize
		if idx+1 == len(panels)-1 {
			result = c.resizePanel(idx, delta, panels)
		} else {
			result = c.resizePanel(idx+1, -delta, panels)
		}
	}

	for i := 0; i < len(panels[idx].Snap); i++ {
		if abs(panels[idx].Snap[i]-panels[idx].Size) < 20 {
			delta = panels[idx].Snap[i] - panels[idx].Size
			min1 := panels[idx].Size + delta
			min2 := panels[idx+1].Size - delta
			if delta != 0 && min1 >= c.panelMinSize(idx, panels) && min2 >= c.panelMinSize(idx+1, panels) {
				result = c.resizePanel(idx, delta, panels)
			}
		}
	}

	for i := 0; i < len(panels[idx+1].Snap); i++ {
		if abs(panels[idx+1].Snap[i]-panels[idx+1].Size) < 20 {
			delta = panels[idx+1].Snap[i] - panels[idx+1].Size
			min1 := panels[idx].Size + delta
			min2 := panels[idx+1].Size - delta
			if delta != 0 && min1 >= c.panelMinSize(idx, panels) && min2 >= c.panelMinSize(idx+1, panels) {
				result = c.resizePanel(idx, -delta, panels)
			}
		}
	}

	return result
}

func (c *PanelGroup) setPanelSize(idx int, sizeX int, sizeY int) {
	size := sizeX
	if c.Direction == "column" {
		size = sizeY
	}
	if size != c.Panels[idx].Size {
		tmpPanels := c.Panels[:]
		if size < tmpPanels[idx].MinSize {
			diff := tmpPanels[idx].MinSize - size
			tmpPanels[idx].Size = tmpPanels[idx].MinSize
			for i := 0; i < len(tmpPanels); i++ {
				if i != idx && tmpPanels[i].Resize == "dynamic" {
					available := tmpPanels[i].Size - tmpPanels[i].MinSize
					cut := min(diff, available)
					tmpPanels[i].Size = tmpPanels[i].Size - cut
					diff = diff - cut
					if diff == 0 {
						break
					}
				}
			}
		} else {
			tmpPanels[idx].Size = size
		}
		c.Panels = tmpPanels

		if idx > 0 {
			c.handleResize(idx-1, 0, 0)
		} else if len(c.Panels) > 2 {
			c.handleResize(idx+1, 0, 0)
		}

		// if cb != nil {
		// 	cb()
		// }
	}
}

func (c *PanelGroup) panelGroupMinSize(spacing int) int {
	size := 0
	for i := 0; i < len(c.Panels); i++ {
		size += c.panelMinSize(i, c.Panels)
	}
	return size + (len(c.Panels)-1)*spacing
}

func (c *PanelGroup) Render() vecty.ComponentOrHTML {
	if c.Panels == nil {
		c.loadPanels()
	}
	minDirKey := fmt.Sprintf("min%s", c.sizeDir(true))
	c.ContainerStyle = InlineStyle(map[string]string{
		minDirKey:       strconv.Itoa(c.panelGroupMinSize(c.Spacing)),
		"flexDirection": c.Direction,
		"width":         "100%",
		"height":        "100%",
		"display":       "flex",
		"flexGrow":      "1",
	})

	if c.OrigChildren == nil && len(c.Children) > 0 {
		c.OrigChildren = c.Children[:]
	}
	c.Children = nil
	stretchIncluded := false

	for i := 0; i < len(c.OrigChildren); i++ {
		dirKey := "width"
		if c.Direction == "row" {
			dirKey = "height"
		}
		minDir := c.Panels[i].Size
		flexGrow := 0
		flexShrink := 0
		if c.Panels[i].Resize == "stretch" {
			minDir = 0
			flexGrow = 1
			flexShrink = 1
		}
		panelStyle := InlineStyle(map[string]string{
			"backgroundColor": c.PanelColor,
			"flexGrow":        strconv.Itoa(flexGrow),
			"flexShrink":      strconv.Itoa(flexShrink),
			"display":         "flex",
			"overflow":        "hidden",
			"position":        "relative",
			dirKey:            "100%",
			minDirKey:         strconv.Itoa(minDir),
			c.sizeDir(false):  strconv.Itoa(c.Panels[i].Size),
		})

		isFirst := (i == 0)
		isLast := (i == len(c.OrigChildren)-1)
		resize := c.Panels[i].Resize
		var onWindowResize func(idx int, sizeX int, sizeY int)
		if resize == "stretch" {
			onWindowResize = c.setPanelSize
			stretchIncluded = true
		}
		if !stretchIncluded && isLast {
			resize = "stretch"
		}

		c.Children = append(c.Children, &PanelWrapper{
			Style:          panelStyle,
			PanelID:        i,
			KeyValue:       fmt.Sprintf("panel%d", i),
			IsFirst:        isFirst,
			IsLast:         isLast,
			Resize:         resize,
			OnWindowResize: onWindowResize,
			Children:       vecty.List{c.OrigChildren[i]},
		})

		if i < len(c.OrigChildren)-1 {
			c.Children = append(c.Children, &Divider{
				BorderColor:  c.BorderColor,
				KeyValue:     fmt.Sprintf("divider%d", i),
				PanelID:      i,
				HandleResize: c.handleResize,
				DividerWidth: c.Spacing,
				Direction:    c.Direction,
				ShowHandles:  c.ShowHandles,
			})
		}

	}

	return webui.Render(c)
}
