package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/progrium/duplex/golang"
)

// Logs going to stderr?

var rpc = duplex.NewRPC(duplex.NewJSONCodec())

func init() {
	rpc.Register("echo2", func(ch *duplex.Channel) error {
		var obj interface{}
		if _, err := ch.Recv(&obj); err != nil {
			return err
		}
		fmt.Printf("echo2: %#v\n", obj)
		return ch.Send(obj, false)
	})
}

type GorillaWSAdapter struct {
	*websocket.Conn
}

func (ws *GorillaWSAdapter) Read(p []byte) (int, error) {
	_, msg, err := ws.Conn.ReadMessage()
	return copy(p, msg), err
}

func (ws *GorillaWSAdapter) Write(p []byte) (int, error) {
	return len(p), ws.Conn.WriteMessage(websocket.TextMessage, p)
}

func (ws *GorillaWSAdapter) Close() error {
	return ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "ws://localhost:8000/"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	fatal(err)
	peer, err := rpc.Handshake(&GorillaWSAdapter{ws})
	fatal(err)
	defer peer.Close()

	fatal(peer.Call("register", []string{"echo2"}, nil))

	fmt.Println("routing...")
	peer.Route()
}
