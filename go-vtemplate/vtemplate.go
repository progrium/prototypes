package vtemplate

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	reflected "github.com/progrium/prototypes/go-reflected"
	"golang.org/x/net/html"
)

type NodeType uint32

const (
	NullNode NodeType = iota
	TextNode
	ElementNode
	CustomNode
)

type Node struct {
	Name     string
	Text     string
	Type     NodeType
	Attrs    map[string]interface{}
	Children []*Node
	Html     *html.Node
	Data     reflected.Value
}

type Binding struct {
	Name       string
	Argument   string
	Expression string
	IterVar    string
	Value      reflected.Value
	Node       *Node
	Parser     *Parser
}

type CustomElement interface {
	Parse(*Node, *html.Node, reflected.Value) error
}

type Directive interface {
	Apply(binding Binding) error
}

type Evaluator interface {
	Set(name string, value interface{})
	Unset(name string)
	Resolve(exp string) (reflected.Value, error)
}

type Parser struct {
	Directives     map[string]Directive
	CustomElements map[string]CustomElement
	Evaluator      Evaluator

	interpRegex *regexp.Regexp
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) init() {
	if p.interpRegex == nil {
		p.interpRegex = regexp.MustCompile("{{[^}]+}}")
	}
	if p.CustomElements == nil {
		p.CustomElements = make(map[string]CustomElement)
	}
	if p.Directives == nil {
		p.Directives = make(map[string]Directive)
	}
}

func (p *Parser) Parse(r io.Reader, data interface{}) (*Node, error) {
	p.init()
	var buf bytes.Buffer
	doc, err := html.Parse(io.TeeReader(r, &buf))
	if err != nil {
		return nil, err
	}
	rdata := reflected.ValueOf(data)
	// html.Parse always returns a full html document with head and body.
	// if our template was for the body, grab the body, otherwise
	// grab what's inside the body
	if strings.HasPrefix(buf.String(), "<body") {
		return p.ParseNode(doc.LastChild.LastChild, rdata)
	}
	return p.ParseNode(doc.LastChild.LastChild.LastChild, rdata)
}

func (p *Parser) ParseNode(h *html.Node, data reflected.Value) (*Node, error) {
	n := &Node{Html: h, Data: data}
	if p.Evaluator != nil {
		for _, f := range data.Props() {
			p.Evaluator.Set(f, data.Get(f).Interface())
		}
	}
	switch h.Type {
	case html.ElementNode:
		return p.parseElement(n, h, data)
	case html.TextNode:
		return p.parseText(n, h, data)
	default:
		return nil, nil
	}
}

func (p *Parser) parseElement(n *Node, h *html.Node, data reflected.Value) (*Node, error) {
	n.Type = ElementNode
	n.Name = h.Data
	n.Attrs = make(map[string]interface{})
	for _, attr := range h.Attr {
		n.Attrs[attr.Key] = attr.Val
	}
	if err := p.applyDirectives(n, data); err != nil {
		return nil, err
	}
	if n.Type == NullNode {
		return nil, nil
	}
	if n.Children == nil {
		for c := h.FirstChild; c != nil; c = c.NextSibling {
			cn, err := p.ParseNode(c, data)
			if err != nil {
				return nil, err
			}
			if cn != nil {
				n.Children = append(n.Children, cn)
			}
		}
	}
	custom, ok := p.CustomElements[n.Name]
	if ok {
		n.Type = CustomNode
		var err error
		if custom != nil {
			err = custom.Parse(n, h, data)
		}
		return n, err
	}
	return n, nil
}

func (p *Parser) applyDirectives(n *Node, data reflected.Value) error {
	for k, v := range n.Attrs {
		exp, ok := v.(string)
		if !ok {
			continue
		}
		if strings.HasPrefix(k, "v-") {
			parts := strings.Split(k, ":")
			var name, arg string
			name = strings.Replace(parts[0], "v-", "", 1)
			if len(parts) > 1 {
				arg = parts[1]
			}
			dir, ok := p.Directives[name]
			if !ok {
				continue
			}
			parts = strings.Split(exp, " in ")
			var iterVar string
			if len(parts) < 2 {
				exp = parts[0]
			} else {
				iterVar = parts[0]
				exp = parts[1]
			}
			val, err := p.resolveExp(exp, data)
			if err != nil {
				return err
			}
			binding := Binding{
				Name:       name,
				Argument:   arg,
				Expression: exp,
				Value:      val,
				IterVar:    iterVar,
				Node:       n,
				Parser:     p,
			}
			if err := dir.Apply(binding); err != nil {
				return err
			}
			delete(n.Attrs, k)
		}
	}
	return nil
}

func (p *Parser) parseText(n *Node, h *html.Node, data reflected.Value) (*Node, error) {
	text := strings.Trim(h.Data, " \n\t")
	if text == "" {
		return nil, nil
	}
	var errs []error
	n.Type = TextNode
	n.Text = p.interpRegex.ReplaceAllStringFunc(text, func(tag string) string {
		v, err := p.resolveExp(tag[2:len(tag)-2], data)
		if err != nil {
			errs = append(errs, err)
			return ""
		}
		return fmt.Sprint(v.Interface())
	})
	if len(errs) > 0 {
		return n, errs[0]
	}
	return n, nil
}

func (p *Parser) resolveExp(expression string, data reflected.Value) (reflected.Value, error) {
	expression = strings.Trim(expression, " ")
	for _, prop := range data.Members() {
		if expression == prop {
			return data.Get(prop), nil
		}
	}
	if p.Evaluator == nil {
		return reflected.Undefined(), fmt.Errorf("unable to resolve expression: '%s'", expression)
	}
	return p.Evaluator.Resolve(expression)
}
