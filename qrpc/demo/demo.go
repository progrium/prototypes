package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/progrium/prototypes/qrpc"
	"github.com/progrium/prototypes/qrpc/transport"
	"golang.org/x/crypto/ssh"
)

const addr = "localhost:4242"

type EchoMessage struct {
	Message string
	Req     bool
	Resp    bool
}

type DemoService struct {
	Prefix string
}

func (s *DemoService) Upper(message string) string {
	return strings.ToUpper(s.Prefix + message)
}

func (s *DemoService) Echo(message string) string {
	return s.Prefix + message
}

func (s *DemoService) Err(message string) error {
	return errors.New(s.Prefix + message)
}

func simpleEcho(message string) (string, error) {
	log.Println("simple echo called")
	return message, nil
}

func main() {
	done := make(chan bool)
	api := qrpc.NewAPI()
	handler, err := qrpc.ExportFunc(simpleEcho)
	if err != nil {
		panic(err)
	}
	api.Handle("simple-echo", handler)
	handler, err = qrpc.Export(&DemoService{
		Prefix: "TEST: ",
	})
	if err != nil {
		panic(err)
	}
	api.Handle("demo", handler)
	api.HandleFunc("echo-client", func(r qrpc.Responder, c *qrpc.Call) {
		log.Println("echo-client called")
		var msg EchoMessage
		err := c.Decode(&msg)
		if err != nil {
			r.Return(err)
			return
		}
		msg.Req = false
		msg.Resp = true
		r.Return(msg)
	})
	api.HandleFunc("echo-server", func(r qrpc.Responder, c *qrpc.Call) {
		log.Println("echo-server called")
		var msg EchoMessage
		err := c.Decode(&msg)
		if err != nil {
			r.Return(err)
			return
		}
		msg.Req = false
		msg.Resp = true
		r.Return(msg)

		req := &EchoMessage{
			Message: "hello client",
			Req:     true,
		}
		fmt.Printf("req: %#v\n", req)
		var resp EchoMessage
		err = c.Caller.Call("echo-client", req, &resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("resp: %#v\n", resp)
		done <- true
	})

	// server

	//l, err := transport.ListenSSH(addr, generateSSHServerConfig())
	//l, err := transport.ListenQuic(addr, generateTLSConfig(), nil)
	//l, err := transport.ListenMuxado(addr, nil)
	l, err := transport.ListenTCP(addr)
	if err != nil {
		panic(err)
	}

	server := &qrpc.Server{}
	go func() {
		log.Fatal(server.Serve(l, api))
	}()

	// client

	//sess, err := transport.DialSSH(addr, generateSSHClientConfig())
	//sess, err := transport.DialQuic(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	//sess, err := transport.DialMuxado(addr, nil)
	sess, err := transport.DialTCP(addr)
	if err != nil {
		panic(err)
	}
	client := &qrpc.Client{
		Session: sess,
		API:     api,
	}
	go client.ServeAPI()

	flag.Parse()

	var resp string
	err = client.Call(flag.Arg(0), "Hello", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %#v\n", resp)
	//<-done
}

// ================== HELPERS ===================

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

func generateSSHClientConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:            "qrpc",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func generateSSHServerConfig() *ssh.ServerConfig {
	cfg := &ssh.ServerConfig{
		NoClientAuth: true,
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.NewSignerFromKey(key)
	if err != nil {
		panic(err)
	}
	cfg.AddHostKey(signer)
	return cfg
}
