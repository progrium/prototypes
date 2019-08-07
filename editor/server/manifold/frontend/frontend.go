package frontend

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/progrium/prototypes/editor/server/manifold"
	"github.com/progrium/prototypes/editor/server/manifold/workspace"
	"github.com/progrium/prototypes/editor/server/repl"
	reflected "github.com/progrium/prototypes/go-reflected"
	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
)

type Field struct {
	Type       string      `msgpack:"type"`
	Name       string      `msgpack:"name"`
	Path       string      `msgpack:"path"`
	Value      interface{} `msgpack:"value"`
	Expression *string     `msgpack:"expression"`
	Fields     []Field     `msgpack:"fields"`
}

type Button struct {
	Name    string `msgpack:"name"`
	Path    string `msgpack:"path"`
	OnClick string `msgpack:"onclick"`
}

type Component struct {
	Name    string   `msgpack:"name"`
	Fields  []Field  `msgpack:"fields"`
	Buttons []Button `msgpack:"buttons"`
}

type Node struct {
	Name       string      `msgpack:"name"`
	Active     bool        `msgpack:"active"`
	Components []Component `msgpack:"components"`
}

type Project struct {
	Name string `msgpack:"name"`
	Path string `msgpack:"path"`
}

type State struct {
	Projects       []Project         `msgpack:"projects"`
	CurrentProject string            `msgpack:"currentProject"`
	Components     []string          `msgpack:"components"`
	Hierarchy      []string          `msgpack:"hierarchy"`
	Nodes          map[string]Node   `msgpack:"nodes"`
	NodePaths      map[string]string `msgpack:"nodePaths"`
}

func exportElem(v reflected.Value, path string, idx int, n *manifold.Node) (Field, bool) {
	elemPath := path + "/" + strconv.Itoa(idx)
	switch v.Type().Kind() {
	case reflect.Bool:
		return Field{
			Path:  elemPath,
			Type:  "boolean",
			Value: v.Interface(),
		}, true
	case reflect.String:
		return Field{
			Path:  elemPath,
			Type:  "string",
			Value: v.Interface(),
		}, true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return Field{
			Path:  elemPath,
			Type:  "number",
			Value: v.Interface(),
		}, true
	default:
		return Field{}, false
	}
}

func exportField(o reflected.Value, field, path string, n *manifold.Node) Field {
	var kind reflect.Kind
	if o.Type().Kind() == reflect.Struct {
		kind = o.Type().FieldType(field).Kind()
	} else {
		if !o.Get(field).IsValid() {
			kind = reflect.Invalid
		} else {
			kind = o.Get(field).Type().Kind()
		}
	}
	fieldPath := path + "/" + field
	var expr *string
	exprPath := fieldPath[len(n.FullPath())+1:]
	if e := n.Expression(exprPath); e != "" {
		expr = &e
	}
	switch kind {
	case reflect.Invalid:
		return Field{
			Name:       field,
			Path:       fieldPath,
			Expression: expr,
			Type:       "string",
			Value:      "INVALID",
		}
	case reflect.Bool:
		return Field{
			Name:       field,
			Path:       fieldPath,
			Expression: expr,
			Type:       "boolean",
			Value:      o.Get(field).Interface(),
		}
	case reflect.String:
		return Field{
			Name:       field,
			Path:       fieldPath,
			Expression: expr,
			Type:       "string",
			Value:      o.Get(field).Interface(),
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return Field{
			Path:       fieldPath,
			Name:       field,
			Expression: expr,
			Type:       "number",
			Value:      o.Get(field).Interface(),
		}
	case reflect.Struct:
		var fields []Field
		v := o.Get(field)
		for _, f := range v.Type().Fields() {
			fields = append(fields, exportField(v, f, fieldPath, n))
		}
		return Field{
			Path:       fieldPath,
			Name:       field,
			Expression: expr,
			Type:       "struct",
			Fields:     fields,
		}
	case reflect.Map:
		var fields []Field
		v := o.Get(field)
		for _, f := range v.Keys() {
			fields = append(fields, exportField(v, f, fieldPath, n))
		}
		return Field{
			Path:       fieldPath,
			Name:       field,
			Expression: expr,
			Type:       "map",
			Fields:     fields,
		}
	case reflect.Slice:
		var fields []Field
		v := o.Get(field)
		for idx, e := range v.Iter() {
			f, ok := exportElem(e, fieldPath, idx, n)
			if !ok {
				return Field{
					Name:       field,
					Path:       fieldPath,
					Expression: expr,
					Type:       "string",
					Value:      "UNSUPPORTED SLICE",
				}
			}
			fields = append(fields, f)
		}
		return Field{
			Path:       fieldPath,
			Name:       field,
			Expression: expr,
			Type:       "array",
			Fields:     fields,
		}
	case reflect.Ptr, reflect.Interface:
		var v interface{}
		if o.Get(field).IsValid() {
			v = o.Get(field).Interface()
		}
		t := o.Type().FieldType(field)
		if kind == reflect.Ptr {
			t = reflected.Type{Type: t.Elem()}
		}
		var path string
		if v != nil {
			refNode := n.Root().FindPtr(v)
			if refNode != nil {
				path = refNode.FullPath()
			}
		}
		return Field{
			Path:       fieldPath,
			Name:       field,
			Expression: expr,
			Type:       fmt.Sprintf("reference:%s", t.Name()),
			Value:      path,
		}
	default:
		panic(o.Type().FieldType(field).Kind())
	}
}

type ButtonProvider interface {
	InspectorButtons() []Button
}

func exportNodes(s *State, root *manifold.Node) {
	s.Hierarchy = []string{}
	s.Nodes = make(map[string]Node)
	manifold.Walk(root, func(n *manifold.Node) {
		s.Hierarchy = append(s.Hierarchy, n.FullPath())
		node := Node{
			Name:       n.Name,
			Active:     n.Active,
			Components: []Component{},
		}
		for _, com := range n.Components {
			var fields []Field
			c := reflected.ValueOf(com.Ref)
			path := n.FullPath() + "/" + com.Name
			for _, field := range c.Type().Fields() {
				fields = append(fields, exportField(c, field, path, n))
			}
			var buttons []Button
			p, ok := com.Ref.(ButtonProvider)
			if ok {
				buttons = p.InspectorButtons()
				for idx, button := range buttons {
					if button.OnClick != "" {
						continue
					}
					typ := reflect.ValueOf(com.Ref).Type()
					for i := 0; i < typ.NumMethod(); i++ {
						method := typ.Method(i)
						if method.Name != button.Name {
							break
						}
						if method.Type.NumIn() == 1 {
							buttons[idx].Path = path + "/" + method.Name
						}
					}
				}
			}

			node.Components = append(node.Components, Component{
				Name:    com.Name,
				Fields:  fields,
				Buttons: buttons,
			})
		}
		s.Nodes[n.ID] = node
		s.NodePaths[n.FullPath()] = n.ID
	})
}

type AppendNodeParams struct {
	ID   string
	Name string
}

type SetValueParams struct {
	Path     string
	Value    interface{}
	IntValue *int
	RefValue *string
}

type RemoveComponentParams struct {
	ID        string
	Component string
}

type NodeParams struct {
	ID     string
	Name   *string
	Active *bool
}

type DelegateParams struct {
	ID       string
	Contents string
}

func ListenAndServe(root *manifold.Node, addr string) error {
	state := State{
		Projects: []Project{
			{Name: "project1", Path: "/Project1"},
			{Name: "project2", Path: "/Project2"},
		},
		CurrentProject: "project1",
		Components:     manifold.RegisteredComponents(),
		Nodes:          make(map[string]Node),
		NodePaths:      make(map[string]string),
	}
	exportNodes(&state, root)

	clients := make(map[qrpc.Caller]string)

	sendState := func() {
		for client, callback := range clients {
			_, err := client.Call(callback, state, nil)
			if err != nil {
				delete(clients, client)
				log.Println(err)
			}
		}
	}

	// repl
	repl := repl.NewREPL(func(v interface{}) {
		fmt.Fprintf(os.Stdout, "%s\n", v)
	})

	// define api
	api := qrpc.NewAPI()
	api.HandleFunc("reload", func(r qrpc.Responder, c *qrpc.Call) {
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("repl", func(r qrpc.Responder, c *qrpc.Call) {
		var params DelegateParams
		_ = c.Decode(&params)
		// ^^ hmmm
		ch, err := r.Hijack(nil)
		if err != nil {
			log.Println(err)
		}
		repl.Run(ch, ch, map[string]interface{}{
			"Root": root,
		})
	})
	api.HandleFunc("readDelegate", func(r qrpc.Responder, c *qrpc.Call) {
		var params DelegateParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		src, err := workspace.DelegateSource(params.ID)
		if err != nil {
			r.Return(err)
			return
		}
		params.Contents = string(src)

		r.Return(params)
	})
	api.HandleFunc("writeDelegate", func(r qrpc.Responder, c *qrpc.Call) {
		var params DelegateParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		err = workspace.WriteDelegate(params.ID, []byte(params.Contents))
		if err != nil {
			r.Return(err)
			return
		}

		r.Return(nil)
	})
	api.HandleFunc("removeComponent", func(r qrpc.Responder, c *qrpc.Call) {
		var params RemoveComponentParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		n := root.FindID(params.ID)
		if n == nil {
			return
		}
		n.RemoveComponent(params.Component)
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("addDelegate", func(r qrpc.Responder, c *qrpc.Call) {
		var params NodeParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		n := root.FindID(params.ID)
		if n == nil {
			return
		}
		r.Return(workspace.CreateDelegate(n))
	})
	api.HandleFunc("updateNode", func(r qrpc.Responder, c *qrpc.Call) {
		var params NodeParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		n := root.FindID(params.ID)
		if n == nil {
			return
		}
		if params.Name != nil {
			n.Name = *params.Name
		}
		if params.Active != nil {
			n.Active = *params.Active
		}
		n.Sync()
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("callMethod", func(r qrpc.Responder, c *qrpc.Call) {
		var path string
		err := c.Decode(&path)
		if err != nil {
			r.Return(err)
			return
		}
		if path == "" {
			return
		}
		n := root.FindNode(path)
		localPath := path[len(n.FullPath())+1:]
		n.CallMethod(localPath)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("setExpression", func(r qrpc.Responder, c *qrpc.Call) {
		var params SetValueParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		n := root.FindNode(params.Path)
		localPath := params.Path[len(n.FullPath())+1:]
		n.SetExpression(localPath, params.Value.(string))
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("setValue", func(r qrpc.Responder, c *qrpc.Call) {
		var params SetValueParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		n := root.FindNode(params.Path)
		localPath := params.Path[len(n.FullPath())+1:]
		switch {
		case params.IntValue != nil:
			n.SetValue(localPath, *params.IntValue)
		case params.RefValue != nil:
			refNode := root.FindNode(*params.RefValue)
			if refNode != nil {
				comName := (*params.RefValue)[len(refNode.FullPath())+1:]
				c := refNode.Component(comName)
				// TODO: check interface
				if c != nil {
					n.SetValue(localPath, c)
				}
			}
		default:
			n.SetValue(localPath, params.Value)
		}
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("appendComponent", func(r qrpc.Responder, c *qrpc.Call) {
		var params AppendNodeParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		if params.Name == "" {
			return
		}
		p := root.FindID(params.ID)
		if p == nil {
			p = root
		}
		v := manifold.NewComponent(params.Name)
		p.AppendComponent(v)
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("deleteNode", func(r qrpc.Responder, c *qrpc.Call) {
		var id string
		err := c.Decode(&id)
		if err != nil {
			r.Return(err)
			return
		}
		if id == "" {
			return
		}
		n := root.FindID(id)
		n.Remove()
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("appendNode", func(r qrpc.Responder, c *qrpc.Call) {
		var params AppendNodeParams
		err := c.Decode(&params)
		if err != nil {
			r.Return(err)
			return
		}
		if params.Name == "" {
			return
		}
		p := root.FindID(params.ID)
		if p == nil {
			p = root
		}
		n := manifold.NewNode(params.Name)
		p.Append(n)
		n.Sync()
		exportNodes(&state, root)
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("subscribe", func(r qrpc.Responder, c *qrpc.Call) {
		clients[c.Caller] = "state"
		sendState()
		r.Return(nil)
	})
	api.HandleFunc("selectProject", func(r qrpc.Responder, c *qrpc.Call) {
		var name string
		err := c.Decode(&name)
		if err != nil {
			r.Return(err)
			return
		}
		state.CurrentProject = name
		sendState()
		r.Return(nil)
	})

	// sess, err := mux.DialWebsocket(addr)
	// if err != nil {
	// 	panic(err)
	// }
	// backend := &qrpc.Client{Session: sess, API: api}
	// _, err = backend.Call("register", []string{
	// 	"updateNode",
	// 	"callMethod",
	// 	"setExpression",
	// 	"setValue",
	// 	"appendComponent",
	// 	"deleteNode",
	// 	"appendNode",
	// 	"subscribe",
	// 	"selectProject",
	// }, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("connected to daemon")
	// backend.ServeAPI()
	// return nil

	// start server with api
	server := &qrpc.Server{}
	l, err := mux.ListenWebsocket(addr)
	if err != nil {
		panic(err)
	}
	log.Println("websocket server listening at", addr)
	return server.Serve(l, api)
}
