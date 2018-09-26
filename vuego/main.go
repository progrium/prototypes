package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Data map[string]interface{}

type Renderer func(out io.Writer, data Data) error

func main() {
	s := `<p>Links:</p><ul><li v-for="Name in Names"><a v-if="UseFoo" v-bind:href="Foo">{{ .Foo }} {{ .Name }}</a><li><a href="/bar/baz">BarBaz</a></ul>`
	fmt.Println(render(s, Data{"Foo": "Hi", "UseFoo": true, "Names": []string{"Jeff", "Buster", "James", "Progrium"}}))
}

func render(tmpl string, data Data) string {
	buf := &bytes.Buffer{}
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
	var f func(*html.Node) Renderer
	f = func(n *html.Node) Renderer {
		if n.Type == html.DocumentNode {
			return func(out io.Writer, data Data) error {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if err := f(c)(out, data); err != nil {
						return err
					}
				}
				return nil
			}
		}
		if n.Type == html.ElementNode {
			return func(out io.Writer, data Data) error {
				var iter []string
				iterName := ""
				attrs := make(map[string]string)
				for _, a := range n.Attr {
					if strings.HasPrefix(a.Key, "v-if") {
						ok, exists := data[a.Val]
						if !exists || !ok.(bool) {
							return nil
						}
					} else if strings.HasPrefix(a.Key, "v-bind:") {
						key := strings.Replace(a.Key, "v-bind:", "", 1)
						attrs[key] = data[a.Val].(string)
					} else if strings.HasPrefix(a.Key, "v-on:") {
						key := strings.Replace(a.Key, "v-on:", "", 1)
						attrs["on"+key] = data[a.Val].(string)
					} else if strings.HasPrefix(a.Key, "v-for") {
						parts := strings.Split(a.Val, " in ")
						// TODO: reflect and iterate over idx not a placeholder type
						iter = data[parts[1]].([]string)
						iterName = parts[0]
					} else {
						attrs[a.Key] = a.Val
					}
				}
				fmt.Fprintf(out, "  <%s%s>\n", n.Data, htmlAttrs(attrs))
				indentOut := NewIndentWriter(out, []byte("  "))
				if iter == nil {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if err := f(c)(indentOut, data); err != nil {
							return err
						}
					}
				} else {
					for _, el := range iter {
						data[iterName] = el
						for c := n.FirstChild; c != nil; c = c.NextSibling {
							if err := f(c)(indentOut, data); err != nil {
								return err
							}
						}
					}
				}
				fmt.Fprintf(out, "  </%s>\n", n.Data)
				return nil
			}
		}
		if n.Type == html.TextNode {
			t := template.Must(template.New("").Parse(n.Data + "\n"))
			return func(out io.Writer, data Data) error {
				out = NewIndentWriter(out, []byte("  "))
				return t.Execute(out, data)
			}
		}
		return nil
	}
	f(doc)(buf, data)
	return buf.String()
}

func htmlAttrs(attrs map[string]string) string {
	s := ""
	for k, v := range attrs {
		s = fmt.Sprintf("%s %s=\"%s\"", s, k, v)
	}
	return s
}

// Writer indents each line of its input.
type indentWriter struct {
	w   io.Writer
	bol bool
	pre [][]byte
	sel int
	off int
}

// NewIndentWriter makes a new write filter that indents the input
// lines. Each line is prefixed in order with the corresponding
// element of pre. If there are more lines than elements, the last
// element of pre is repeated for each subsequent line.
func NewIndentWriter(w io.Writer, pre ...[]byte) io.Writer {
	return &indentWriter{
		w:   w,
		pre: pre,
		bol: true,
	}
}

// The only errors returned are from the underlying indentWriter.
func (w *indentWriter) Write(p []byte) (n int, err error) {
	for _, c := range p {
		if w.bol {
			var i int
			i, err = w.w.Write(w.pre[w.sel][w.off:])
			w.off += i
			if err != nil {
				return n, err
			}
		}
		_, err = w.w.Write([]byte{c})
		if err != nil {
			return n, err
		}
		n++
		w.bol = c == '\n'
		if w.bol {
			w.off = 0
			if w.sel < len(w.pre)-1 {
				w.sel++
			}
		}
	}
	return n, nil
}
