package supervisor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/afero"

	"github.com/armon/circbuf"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
)

type lockedBuffer struct {
	*circbuf.Buffer
	sync.Mutex
}

func (lb *lockedBuffer) Write(b []byte) (n int, err error) {
	lb.Lock()
	defer lb.Unlock()
	return lb.Buffer.Write(b)
}

func (lb *lockedBuffer) Bytes() []byte {
	lb.Lock()
	defer lb.Unlock()
	return lb.Buffer.Bytes()
}

type LineWriter struct {
	output io.Writer
	sync.Mutex
}

func (lw *LineWriter) WriteLine(left, right string, isError bool) {
	lw.Lock()
	defer lw.Unlock()

	var formatter string
	if isError {
		formatter = fmt.Sprintf("%%-%ds > ", 20)
	} else {
		formatter = fmt.Sprintf("%%-%ds | ", 20)
	}
	fmt.Fprintf(lw.output, formatter, left)
	fmt.Fprintln(lw.output, right)
}

func (lw *LineWriter) SystemOutput(str string) {
	lw.WriteLine("SYSTEM", str, false)
}

// TODO: pass in logs buffer and tee into it
func (lw *LineWriter) ReadLines(wg *sync.WaitGroup, name string, r io.Reader, isError bool) {
	defer wg.Done()

	reader := bufio.NewReader(r)

	var buffer bytes.Buffer

	for {
		buf := make([]byte, 1024)

		if n, err := reader.Read(buf); err != nil {
			return
		} else {
			buf = buf[:n]
		}

		for {
			i := bytes.IndexByte(buf, '\n')
			if i < 0 {
				break
			}
			buffer.Write(buf[0:i])
			lw.WriteLine(name, buffer.String(), isError)
			buffer.Reset()
			buf = buf[i+1:]
		}

		buffer.Write(buf)
	}
}

type Process struct {
	cmd    *exec.Cmd
	logs   *lockedBuffer
	exited chan int
	name   string
}

func StartProcess(cmd *exec.Cmd, lw *LineWriter) (*Process, error) {
	name := filepath.Base(cmd.Args[0])

	buf, err := circbuf.NewBuffer(1024 * 64)
	if err != nil {
		panic(err)
	}
	lbuf := &lockedBuffer{Buffer: buf}

	p := &Process{
		logs:   lbuf,
		cmd:    cmd,
		exited: make(chan int),
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go lw.ReadLines(pipeWait, name, io.TeeReader(stdout, lbuf), false)
	go lw.ReadLines(pipeWait, name, io.TeeReader(stderr, lbuf), true)

	err = cmd.Start()
	if err != nil {
		return p, err
	}

	go func() {
		pipeWait.Wait()
		status, err := exitStatus(cmd.Wait())
		if err != nil {
			panic(err)
		}
		p.exited <- status
	}()

	return p, nil
}

func exitStatus(err error) (int, error) {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// There is no platform independent way to retrieve
			// the exit code, but the following will work on Unix
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return int(status.ExitStatus()), nil
			}
		}
		return 0, err
	}
	return 0, nil
}

type Supervisor struct {
	sidecars     map[string]*Process
	output       *LineWriter
	cmdFactory   func(*exec.Cmd) *exec.Cmd
	restartDelay time.Duration
	maxRestarts  int
	wg           sync.WaitGroup
	fs           afero.Fs
	watcher      Watcher

	sync.Mutex
}

func NewSupervisor(output io.Writer) (*Supervisor, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Supervisor{
		sidecars: make(map[string]*Process),
		output:   &LineWriter{output: output},
		cmdFactory: func(cmd *exec.Cmd) *exec.Cmd {
			if cmd == nil {
				cmd = &exec.Cmd{}
			}
			return cmd
		},
		restartDelay: time.Duration(500 * time.Millisecond),
		maxRestarts:  -1,
		fs:           afero.NewOsFs(),
		watcher:      &wrappedWatcher{watcher},
	}, nil
}

func (s *Supervisor) Wait() {
	s.wg.Wait()
}

func (s *Supervisor) Logs(path string) []byte {
	s.Lock()
	defer s.Unlock()
	p, ok := s.sidecars[path]
	if !ok {
		return []byte{}
	}
	return p.logs.Bytes()
}

func (s *Supervisor) Watch() {
	for {
		select {
		case event := <-s.watcher.Events():
			switch {
			case event.Op&fsnotify.Write == fsnotify.Write:
				//log.Println("write:", event.Name)
				s.Reload(event.Name)
			case event.Op&fsnotify.Create == fsnotify.Create:
				//log.Println("create:", event.Name)
				s.Reload(event.Name)
			case event.Op&fsnotify.Remove == fsnotify.Remove:
				//log.Println("remove:", event.Name)
				s.Stop(event.Name)
			}
		case err := <-s.watcher.Errors():
			// remove me or make a system message
			log.Println("error:", err)
		}
	}
}

func (s *Supervisor) LoadDir(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return errors.Wrapf(err, "abs %s", path)
	}
	fi, err := s.fs.Stat(path)
	if err != nil {
		return errors.Wrapf(err, "stat %s", path)
	}
	if !fi.IsDir() {
		return errors.New("path is not a directory")
	}
	err = s.watcher.Add(path)
	if err != nil {
		return errors.Wrapf(err, "add %s", path)
	}
	dir, err := afero.ReadDir(s.fs, path)
	if err != nil {
		return errors.Wrapf(err, "readdir %s", path)
	}
	for _, fi := range dir {
		if !fi.IsDir() && fi.Mode()&0111 != 0 {
			err := s.Reload(filepath.Join(path, fi.Name()))
			if err != nil {
				// TODO REMOVE
				log.Println(err)
			}
		}
	}
	return nil
}

func (s *Supervisor) Reload(path string) error {
	if err := s.reloadable(path); err != nil {
		return err
	}
	return s.reload(path, 0)
}

func (s *Supervisor) reloadable(path string) error {
	s.Lock()
	defer s.Unlock()
	fi, err := s.fs.Stat(path)
	if err != nil {
		return err
	}
	if fi.IsDir() || fi.Mode()&0111 == 0 {
		return errors.New("path is not executable file")
	}
	return nil
}

func (s *Supervisor) dumpLogs(process *Process) {
	var (
		filename  = process.cmd.Args[0]
		extension = filepath.Ext(filename)
		path      = filename[0 : len(filename)-len(extension)]
		logpath   = fmt.Sprintf("%s-debug.log", path)
	)
	if err := afero.WriteFile(s.fs, logpath, process.logs.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func (s *Supervisor) Stop(path string) error {
	s.Lock()
	defer s.Unlock()
	p, ok := s.sidecars[path]
	if ok {
		delete(s.sidecars, path)
		p.cmd.Process.Kill()
	}
	return nil
}

func (s *Supervisor) Get(path string) (*os.Process, error) {
	s.Lock()
	defer s.Unlock()
	p, ok := s.sidecars[path]
	if !ok {
		return nil, errors.New("no process registered")
	}
	return p.cmd.Process, nil
}

func (s *Supervisor) reload(path string, retries int) (err error) {
	s.Lock()
	defer s.Unlock()
	p, ok := s.sidecars[path]
	if ok {
		p.cmd.Process.Kill()
	}
	cmd := s.cmdFactory(&exec.Cmd{
		Path: path,
		Args: []string{path},
		Env:  os.Environ(),
		Dir:  filepath.Dir(path),
	})
	s.sidecars[path], err = StartProcess(cmd, s.output)
	if err != nil {
		return
	}
	exited := s.sidecars[path].exited
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		status := <-exited
		s.Lock()
		_, ok = s.sidecars[path]
		s.Unlock()
		if !ok {
			return
		}
		if status > 0 {
			s.dumpLogs(s.sidecars[path])
		}
		if s.maxRestarts > -1 && retries >= s.maxRestarts {
			return
		}
		retries += 1
		time.Sleep(s.restartDelay)
		s.reload(path, retries)
	}()
	return nil
}

type Watcher interface {
	Close() error
	Add(string) error
	Events() chan fsnotify.Event
	Errors() chan error
}

type wrappedWatcher struct {
	*fsnotify.Watcher
}

func (w *wrappedWatcher) Events() chan fsnotify.Event {
	return w.Watcher.Events
}

func (w *wrappedWatcher) Errors() chan error {
	return w.Watcher.Errors
}
