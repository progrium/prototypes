package supervisor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
)

func helperCommand(cmd *exec.Cmd) *exec.Cmd {
	cs := []string{cmd.Path, "-test.run=TestHelperProcess", "--"}
	cs = append(cs, cmd.Path)
	cs = append(cs, cmd.Args...)
	cmd.Path = os.Args[0]
	cmd.Args = cs
	cmd.Env = append(cmd.Env, "GO_WANT_HELPER_PROCESS=1")
	return cmd
}

func helperCommandPath(s ...string) (cmd *exec.Cmd) {
	return helperCommand(&exec.Cmd{Path: s[0], Args: s[1:]})
}

type mockWatcher struct {
	events chan fsnotify.Event
	errors chan error
}

func (w *mockWatcher) Close() error {
	return nil
}

func (w *mockWatcher) Add(path string) error {
	return nil
}

func (w *mockWatcher) Events() chan fsnotify.Event {
	return w.events
}

func (w *mockWatcher) Errors() chan error {
	return w.errors
}

type Line struct {
	Header  string
	Message string
	Error   bool
}

func (l Line) String() string {
	if len(l.Header) == 0 && len(l.Message) == 0 {
		return ""
	}
	sep := "|"
	if l.Error {
		sep = ">"
	}
	return fmt.Sprintf("%s %s %s", l.Header, sep, l.Message)
}

func NewLine(line string) Line {
	outline := strings.SplitN(line, "| ", 2)
	if len(outline) > 1 {
		return Line{
			Header:  outline[0],
			Message: outline[1],
			Error:   false,
		}
	}
	errline := strings.SplitN(line, "> ", 2)
	if len(errline) > 1 {
		return Line{
			Header:  errline[0],
			Message: errline[1],
			Error:   true,
		}
	}
	return Line{Message: line}
}

func ParseLines(buf *bytes.Buffer) []Line {
	lines := []Line{}
	for _, line := range strings.Split(buf.String(), "\n") {
		if line != "" {
			lines = append(lines, NewLine(line))
		}
	}
	return lines
}

func TestProcessStdout(t *testing.T) {
	buf := new(bytes.Buffer)
	lw := &LineWriter{output: buf}
	cmd := helperCommandPath("stdout")
	p, err := StartProcess(cmd, lw)
	if err != nil {
		t.Fatal(err)
	}
	<-p.exited
	lines := ParseLines(buf)
	if lines[0].Message != "stdout" {
		t.Fatalf("got %v, want %v", lines[0].Message, "stdout")
	}
}

func TestProcessStderr(t *testing.T) {
	buf := new(bytes.Buffer)
	lw := &LineWriter{output: buf}
	cmd := helperCommandPath("stderr")
	p, err := StartProcess(cmd, lw)
	if err != nil {
		t.Fatal(err)
	}
	<-p.exited
	lines := ParseLines(buf)
	if lines[0].Message != "stderr" {
		t.Fatalf("got %v, want %v", lines[0].Message, "stderr")
	}
	if lines[0].Error != true {
		t.Fatalf("got %v, want %v", lines[0].Error, true)
	}
}

func TestProcessStartError(t *testing.T) {
	buf := new(bytes.Buffer)
	lw := &LineWriter{output: buf}
	cmd := exec.Command("/foobar")
	_, err := StartProcess(cmd, lw)
	if err == nil {
		t.Fatal("expected exec of /foobar to fail")
	}
}

func TestProcessFinished(t *testing.T) {
	buf := new(bytes.Buffer)
	lw := &LineWriter{output: buf}
	cmd := helperCommandPath("exit-nonzero")
	p, err := StartProcess(cmd, lw)
	if err != nil {
		t.Fatal(err)
	}
	<-p.exited
}

func TestSupervisorLoadDir(t *testing.T) {
	buf := new(bytes.Buffer)
	s, _ := NewSupervisor(buf)
	s.watcher = &mockWatcher{
		events: make(chan fsnotify.Event),
		errors: make(chan error),
	}
	s.restartDelay = time.Duration(0)
	s.cmdFactory = helperCommand
	s.maxRestarts = 0
	s.fs = afero.NewMemMapFs()
	afero.WriteFile(s.fs, "/bin/panic", []byte{}, 0700)
	afero.WriteFile(s.fs, "/bin/stderr", []byte{}, 0700)
	err := s.LoadDir("/bin")
	if err != nil {
		t.Fatal(err)
	}
	s.Wait()
}

func TestSupervisorPanic(t *testing.T) {
	buf := new(bytes.Buffer)
	s, _ := NewSupervisor(buf)
	s.watcher = &mockWatcher{
		events: make(chan fsnotify.Event),
		errors: make(chan error),
	}
	s.restartDelay = time.Duration(0)
	s.cmdFactory = helperCommand
	s.maxRestarts = 0
	s.fs = afero.NewMemMapFs()
	afero.WriteFile(s.fs, "panic", []byte{}, 0700)
	err := s.Reload("panic")
	if err != nil {
		t.Fatal(err)
	}
	s.Wait()
	b, e := afero.ReadFile(s.fs, "panic-debug.log")
	if e != nil {
		t.Fatal(e)
	}
	lines := ParseLines(buf)
	if !strings.HasPrefix(string(b), lines[0].Message) {
		t.Fatal("first line of debug log not equal to log output")
	}
}

func TestSupervisorStopProcess(t *testing.T) {
	buf := new(bytes.Buffer)
	s, _ := NewSupervisor(buf)
	s.watcher = &mockWatcher{
		events: make(chan fsnotify.Event),
		errors: make(chan error),
	}
	s.restartDelay = time.Duration(0)
	s.cmdFactory = helperCommand
	s.maxRestarts = 0
	s.fs = afero.NewMemMapFs()
	afero.WriteFile(s.fs, "/bin/wait", []byte{}, 0700)
	err := s.LoadDir("/bin")
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.Get("/bin/wait")
	if err != nil {
		t.Fatal("unexpected error getting process")
	}
	err = s.Stop("/bin/wait")
	if err != nil {
		t.Fatal("unexpected error stopping process")
	}
	s.Wait()
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]
	switch filepath.Base(cmd) {
	case "wait":
		select {}
	case "stdout":
		fmt.Println("stdout")
	case "stderr":
		fmt.Fprintln(os.Stderr, "stderr")
	case "exit-nonzero":
		os.Exit(1)
	case "panic":
		panic(nil)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}

}
