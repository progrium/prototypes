package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/rjeczalik/notify"
	"golang.org/x/net/websocket"
)

var (
	listen       = flag.String("listen", ":8080", "listen address")
	dir          = flag.String("dir", ".", "directory to serve")
	buildErrored = false
)

var reloadCh = make(map[chan bool]bool)
var reloadMu sync.Mutex

func ReloadServer(ws *websocket.Conn) {
	ch := make(chan bool)
	reloadMu.Lock()
	reloadCh[ch] = true
	reloadMu.Unlock()
	<-ch
	ws.Write([]byte("Reload"))
	reloadMu.Lock()
	delete(reloadCh, ch)
	reloadMu.Unlock()
}

func main() {
	flag.Parse()

	go func() {
		notifyChanges("..", []string{".go"}, false, func(path string) {
			buildAndRun()
		})
	}()

	log.Printf("listening on %q...", *listen)
	http.Handle("/reload", websocket.Handler(ReloadServer))
	http.HandleFunc("/", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		http.FileServer(http.Dir(*dir)).ServeHTTP(resp, req)
	}))
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func extensionIn(path string, exts []string) bool {
	for _, ext := range exts {
		if filepath.Ext(path) == ext {
			return true
		}
	}
	return false
}

func notifyChanges(dir string, exts []string, onlyCreate bool, cb func(path string)) {
	c := make(chan notify.EventInfo, 1)
	types := notify.All
	if onlyCreate {
		types = notify.Create
	}
	if err := notify.Watch(dir, c, types); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)
	for event := range c {
		path := event.Path()
		dir, file := filepath.Split(path)
		if filepath.Base(dir) == ".git" {
			continue
		}
		if filepath.Base(file)[0] == '.' {
			continue
		}
		if extensionIn(path, exts) {
			cb(path)
		}
	}
}

func buildAndRun() error {
	start := time.Now()
	cmd := exec.Command("sh", "-c", "GOARCH=wasm GOOS=js go build -o test.wasm ../main.go")
	output, err := cmd.CombinedOutput()
	if !cmd.ProcessState.Success() {
		buildErrored = true
		log.Println("ERROR! Build failed:")
		fmt.Println(string(output))
	} else {
		log.Println("New build:", time.Now().Sub(start))
		buildErrored = false
		for ch, _ := range reloadCh {
			close(ch)
		}
	}
	time.Sleep(100 * time.Millisecond)
	return err
}
