package mux

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/progrium/prototypes/qmux"
	"golang.org/x/net/websocket"
)

func DialWebsocket(addr string) (Session, error) {
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/", addr), "", fmt.Sprintf("http://%s/", addr))
	return &qmuxSession{
		Session: qmux.NewSession(ws),
		ctx:     context.Background(),
	}, err
}

func ListenWebsocket(addr string) (Listener, error) {
	sessCh := make(chan qmux.Session)
	listener, err := net.Listen("tcp", addr)
	s := &http.Server{
		Addr: addr,
		Handler: websocket.Handler(func(ws *websocket.Conn) {
			sess := qmux.NewSession(ws)
			sessCh <- sess
			sess.Wait()
		}),
	}
	go s.Serve(listener)
	return &websocketListener{listener, sessCh}, err
}

type websocketListener struct {
	net.Listener
	sessCh chan qmux.Session
}

func (l *websocketListener) Accept() (Session, error) {
	return &qmuxSession{
		Session: <-l.sessCh,
		ctx:     context.Background(),
	}, nil
}
