package app

import "github.com/gowasm/vecty"

type PageView struct {
	vecty.Core
	Input string
}

func (p *PageView) OnTextAreaChange(e *vecty.Event) {
	// When input is typed into the textarea, update the local
	// component state and rerender.
	p.Input = e.Target.Get("value").String()
	vecty.Rerender(p)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return render(`<body> 
			<div style="float: right">
				<textarea 
					v-on:input="OnTextAreaChange" 
					rows="14" 
					cols="70" 
					style="font-family: monospace;">{{Input}}</textarea>
			</div>
			<Markdown v-bind:Input="Input"></Markdown>
			<Footer Copyright="2018 Jeff">
				<p>Hello worlds</p>
			</Footer> 
		</body>`, p)
}
