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

type Request struct {
	Destination string
}

type Response struct {
	Error error
}

type Destinations []Destination

type Destination struct {
	Path    string
	Handler Handler
}

type Handler func(req interface{}) (interface{}, error)

type Peer interface {
	Session() transport.Session
	Register(path string, obj interface{})
	Call(path string, args, reply interface{}) error
	Close() error
}

type peer struct {
	session  transport.Session
	handlers map[string]Handler
	mu       sync.Mutex
}

func ListenAndServe(addr string, dests Destinations) error {
	listener, err := transport.ListenQuic(addr, generateTLSConfig(), nil)
	if err != nil {
		return err
	}
	for {
		sess, err := listener.Accept()
		if err != nil {
			return err
		}
		peer := NewPeer(sess)
		for _, v := range dests {
			peer.Register(v.Path, v.Handler)
		}
		go peer.Serve()
	}
}

func DialPeer(addr string) (*peer, error) {
	sess, err := transport.DialQuic(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		return nil, err
	}
	return NewPeer(sess), nil
}

func NewPeer(sess transport.Session) *peer {
	return &peer{
		session:  sess,
		handlers: make(map[string]Handler),
	}
}

func (p *peer) Session() transport.Session {
	return p.session
}

func (p *peer) Close() error {
	return p.session.Close()
}

func (p *peer) Register(path string, handler Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.handlers[path] = handler
}

func (p *peer) Call(path string, args, reply interface{}) error {
	ch, err := p.session.Open()
	if err != nil {
		return err
	}
	defer ch.Close()
	// request
	enc := msgpack.NewEncoder(ch)
	err = enc.Encode(Request{
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
	var resp Response
	err = dec.Decode(&resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}
	return dec.Decode(reply)
}

func (p *peer) sendResponse(ch transport.Channel, value interface{}, err error) error {
	enc := msgpack.NewEncoder(ch)
	e := enc.Encode(Response{
		Error: err,
	})
	if e != nil {
		return e
	}
	err = enc.Encode(value)
	if e != nil {
		return e
	}
	return ch.Close()
}

func (p *peer) Serve() {
	for {
		ch, err := p.session.Accept()
		if err != nil {
			panic(err)
		}
		dec := msgpack.NewDecoder(ch)
		var req Request
		err = dec.Decode(&req)
		if err != nil {
			panic(err)
		}
		p.mu.Lock()
		handler, exists := p.handlers[req.Destination]
		p.mu.Unlock()
		if !exists {
			p.sendResponse(ch, nil, errors.New("handler does not exist for this destination"))
			continue
		}
		var payload interface{}
		err = dec.Decode(&payload)
		if err != nil {
			panic(err)
		}
		reply, err := handler(payload)
		if err != nil {
			p.sendResponse(ch, nil, err)
			continue
		}
		p.sendResponse(ch, reply, nil)
	}
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
