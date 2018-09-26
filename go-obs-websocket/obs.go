package obs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"sync"

	"golang.org/x/net/websocket"
)

const maxID = 1 << 63

type Message map[string]interface{}

type WebsocketConn struct {
	mu        sync.Mutex
	enc       *json.Encoder
	dec       *json.Decoder
	reqs      map[string]chan Message
	nextReqID uint64

	Events chan Message

	*websocket.Conn
}

func (c *WebsocketConn) Request(reqType string, msg map[string]string) (Message, error) {
	c.mu.Lock()
	if msg == nil {
		msg = map[string]string{}
	}
	msg["request-type"] = reqType
	msg["message-id"] = strconv.FormatUint(c.nextReqID, 10)
	respCh := make(chan Message)
	c.reqs[msg["message-id"]] = respCh
	if err := c.enc.Encode(msg); err != nil {
		c.mu.Unlock()
		return nil, err
	}
	c.nextReqID = (c.nextReqID + 1) % maxID
	c.mu.Unlock()
	resp := <-respCh // TODO: timeout, context
	return resp, nil
}

func (c *WebsocketConn) route() {
	for {
		resp := map[string]interface{}{}
		if err := c.dec.Decode(&resp); err != nil {
			return
		}
		v, ok := resp["message-id"]
		if !ok {
			c.Events <- resp
			continue
		}
		msgID, ok := v.(string)
		if !ok {
			continue
		}
		r, ok := c.reqs[msgID]
		if ok {
			delete(c.reqs, msgID)
			r <- resp
		}
	}

}

func generateAuth(msg Message, password string) string {
	sinput := sha256.Sum256([]byte(password + msg["salt"].(string)))
	secret := base64.StdEncoding.EncodeToString(sinput[:])
	rinput := sha256.Sum256([]byte(secret + msg["challenge"].(string)))
	return base64.StdEncoding.EncodeToString(rinput[:])
}

func DialWebsocket(url, password string) (*WebsocketConn, error) {
	origin := "http://localhost/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		return nil, err
	}
	conn := &WebsocketConn{
		Conn:   ws,
		enc:    json.NewEncoder(ws),
		dec:    json.NewDecoder(ws),
		reqs:   make(map[string]chan Message),
		Events: make(chan Message),
	}
	go conn.route()

	resp, err := conn.Request("GetAuthRequired", nil)
	if err != nil {
		return nil, err
	}

	if resp["authRequired"].(bool) {
		resp, err = conn.Request("Authenticate", map[string]string{
			"auth": generateAuth(resp, password),
		})
		if err != nil {
			return nil, err
		}
		if resp["status"].(string) != "ok" {
			return nil, errors.New("auth failed") // TODO: better error
		}
	}

	return conn, nil
}
