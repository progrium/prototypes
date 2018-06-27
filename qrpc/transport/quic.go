package transport

import (
	"context"
	"crypto/tls"
	"io"
	"net"

	quic "github.com/lucas-clemente/quic-go"
)

type Session interface {
	Context() context.Context
	Close() error
	Open() (Channel, error)
	Accept() (Channel, error)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

type Channel interface {
	ID() uint64
	Context() context.Context
	io.Reader
	io.Writer
	io.Closer
}

type Listener interface {
	Close() error
	Addr() net.Addr
	Accept() (Session, error)
}

type quicSession struct {
	quic.Session
}

func (s *quicSession) Close() error {
	return s.Session.Close(nil)
}

func (s *quicSession) Open() (Channel, error) {
	stream, err := s.OpenStreamSync()
	return &quicChannel{stream}, err
}

func (s *quicSession) Accept() (Channel, error) {
	stream, err := s.AcceptStream()
	return &quicChannel{stream}, err
}

type quicChannel struct {
	quic.Stream
}

func (c *quicChannel) ID() uint64 {
	return uint64(c.StreamID())
}

func DialQuic(addr string, tlsConf *tls.Config, config *quic.Config) (Session, error) {
	sess, err := quic.DialAddr(addr, tlsConf, config)
	return &quicSession{sess}, err
}

func ListenQuic(addr string, tlsConf *tls.Config, config *quic.Config) (Listener, error) {
	listener, err := quic.ListenAddr(addr, tlsConf, config)
	return &quicListener{listener}, err
}

type quicListener struct {
	quic.Listener
}

func (l *quicListener) Accept() (Session, error) {
	sess, err := l.Listener.Accept()
	return &quicSession{sess}, err
}
