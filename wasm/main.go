package main

import (
	"bytes"
	"html/template"
	"log"
	"reflect"
	"strings"

	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"golang.org/x/net/html"
)

var registry map[string]reflect.Type

func init() {
	registry = map[string]reflect.Type{
		"footer":   reflect.TypeOf(Footer{}),
		"markdown": reflect.TypeOf(Markdown{}),
	}
}

func GetPropFields(t reflect.Type) []string {
	propFields := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag == `vecty:"prop"` {
			propFields = append(propFields, field.Name)
		}
	}
	return propFields
}

func SetFieldString(v reflect.Value, name, value string) {
	v.Elem().FieldByName(strings.Title(name)).SetString(value)
}

func GetFieldString(v reflect.Value, name string) string {
	return v.Elem().FieldByName(name).String()
}

func CallMethod(v reflect.Value, method string, args ...interface{}) {
	var in []reflect.Value
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	v.MethodByName(method).Call(in)
}

func main() {
	c := make(chan struct{}, 0)
	vecty.SetTitle("Markdown Demo")
	vecty.RenderBody(&PageView{
		Input: `# Markdown Example

This is a live editor, try editing the Markdown on the right of the page.
`,
	})
	<-c
}

// PageView is our main page component.
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
	// return render(`
	// <body>
	// 	<div style="float: right">
	// 		<textarea
	// 			v-on:input="OnTextAreaChange"
	// 			rows="14"
	// 			cols="70"
	// 			style="font-family: monospace;">{{.Input}}</textarea>
	// 	</div>
	// 	<Markdown v-bind:Input="Input" />
	// 	<Footer Copyright="2018 Brian" />
	// </body>
	// `, p))
	return elem.Body(
		render(`<div style="float: right">
				<textarea 
					v-on:input="OnTextAreaChange" 
					rows="14" 
					cols="70" 
					style="font-family: monospace;">{{.Input}}</textarea>
			</div>`, p),
		render(`<Markdown v-bind:Input="Input" />`, p),
		render(`<Footer Copyright="2018 Jeff" />`, p),
	)
}

// Markdown is a simple component which renders the Input markdown as sanitized
// HTML into a div.
type Markdown struct {
	vecty.Core
	Input string `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (m *Markdown) Render() vecty.ComponentOrHTML {
	// Render the markdown input into HTML using Blackfriday.
	unsafeHTML := blackfriday.MarkdownCommon([]byte(m.Input))

	// Sanitize the HTML.
	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	// Return the HTML, which we guarantee to be safe / sanitized.
	return elem.Div(
		vecty.Markup(
			vecty.UnsafeHTML(safeHTML),
		),
	)
}

type Footer struct {
	vecty.Core
	Copyright string `vecty:"prop"`
}

func (m *Footer) template() string {
	return `<div foo="bar" class="footer">{{ .Copyright }}</div>`
}

// Render implements the vecty.Component interface.
func (m *Footer) Render() vecty.ComponentOrHTML {
	return render(m.template(), m)
}

func render(tmpl string, v interface{}) vecty.ComponentOrHTML {
	doc := &html.Node{
		Type: html.ElementNode,
		Data: "doc",
	}
	els, err := html.ParseFragment(strings.NewReader(tmpl), doc)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range els {
		doc.AppendChild(v)
	}
	return nodeToVecty(doc.LastChild, v)
}

func nodeToVecty(n *html.Node, v interface{}) vecty.ComponentOrHTML {
	switch n.Type {
	case html.ElementNode:
		for name, comType := range registry {
			if name == n.Data {
				c := reflect.New(comType)
				for _, a := range n.Attr {
					for _, prop := range GetPropFields(comType) {
						if a.Key == strings.ToLower(prop) {
							SetFieldString(c, prop, a.Val)
						}
					}
					if strings.HasPrefix(a.Key, "v-bind:") {
						key := strings.Replace(a.Key, "v-bind:", "", 1)
						val := GetFieldString(reflect.ValueOf(v), a.Val)
						SetFieldString(c, key, val)
					}
				}

				ch := c.Interface().(vecty.ComponentOrHTML)
				return ch
			}
		}
		var applyers []vecty.Applyer
		for _, a := range n.Attr {
			// todo: v-bind
			if strings.HasPrefix(a.Key, "v-on:") {
				attr := a
				key := strings.Replace(attr.Key, "v-on:", "", 1)
				rcvr := reflect.ValueOf(v)
				applyers = append(applyers, &vecty.EventListener{
					Name: key,
					Listener: func(e *vecty.Event) {
						CallMethod(rcvr, attr.Val, e)
					},
				})
				continue
			}
			applyers = append(applyers, vecty.Attribute(a.Key, a.Val))
		}
		children := []vecty.MarkupOrChild{vecty.Markup(applyers...)}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			children = append(children, nodeToVecty(c, v))
		}
		return vecty.Tag(n.Data, children...)
	case html.TextNode:
		t := template.Must(template.New("").Parse(n.Data))
		b := &bytes.Buffer{}
		if err := t.Execute(b, v); err != nil {
			panic(err)
		}
		return vecty.Text(b.String())
	default:
		return vecty.Text("")
	}
}
