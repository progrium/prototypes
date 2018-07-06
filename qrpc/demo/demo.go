package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
	"golang.org/x/crypto/ssh"
)

const addr = "localhost:4242"

type Person struct {
	name string
	age  int
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) IncrAge() {
	p.age += 1
}

type PeopleService struct{}

func (_ *PeopleService) MakePerson(name string) {

}

func (_ *PeopleService) Echo(msg string) string {
	return msg
}

func main() {
	api := qrpc.NewAPI()
	handler, err := qrpc.Export(&PeopleService{})
	if err != nil {
		panic(err)
	}
	api.Handle("demo", handler)

	// server
	//l, err := mux.ListenSSH(addr, generateSSHServerConfig())
	//l, err := mux.ListenQuic(addr, generateTLSConfig(), nil)
	//l, err := mux.ListenMuxado(addr, nil)
	//l, err := mux.ListenTCP(addr)
	l, err := mux.ListenWebsocket(addr)
	if err != nil {
		panic(err)
	}

	server := &qrpc.Server{}
	go func() {
		log.Fatal(server.Serve(l, api))
	}()

	// client
	//sess, err := mux.DialSSH(addr, generateSSHClientConfig())
	//sess, err := mux.DialQuic(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	//sess, err := mux.DialMuxado(addr, nil)
	//sess, err := mux.DialTCP(addr)
	sess, err := mux.DialWebsocket(addr)
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
