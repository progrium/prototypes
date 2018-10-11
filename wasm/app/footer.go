package app

import "github.com/gowasm/vecty"

type Footer struct {
	vecty.Core
	Copyright string     `vecty:"prop"`
	Children  vecty.List `vecty:"slot"`
}

func (m *Footer) template() string {
	return `<div class="footer">{{ Copyright }} <slot></slot></div>`
}

// Render implements the vecty.Component interface.
func (m *Footer) Render() vecty.ComponentOrHTML {
	return render(m.template(), m)
}
