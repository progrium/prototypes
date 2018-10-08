package main

import (
	"bytes"
	"os"

	reflected "github.com/progrium/prototypes/go-reflected"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
	"github.com/progrium/prototypes/go-vtemplate/goja"
	"github.com/progrium/prototypes/go-vtemplate/html"
	h "golang.org/x/net/html"
)

type DemoElement struct{}

func (e *DemoElement) Parse(n *vtemplate.Node, h *h.Node, data reflected.Value) error {
	n.Name = "demo"
	return nil
}

type Person struct {
	Name string
	YOB  int
}

func main() {
	data := map[string]interface{}{
		"People": []Person{
			{
				Name: "Jeff Lindsay",
				YOB:  1985,
			},
			{
				Name: "Gary Oldman",
				YOB:  1958,
			},
		},
	}
	template := `
	<ul v-for="Person in People">
		<li v-bind:class="Person.YOB >= 1982 ? 'millenial' : undefined">
			<foobar></foobar>
			{{ Person.Name.split(" ")[0] }} is {{ new Date().getFullYear() - Person.YOB }}.
		</li>
	</ul>`
	p := &vtemplate.Parser{
		Directives: html.BuiltinDirectives(),
		Evaluator:  goja.Evaluator(),
		CustomElements: map[string]vtemplate.CustomElement{
			"foobar": &DemoElement{},
		},
	}
	n, _ := p.Parse(bytes.NewBufferString(template), data)
	html.Render(os.Stdout, n)
}
