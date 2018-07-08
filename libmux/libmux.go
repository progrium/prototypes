package main

import "C"

import (
	"bytes"
	"fmt"
	"io"
	"log"
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
	defer log.Println(err)
	for idx, vv := range m.values {
		if vv == nil {
			m.values[idx] = err
			return idx + 1
		}
	}
	m.values = append(m.values, err)
	return len(m.values)
}

func (m *refmanager) Err(id int) error {
	m.Lock()
	defer m.Unlock()
	if id-1 >= len(m.errors) || id < 1 {
		log.Println("bad errid:", id)
		return nil
	}
	err := m.errors[id-1]
	m.errors[id-1] = nil
	return err
}

var refs = &refmanager{}

var errNoValueFmt = "%s id has no value"
var errTypeFmt = "%s id has wrong type"

//export Error
func Error(id int, buf *C.uchar, len int) int {
	err := refs.Err(id)
	if err == nil {
		return 0
	}
	errBuf := bytes.NewBufferString(err.Error())
	b := (*[1 << 30]byte)(unsafe.Pointer(buf))[:len:len]
	return copy(b, errBuf.Bytes())
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
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Listener"))
	}
	l, ok := r.(mux.Listener)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "ListenerClose"))
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
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Listener"))
	}
	l, ok := r.(mux.Listener)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "ListenerAccept"))
	}
	sess, err := l.Accept()
	if err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return refs.StoreVal(sess)
}

//export SessionClose
func SessionClose(id int) int {
	r := refs.Val(id)
	if r == nil {
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Session"))
	}
	sess, ok := r.(mux.Session)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "SessionClose"))
	}
	if err := sess.Close(); err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return refs.ReleaseVal(id)
}

//export SessionOpen
func SessionOpen(id int) int {
	r := refs.Val(id)
	if r == nil {
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Session"))
	}
	l, ok := r.(mux.Session)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "SessionOpen"))
	}
	ch, err := l.Open()
	if err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return refs.StoreVal(ch)
}

//export SessionAccept
func SessionAccept(id int) int {
	r := refs.Val(id)
	if r == nil {
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Session"))
	}
	l, ok := r.(mux.Session)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "SessionAccept"))
	}
	ch, err := l.Accept()
	if err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return refs.StoreVal(ch)
}

//export ChannelClose
func ChannelClose(id int) int {
	r := refs.Val(id)
	if r == nil {
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Channel"))
	}
	ch, ok := r.(mux.Channel)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "ChannelClose"))
	}
	if err := ch.Close(); err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return refs.ReleaseVal(id)
}

//export ChannelWrite
func ChannelWrite(id int, buf *C.uchar, len int) int {
	r := refs.Val(id)
	if r == nil {
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Channel"))
	}
	ch, ok := r.(mux.Channel)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "ChannelWrite"))
	}
	b := (*[1 << 30]byte)(unsafe.Pointer(buf))[:len:len]
	n, err := ch.Write(b)
	if err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return n
}

//export ChannelRead
func ChannelRead(id int, buf *C.uchar, len int) int {
	r := refs.Val(id)
	if r == nil {
		return 0 //refs.StoreErr(fmt.Errorf(errNoValueFmt, "Channel"))
	}
	ch, ok := r.(mux.Channel)
	if !ok {
		return refs.StoreErr(fmt.Errorf(errTypeFmt, "ChannelRead"))
	}
	b := (*[1 << 30]byte)(unsafe.Pointer(buf))[:len:len]
	n, err := ch.Read(b)
	if err != nil {
		if err == io.EOF {
			return 0
		}
		return refs.StoreErr(err)
	}
	return n
}

func main() {}
