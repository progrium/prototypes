package qrpc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"
	"sync"

	"github.com/progrium/prototypes/qrpc/transport"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type Call struct {
	Destination string
	ObjectPath  string // todo
	Method      string // todo
	Caller      Caller

	dec *msgpack.Decoder
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

func Dial(addr string, api *API) (*Client, error) {
	sess, err := transport.DialQuic(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		return nil, err
	}
	if api == nil {
		api = NewAPI()
	}
	return &Client{
		Session: sess,
		API:     api,
	}, nil
}

func (c *Client) Close() error {
	return c.Session.Close()
}

func (c *Client) ServeAPI() {
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
		resp.Return(errors.New("handler does not exist for this destination"))
		return
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

func (s *Server) ListenAndServe(addr string, api *API) error {
	listener, err := transport.ListenQuic(addr, generateTLSConfig(), nil)
	if err != nil {
		return err
	}
	if api != nil {
		s.API = api
	}
	if s.API == nil {
		s.API = NewAPI()
	}
	for {
		sess, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.ServeAPI(sess)
	}
}

func ExportFunc(fn interface{}) Handler {
	return HandlerFunc(func(r Responder, c *Call) {
		var obj interface{}
		err := c.Decode(&obj)
		if err != nil {
			panic(err)
		}
		params, ok := obj.([]interface{})
		if !ok {
			panic("only positional arguments in form of array are supported")
		}

		// invoke fn with reflect

		// look error type in return values
		// convert remaining multiple values to array

		//r.Return(ret)
	})
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}
