package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/progrium/duplex/golang"
	"github.com/progrium/prototypes/supervisor"
)

// toDO: rpcdaemon possible halts when reloading a program that exits

var rpc = duplex.NewRPC(duplex.NewJSONCodec())

var backendTable map[string]*HubConn
var frontendTable map[int]*Channel
var frontendMutex sync.Mutex
var idCounter = 0

type Channel struct {
	Backend  *HubConn
	Frontend *HubConn
	OrigID   int
	NewID    int
}

func init() {
	backendTable = make(map[string]*HubConn)
	frontendTable = make(map[int]*Channel)
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

var upgrader = websocket.Upgrader{} // use default options

func WebsocketServer(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	conn, err := Accept(&GorillaWSAdapter{ws})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	channels := make(map[int]*Channel)
	for {
		msg, err := conn.RecvMsg()
		if err != nil {
			//panic(err)
			log.Println(err)
			return
		}
		if msg.Method == "register" {
			iter, ok := msg.Payload.([]interface{})
			if !ok {
				continue
			}
			for _, v := range iter {
				backendTable[v.(string)] = conn
			}
			conn.SendMsg(&duplex.Message{
				Id:   msg.Id,
				Type: duplex.TypeReply,
			})
			continue
		}
		log.Println(msg)
		switch msg.Type {
		case duplex.TypeReply:
			frontendMutex.Lock()
			ch, exists := frontendTable[msg.Id]
			if !exists {
				panic("channel id not found")
			}
			if !msg.More {
				delete(frontendTable, ch.NewID)
			}
			frontendMutex.Unlock()
			msg.Id = ch.OrigID
			if err := ch.Frontend.SendMsg(msg); err != nil {
				panic(err)
			}
		case duplex.TypeRequest:
			frontendMutex.Lock()
			ch, exists := frontendTable[msg.Id]
			if !exists {
				backend, ok := backendTable[msg.Method]
				if !ok {
					panic("method not found")
				}
				idCounter += 1
				ch = &Channel{
					Frontend: conn,
					Backend:  backend,
					OrigID:   msg.Id,
					NewID:    idCounter,
				}
				frontendTable[ch.NewID] = ch
				channels[ch.OrigID] = ch
			}
			if !msg.More {
				delete(channels, ch.OrigID)
			}
			frontendMutex.Unlock()
			msg.Id = ch.NewID
			if err := ch.Backend.SendMsg(msg); err != nil {
				panic(err)
			}
		default:
			panic("message type not supported")
		}

	}
}

func Accept(conn io.ReadWriteCloser) (*HubConn, error) {
	hconn := &HubConn{duplex.NewJSONCodec(), conn}
	buf := make([]byte, 32)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	// TODO: check handshake
	_, err = conn.Write([]byte(duplex.HandshakeAccept))
	if err != nil {
		return nil, err
	}
	return hconn, nil
}

type HubConn struct {
	codec *duplex.Codec

	io.ReadWriteCloser
}

func (conn *HubConn) SendMsg(msg *duplex.Message) error {
	frame, err := conn.codec.Encode(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(frame)
	return err
}

func (conn *HubConn) RecvMsg() (*duplex.Message, error) {
	frame := make([]byte, duplex.MaxFrameSize)
	n, err := conn.Read(frame)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("some error case thats not important because im rewriting this")
	}
	var msg duplex.Message
	err = conn.codec.Decode(frame[:n], &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func triggerHook(name string, args ...string) {
	path := fmt.Sprintf("hooks/%s", name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}
	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("%s: %s\n", path, err)
	}
}

func main() {
	http.HandleFunc("/", WebsocketServer)
	fmt.Println("Listening on 8000...")
	go http.ListenAndServe(":8000", nil)

	s, err := supervisor.NewSupervisor(os.Stdout)
	s.ChangeCallback = func(path string, reloadable bool, deleted bool) {
		if !deleted {
			triggerHook("change", path)
		}
	}
	if err != nil {
		panic(err)
	}
	go s.Watch()
	err = s.LoadDir("modules")
	if err != nil {
		panic(err)
	}
	s.Wait()
}
