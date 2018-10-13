package main

import (
	"bytes"
	"os"

	vtemplate "github.com/progrium/prototypes/go-vtemplate"
	"github.com/progrium/prototypes/go-vtemplate/goja"
	"github.com/progrium/prototypes/go-vtemplate/html"
)

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
			{{ Person.Name.split(" ")[0] }} is {{ new Date().getFullYear() - Person.YOB }}.
		</li>
	</ul>`
	p := &vtemplate.Parser{
		Directives: html.BuiltinDirectives(),
		Evaluator:  goja.Evaluator(),
	}
	n, _ := p.Parse(bytes.NewBufferString(template), data)
	html.Render(os.Stdout, n)
}
