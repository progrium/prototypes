package transport

import (
	"context"
	"net"

	"github.com/progrium/prototypes/qmux"
)

func DialTCP(addr string) (Session, error) {
	conn, err := net.Dial("tcp", addr)
	return &qmuxSession{
		Session: qmux.NewSession(conn),
		ctx:     context.Background(),
	}, err
}

func ListenTCP(addr string) (Listener, error) {
	listener, err := net.Listen("tcp", addr)
	return &qmuxListener{listener}, err
}

type qmuxSession struct {
	qmux.Session

	ctx context.Context
}

func (s *qmuxSession) Context() context.Context {
	return s.ctx
}

func (s *qmuxSession) Open() (Channel, error) {
	ch, err := s.Session.Open()
	return &qmuxChannel{
		Channel: ch,
		ctx:     s.ctx,
	}, err
}

func (s *qmuxSession) Accept() (Channel, error) {
	ch, err := s.Session.Accept()
	return &qmuxChannel{
		Channel: ch,
		ctx:     s.ctx,
	}, err
}

type qmuxChannel struct {
	qmux.Channel
	ctx context.Context
}

func (c *qmuxChannel) Context() context.Context {
	return c.ctx
}

func (c *qmuxChannel) ID() uint64 {
	return uint64(c.ID())
}

type qmuxListener struct {
	net.Listener
}

func (l *qmuxListener) Accept() (Session, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return &qmuxSession{
		Session: qmux.NewSession(conn),
		ctx:     context.Background(),
	}, err
}
