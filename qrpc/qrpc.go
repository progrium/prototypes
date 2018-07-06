package qrpc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/progrium/prototypes/libmux/mux"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type Call struct {
	Destination string
	ObjectPath  string
	Method      string
	Caller      Caller
	Decoder     *msgpack.Decoder
}

func (c *Call) Parse() error {
	if len(c.Destination) == 0 {
		return fmt.Errorf("no destination specified")
	}
	if c.Destination[0] == '/' {
		c.Destination = c.Destination[1:]
	}
	parts := strings.Split(c.Destination, "/")
	if len(parts) == 1 {
		c.ObjectPath = "/"
		c.Method = parts[0]
		return nil
	}
	c.ObjectPath = strings.Join(parts[0:len(parts)-2], "/")
	c.Method = parts[len(parts)-1]
	return nil
}

func (c *Call) Decode(v interface{}) error {
	return c.Decoder.Decode(v)
}

type ResponseHeader struct {
	Error error
}

type Responder interface {
	Header() *ResponseHeader
	Return(interface{}) error
}

type Handler interface {
	ServeRPC(Responder, *Call)
}

type HandlerFunc func(Responder, *Call)

func (f HandlerFunc) ServeRPC(resp Responder, call *Call) {
	f(resp, call)
}

type Caller interface {
	Call(path string, args, reply interface{}) error
}

type Client struct {
	Session mux.Session
	API     API
}

func (c *Client) Close() error {
	return c.Session.Close()
}

func (c *Client) ServeAPI() {
	if c.API == nil {
		c.API = NewAPI()
	}
	for {
		ch, err := c.Session.Accept()
		if err != nil {
			panic(err)
		}
		go c.API.ServeAPI(c.Session, ch)
	}
}

func (c *Client) Call(path string, args, reply interface{}) error {
	ch, err := c.Session.Open()
	if err != nil {
		return err
	}
	defer ch.Close()
	// request
	enc := msgpack.NewEncoder(ch)
	err = enc.Encode(Call{
		Destination: path,
	})
	if err != nil {
		return err
	}
	err = enc.Encode(args)
	if err != nil {
		return err
	}
	// response
	dec := msgpack.NewDecoder(ch)
	var resp ResponseHeader
	err = dec.Decode(&resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}
	if reply != nil {
		return dec.Decode(reply)
	}
	return nil
}

type API interface {
	Handle(path string, handler Handler)
	HandleFunc(path string, handler func(Responder, *Call))
	Handler(path string) Handler
	ServeAPI(sess mux.Session, ch mux.Channel)
}

type api struct {
	handlers map[string]Handler
	mu       sync.Mutex
}

func NewAPI() *api {
	return &api{
		handlers: make(map[string]Handler),
	}
}

func (a *api) Handle(path string, handler Handler) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.handlers[path] = handler
}

func (a *api) HandleFunc(path string, handler func(Responder, *Call)) {
	a.Handle(path, HandlerFunc(handler))
}

func (a *api) Handler(path string) Handler {
	var handler Handler
	a.mu.Lock()
	for k, v := range a.handlers {
		if strings.HasPrefix(path, k) {
			handler = v
			break
		}
	}
	a.mu.Unlock()
	return handler
}

func (a *api) ServeAPI(sess mux.Session, ch mux.Channel) {
	dec := msgpack.NewDecoder(ch)
	var call Call
	err := dec.Decode(&call)
	if err != nil {
		panic(err)
	}
	err = call.Parse()
	if err != nil {
		panic(err)
	}
	call.Decoder = dec
	call.Caller = &Client{
		Session: sess,
	}
	header := &ResponseHeader{}
	resp := NewResponder(ch, header)
	handler := a.Handler(call.Destination)
	if handler == nil {
		resp.Return(errors.New("handler does not exist for this destination"))
		return
	}
	handler.ServeRPC(resp, &call)
}

type responder struct {
	ch     mux.Channel
	header *ResponseHeader
}

func NewResponder(ch mux.Channel, header *ResponseHeader) Responder {
	return &responder{ch, header}
}

func (r *responder) Header() *ResponseHeader {
	return r.header
}

func (r *responder) Return(v interface{}) error {
	enc := msgpack.NewEncoder(r.ch)
	var e error
	var ok bool
	if e, ok = v.(error); ok {
		v = nil
	}
	r.header.Error = e
	err := enc.Encode(r.header)
	if err != nil {
		return err
	}
	err = enc.Encode(v)
	if err != nil {
		return err
	}
	return r.ch.Close()
}

type Server struct {
	API API
}

func (s *Server) ServeAPI(sess mux.Session) {
	for {
		ch, err := sess.Accept()
		if err != nil {
			panic(err)
		}
		go s.API.ServeAPI(sess, ch)
	}
}

func (s *Server) Serve(l mux.Listener, api API) error {
	if api != nil {
		s.API = api
	}
	if s.API == nil {
		s.API = NewAPI()
	}
	for {
		sess, err := l.Accept()
		if err != nil {
			return err
		}
		go s.ServeAPI(sess)
	}
}

func ExportFunc(fn interface{}) (Handler, error) {
	return exportFunc(fn, nil)
}

func exportFunc(fn interface{}, rcvr interface{}) (Handler, error) {
	reflectedFn := reflect.ValueOf(fn)

	reflectedType := reflect.TypeOf(fn)
	if reflectedType.Kind() != reflect.Func {
		return nil, fmt.Errorf("takes only a function")
	}
	var hasParam bool
	if rcvr != nil {
		if reflectedType.NumIn() > 2 {
			return nil, fmt.Errorf("only supports 1 argument atm, got %d", reflectedType.NumIn())
		}
		hasParam = reflectedType.NumIn() > 1
	} else {
		if reflectedType.NumIn() > 1 {
			return nil, fmt.Errorf("only supports 1 argument atm, got %d", reflectedType.NumIn())
		}
		hasParam = reflectedType.NumIn() > 0
	}
	if reflectedType.NumOut() > 2 {
		return nil, fmt.Errorf("only supports up to 1 return value and optional error")
	}

	var paramType reflect.Type
	if hasParam {
		if rcvr != nil {
			paramType = reflectedType.In(1)
		} else {
			paramType = reflectedType.In(0)
		}
	} else {
		var empty interface{}
		paramType = reflect.TypeOf(empty)
	}
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()

	return HandlerFunc(func(r Responder, c *Call) {
		var paramValue reflect.Value
		if hasParam {
			if paramType.Kind() == reflect.Ptr {
				paramValue = reflect.New(paramType.Elem())
			} else {
				paramValue = reflect.New(paramType)
			}

			err := c.Decode(paramValue.Interface())
			if err != nil {
				// arguments weren't what was expected,
				// or any other error
				panic(err)
			}
		}

		var params []reflect.Value
		if rcvr != nil {
			params = append(params, reflect.ValueOf(rcvr))
		}
		if hasParam {
			if paramType.Kind() == reflect.Ptr {
				params = append(params, paramValue)
			} else {
				params = append(params, paramValue.Elem())
			}
		}
		retVals := reflectedFn.Call(params)

		if len(retVals) == 0 {
			r.Return(nil)
			return
		}

		// assuming up to 2 return values, one being an error
		var retVal reflect.Value
		for _, v := range retVals {
			if v.Type().Implements(errorInterface) {
				if !v.IsNil() {
					r.Return(v.Interface().(error))
					return
				}
			} else {
				retVal = v
			}
		}
		r.Return(retVal.Interface())

	}), nil
}

func Export(v interface{}) (Handler, error) {
	methodHandlers := make(map[string]Handler)
	reflectedType := reflect.TypeOf(v)
	for i := 0; i < reflectedType.NumMethod(); i++ {
		rmethod := reflectedType.Method(i)
		handler, err := exportFunc(rmethod.Func.Interface(), v)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", rmethod.Name, err.Error())
		}
		methodHandlers[rmethod.Name] = handler
	}
	return HandlerFunc(func(r Responder, c *Call) {
		handler, ok := methodHandlers[c.Method]
		if !ok {
			r.Return(errors.New("method handler does not exist for this destination"))
			return
		}
		handler.ServeRPC(r, c)
	}), nil
}

// TODO: object manager
type ObjectManager interface {
	Register(v interface{}) ManagedObject
	ServeRPC(r Responder, c *Call)
	Export(api *API, path string)
}

type ManagedObject interface {
	Dispose()
	Path()
	Value()
}
