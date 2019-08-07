package daemon

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
)

var logBus = NewMulticastWriteCloser()
var currentServer *exec.Cmd

const addr = "localhost:4242"

func runServer() error {
	bin, err := exec.LookPath("go")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(bin, "run", "./workspace/server.go")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = logBus
	cmd.Stderr = logBus
	currentServer = cmd
	return cmd.Run()
}

func watchDelegates(watcher *fsnotify.Watcher) {
	files, err := ioutil.ReadDir("./workspace/delegates")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			if err := watcher.Add(path.Join("./workspace/delegates", f.Name(), "delegate.go")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func Run() {
	log.SetOutput(logBus)
	go logBus.WriteTo(os.Stdout)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if strings.HasSuffix(event.Name, "delegates/delegates.go") {
						watchDelegates(watcher)
					}
					if currentServer != nil {
						syscall.Kill(-currentServer.Process.Pid, syscall.SIGTERM)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()
	if err := watcher.Add("./workspace"); err != nil {
		log.Fatal(err)
	}
	if err := watcher.Add("./workspace/delegates/delegates.go"); err != nil {
		log.Fatal(err)
	}
	watchDelegates(watcher)
	// TODO: watch all delegate directories

	go func() {
		for {
			err := runServer()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			if currentServer != nil {
				syscall.Kill(-currentServer.Process.Pid, syscall.SIGTERM)
			}
			os.Exit(0)
		}
	}()

	log.Println("running daemon...")
	log.Fatal(ListenAndServe(addr))
}

func ListenAndServe(addr string) error {
	server := &qrpc.Server{}
	l, err := mux.ListenWebsocket(addr)
	if err != nil {
		panic(err)
	}
	api := qrpc.NewAPI()
	api.HandleFunc("console", func(r qrpc.Responder, c *qrpc.Call) {
		ch, err := r.Hijack(nil)
		if err != nil {
			log.Println(err)
		}
		logBus.WriteTo(ch)
	})
	return server.Serve(l, api)
}
