package qrpc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/progrium/prototypes/qrpc/transport"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type Call struct {
	Destination string
	ObjectPath  string
	Method      string
	Caller      Caller

	dec *msgpack.Decoder
}

func (c *Call) parse() {
	parts := strings.Split(c.Destination, "/")
	c.ObjectPath = strings.Join(parts[0:len(parts)-2], "/")
	c.Method = parts[len(parts)-1]
}

func (c *Call) Decode(v interface{}) error {
	return c.dec.Decode(v)
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
	Session transport.Session
	API     *API
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
		go c.API.Serve(c.Session, ch)
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
	return dec.Decode(reply)
}

type API struct {
	handlers map[string]Handler
	mu       sync.Mutex
}

func NewAPI() *API {
	return &API{
		handlers: make(map[string]Handler),
	}
}

func (a *API) Handle(path string, handler Handler) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.handlers[path] = handler
}

func (a *API) HandleFunc(path string, handler func(Responder, *Call)) {
	a.Handle(path, HandlerFunc(handler))
}

func (a *API) Serve(sess transport.Session, ch transport.Channel) {
	dec := msgpack.NewDecoder(ch)
	var call Call
	err := dec.Decode(&call)
	if err != nil {
		panic(err)
	}
	call.parse()
	call.dec = dec
	call.Caller = &Client{
		Session: sess,
	}
	header := &ResponseHeader{}
	resp := &responder{ch, header}
	a.mu.Lock()
	handler, exists := a.handlers[call.Destination]
	a.mu.Unlock()
	if !exists {
		a.mu.Lock()
		for k, v := range a.handlers {
			if strings.HasPrefix(call.Destination, k) {
				handler = v
				break
			}
		}
		a.mu.Unlock()
		if handler == nil {
			resp.Return(errors.New("handler does not exist for this destination"))
			return
		}
	}
	handler.ServeRPC(resp, &call)
}

type responder struct {
	ch     transport.Channel
	header *ResponseHeader
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
	API *API
}

func (s *Server) ServeAPI(sess transport.Session) {
	for {
		ch, err := sess.Accept()
		if err != nil {
			panic(err)
		}
		go s.API.Serve(sess, ch)
	}
}

func (s *Server) Serve(l transport.Listener, api *API) error {
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
	if rcvr != nil {
		if reflectedType.NumIn() != 2 {
			return nil, fmt.Errorf("only supports 1 argument atm, got %d", reflectedType.NumIn())
		}
	} else {
		if reflectedType.NumIn() != 1 {
			return nil, fmt.Errorf("only supports 1 argument atm, got %d", reflectedType.NumIn())
		}
	}
	if reflectedType.NumOut() > 2 {
		return nil, fmt.Errorf("only supports up to 1 return value and optional error")
	}

	var paramType reflect.Type
	if rcvr != nil {
		paramType = reflectedType.In(1)
	} else {
		paramType = reflectedType.In(0)
	}
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()

	return HandlerFunc(func(r Responder, c *Call) {
		var paramValue reflect.Value
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

		var params []reflect.Value
		if rcvr != nil {
			params = append(params, reflect.ValueOf(rcvr))
		}
		if paramType.Kind() == reflect.Ptr {
			params = append(params, paramValue)
		} else {
			params = append(params, paramValue.Elem())
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
