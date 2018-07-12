package qrpc

import (
	"fmt"
	"io"
	path_ "path"
	"strings"

	"github.com/progrium/prototypes/libmux/mux"
)

type Call struct {
	Destination string
	ObjectPath  string
	Method      string
	Caller      Caller
	Decoder     Decoder
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
	c.ObjectPath = strings.Join(parts[0:len(parts)-1], "/")
	c.Method = parts[len(parts)-1]
	return nil
}

func (c *Call) Decode(v interface{}) error {
	return c.Decoder.Decode(v)
}

type ResponseHeader struct {
	Error error
}

type responder struct {
	ch     mux.Channel
	header *ResponseHeader
	c      Codec
}

func NewResponder(ch mux.Channel, c Codec, header *ResponseHeader) Responder {
	return &responder{
		ch:     ch,
		c:      c,
		header: header,
	}
}

func (r *responder) Header() *ResponseHeader {
	return r.header
}

func (r *responder) Return(v interface{}) error {
	enc := r.c.Encoder(r.ch)
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

type Client struct {
	Session mux.Session
	API     API
	Codec   Codec
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
		go c.API.ServeAPI(c.Session, ch, &frameCodec{c.Codec})
	}
}

func (c *Client) Proxy(path string) Caller {
	return &proxy{path: path, client: c}
}

func (c *Client) Call(path string, args, reply interface{}) error {
	ch, err := c.Session.Open()
	if err != nil {
		return err
	}
	defer ch.Close()
	codec := newFrameCodec(c.Codec)
	// request
	enc := codec.Encoder(ch)
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
	dec := codec.Decoder(ch)
	var resp ResponseHeader
	err = dec.Decode(&resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf("remote: %s", resp.Error)
	}
	if reply != nil {
		return dec.Decode(reply)
	}
	return nil
}

type proxy struct {
	path   string
	client *Client
}

func (p *proxy) Call(path string, args, reply interface{}) error {
	return p.client.Call(path_.Join(p.path, path), args, reply)
}

func (p *proxy) Proxy(path string) Caller {
	return &proxy{path: path_.Join(p.path, path), client: p.client}
}

type Server struct {
	API   API
	Codec Codec
}

func (s *Server) ServeAPI(sess mux.Session) {
	for {
		ch, err := sess.Accept()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		go s.API.ServeAPI(sess, ch, newFrameCodec(s.Codec))
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
