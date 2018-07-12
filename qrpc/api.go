package qrpc

import (
	"errors"
	"strings"
	"sync"

	"github.com/progrium/prototypes/libmux/mux"
)

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
	Proxy(path string) Caller
}

type API interface {
	Handle(path string, handler Handler)
	HandleFunc(path string, handler func(Responder, *Call))
	Handler(path string) Handler
	ServeAPI(sess mux.Session, ch mux.Channel, c Codec)
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

func (a *api) ServeAPI(sess mux.Session, ch mux.Channel, c Codec) {
	dec := c.Decoder(ch)
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
	resp := NewResponder(ch, c, header)
	handler := a.Handler(call.Destination)
	if handler == nil {
		resp.Return(errors.New("handler does not exist for this destination"))
		return
	}
	handler.ServeRPC(resp, &call)
}
