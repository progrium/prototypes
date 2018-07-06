package main

import "C"

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/progrium/prototypes/libmux/mux"
)

type refmanager struct {
	sync.Mutex
	values []interface{}
	errors []error
}

func (m *refmanager) StoreVal(v interface{}) int {
	m.Lock()
	defer m.Unlock()
	for idx, vv := range m.values {
		if vv == nil {
			m.values[idx] = v
			return idx + 1
		}
	}
	m.values = append(m.values, v)
	return len(m.values)
}

func (m *refmanager) ReleaseVal(id int) int {
	m.Lock()
	defer m.Unlock()
	if id-1 >= len(m.values) || id < 1 {
		return 0
	}
	m.values[id-1] = nil
	return 0
}

func (m *refmanager) Val(id int) interface{} {
	m.Lock()
	defer m.Unlock()
	if id-1 >= len(m.values) || id < 1 {
		return nil
	}
	return m.values[id-1]
}

func (m *refmanager) StoreErr(err error) int {
	m.Lock()
	defer m.Unlock()
	for idx, vv := range m.values {
		if vv == nil {
			m.values[idx] = err
			return (idx + 1) * -1
		}
	}
	m.values = append(m.values, err)
	return len(m.values) * -1
}

func (m *refmanager) Err(id int) error {
	m.Lock()
	defer m.Unlock()
	if (id*-1)-1 >= len(m.errors) || id > -1 {
		return nil
	}
	err := m.errors[(id*-1)-1]
	m.errors[(id*-1)-1] = nil
	return err
}

var refs = &refmanager{}

var errNoValue = errors.New("ref id has no value")
var errType = errors.New("ref id has wrong type")

//export Error
func Error(id int) string {
	err := refs.Err(id)
	if err == nil {
		return ""
	}
	return err.Error()
}

//export DialTCP
func DialTCP(addr string) int {
	sess, err := mux.DialTCP(addr)
	if err != nil {
		return -1
	}
	return refs.StoreVal(sess)
}

//export ListenTCP
func ListenTCP(addr string) int {
	l, err := mux.ListenTCP(addr)
	if err != nil {
		return -1
	}
	return refs.StoreVal(l)
}

//export DialWebsocket
func DialWebsocket(addr string) int {
	sess, err := mux.DialWebsocket(addr)
	if err != nil {
		return -1
	}
	return refs.StoreVal(sess)
}

//export ListenWebsocket
func ListenWebsocket(addr string) int {
	l, err := mux.ListenWebsocket(addr)
	if err != nil {
		return -1
	}
	return refs.StoreVal(l)
}

//export ListenerClose
func ListenerClose(id int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	l, ok := r.(mux.Listener)
	if !ok {
		return refs.StoreErr(errType)
	}
	if err := l.Close(); err != nil {
		return refs.StoreErr(err)
	}
	return refs.ReleaseVal(id)
}

//export ListenerAccept
func ListenerAccept(id int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	l, ok := r.(mux.Listener)
	if !ok {
		return refs.StoreErr(errType)
	}
	sess, err := l.Accept()
	if err != nil {
		return refs.StoreErr(err)
	}
	return refs.StoreVal(sess)
}

//export SessionClose
func SessionClose(id int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	sess, ok := r.(mux.Session)
	if !ok {
		return refs.StoreErr(errType)
	}
	if err := sess.Close(); err != nil {
		return refs.StoreErr(err)
	}
	return refs.ReleaseVal(id)
}

//export SessionOpen
func SessionOpen(id int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	l, ok := r.(mux.Session)
	if !ok {
		return refs.StoreErr(errType)
	}
	ch, err := l.Open()
	if err != nil {
		return refs.StoreErr(err)
	}
	return refs.StoreVal(ch)
}

//export SessionAccept
func SessionAccept(id int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	l, ok := r.(mux.Session)
	if !ok {
		return refs.StoreErr(errType)
	}
	ch, err := l.Accept()
	if err != nil {
		return refs.StoreErr(err)
	}
	return refs.StoreVal(ch)
}

//export ChannelClose
func ChannelClose(id int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	ch, ok := r.(mux.Channel)
	if !ok {
		return refs.StoreErr(errType)
	}
	if err := ch.Close(); err != nil {
		return refs.StoreErr(err)
	}
	return refs.ReleaseVal(id)
}

//export ChannelWrite
func ChannelWrite(id int, buf *C.uchar, len int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	ch, ok := r.(mux.Channel)
	if !ok {
		return refs.StoreErr(errType)
	}
	b := (*[1 << 30]byte)(unsafe.Pointer(buf))[:len:len]
	n, err := ch.Write(b)
	if err != nil {
		return refs.StoreErr(err)
	}
	return n
}

//export ChannelRead
func ChannelRead(id int, buf *C.uchar, len int) int {
	r := refs.Val(id)
	if r == nil {
		return refs.StoreErr(errNoValue)
	}
	ch, ok := r.(mux.Channel)
	if !ok {
		return refs.StoreErr(errType)
	}
	b := (*[1 << 30]byte)(unsafe.Pointer(buf))[:len:len]
	n, err := ch.Read(b)
	if err != nil {
		return refs.StoreErr(err)
	}
	return n
}

func main() {}
