package app

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/microcosm-cc/bluemonday"
	"github.com/progrium/prototypes/wasm/pkg/webui"
	"github.com/russross/blackfriday"
)

func init() {
	webui.Register(Markdown{})
}

type Markdown struct {
	vecty.Core
	Input string `vecty:"prop"`
}

func (m *Markdown) Render() vecty.ComponentOrHTML {
	unsafeHTML := blackfriday.MarkdownCommon([]byte(m.Input))
	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	return elem.Div(
		vecty.Markup(
			vecty.UnsafeHTML(safeHTML),
		),
	)
}
