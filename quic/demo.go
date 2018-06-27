package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"sync"

	quic "github.com/lucas-clemente/quic-go"
)

const addr = "localhost:4242"

const message = "foobar"

// We start a server echoing data on the first stream the client opens,
// then connect with a client, send the message, and wait for its receipt.
func main() {
	go echoListener()
	clientDialer()
}

func echoServer(name string, sess quic.Session) error {
	stream, err := sess.AcceptStream()
	if err != nil {
		return err
	}
	// Echo through the loggingWriter
	_, err = io.Copy(loggingWriter{stream, name}, stream)
	return err
}

func echoClient(name string, session quic.Session) error {
	stream, err := session.OpenStreamSync()
	if err != nil {
		return err
	}

	fmt.Printf("%s: Sending '%s'\n", name, message)
	_, err = stream.Write([]byte(message))
	if err != nil {
		return err
	}

	buf := make([]byte, len(message))
	_, err = io.ReadFull(stream, buf)
	if err != nil {
		return err
	}
	fmt.Printf("%s: Received '%s'\n", name, buf)
	stream.Close()
	return nil
}

// Start a server that echos all data on the first stream opened by the client
func echoListener() {
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		log.Fatal(err)
	}
	sess, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := echoClient("server", sess)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := echoServer("server", sess)
		if err != nil {
			log.Fatal(err)
		}
	}()
	//fmt.Println("listener wait")
	wg.Wait()
}

func clientDialer() {
	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := echoClient("client", session)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := echoServer("client", session)
		if err != nil {
			log.Fatal(err)
		}
	}()
	//fmt.Println("dialer wait")
	wg.Wait()
}

// A wrapper for io.Writer that also logs the message.
type loggingWriter struct {
	io.Writer
	name string
}

func (w loggingWriter) Write(b []byte) (int, error) {
	fmt.Printf("%s: Echoing '%s'\n", w.name, string(b))
	return w.Writer.Write(b)
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
