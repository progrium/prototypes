package bus

import (
	"context"
	"io"
	"net"
	"time"
)

// multiplexed transport

type Session interface {
	Context() context.Context
	Close() error
	Open() Channel
	Accept() Channel
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

type Channel interface {
	ID()
	Context()
	io.Reader
	io.Writer
	io.Closer
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	SetDeadline(t time.Time) error
}

// rpc/bus

type BusClient interface {
	Session()

	Serve()
	Register()

	Invoke()
	//Proxy()
}

// Auth
// Export

func NewBusPeer(Session) BusPeer {

}
