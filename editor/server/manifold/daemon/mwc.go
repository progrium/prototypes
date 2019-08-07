package daemon

import (
	"fmt"
	"io"
	"sync"
)

type writeMeta struct {
	closer chan int64
	bytes  int64
}

func (wm writeMeta) close() {
	wm.closer <- wm.bytes
}

type MulticastWriteCloser struct {
	sync.Mutex
	writers map[io.Writer]writeMeta
	closed  bool
}

func NewMulticastWriteCloser() *MulticastWriteCloser {
	return &MulticastWriteCloser{
		writers: make(map[io.Writer]writeMeta),
	}
}

func (mw *MulticastWriteCloser) Write(p []byte) (n int, err error) {
	mw.Lock()
	defer mw.Unlock()
	if mw.closed {
		return 0, fmt.Errorf("closed")
	}
	n = len(p)
	for w, m := range mw.writers {
		n, err := w.Write(p)
		m.bytes += int64(n)
		if err != nil {
			m.close()
			delete(mw.writers, w)
		}
	}
	return
}

func (mw *MulticastWriteCloser) Close() error {
	mw.Lock()
	defer mw.Unlock()
	for _, m := range mw.writers {
		m.close()
	}
	mw.writers = make(map[io.Writer]writeMeta)
	mw.closed = true
	return nil
}

func (mw *MulticastWriteCloser) WriteTo(w io.Writer) (n int64, err error) {
	mw.Lock()
	if mw.closed {
		mw.Unlock()
		return 0, fmt.Errorf("closed")
	}
	if _, exists := mw.writers[w]; !exists {
		mw.writers[w] = writeMeta{make(chan int64), 0}
	}
	m := mw.writers[w]
	mw.Unlock()
	// TODO: error?
	return <-m.closer, nil
}

func (mw *MulticastWriteCloser) Count() int {
	mw.Lock()
	defer mw.Unlock()
	return len(mw.writers)
}
