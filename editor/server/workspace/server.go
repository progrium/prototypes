package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gliderlabs/com/objects"
	"github.com/gliderlabs/ssh"
	"github.com/gliderlabs/stdcom/daemon"
	logapi "github.com/gliderlabs/stdcom/log"
	"github.com/gliderlabs/stdcom/log/std"
	"github.com/gliderlabs/stdcom/web"
	"github.com/gliderlabs/stdcom/web/auth"
	"github.com/gliderlabs/stdcom/web/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/progrium/prototypes/editor/server/manifold"
	"github.com/progrium/prototypes/editor/server/manifold/frontend"
	"github.com/progrium/prototypes/editor/server/manifold/workspace"

	_ "github.com/progrium/prototypes/editor/server/workspace/delegates"
)

func init() {
	manifold.RegisterComponent(&PlaceholderComponent{})
	manifold.RegisterComponent(&AnotherComponent{})
	manifold.RegisterComponent(&SSHServer{})
	manifold.RegisterComponent(&WebServer{})
	manifold.RegisterComponent(&Auth{})
}

const addr = "localhost:4243"

type Auth struct {
	Logger  logapi.DebugLogger `com:"singleton" json:"-"`
	Session sessions.Session   `com:"singleton" json:"-"`
	User    auth.User

	auth     *auth.Component
	registry *objects.Registry
	node     *manifold.Node
}

func (c *Auth) InspectorButtons() []frontend.Button {
	return []frontend.Button{
		{
			Name:    "Login",
			OnClick: "window.open('http://localhost:8080/_auth/login/github?return=/', '_blank')",
		},
		{
			Name:    "Logout",
			OnClick: "window.open('http://localhost:8080/_auth/logout?return=/', '_blank')",
		},
	}
}

func (c *Auth) InitializeComponent(n *manifold.Node) {
	c.registry = n.Registry
	c.node = n
}

func (c *Auth) InitializeDaemon() error {
	c.auth = &auth.Component{
		Config: auth.Config{
			BasePath: "/_auth",
		},
		Log:     c.Logger,
		Session: c.Session,
	}
	c.registry.Register(&objects.Object{Value: c.auth})
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_CLIENT"), os.Getenv("GITHUB_SECRET"), "http://localhost:8080/_auth/callback/github"),
	)
	return nil
}

func (c *Auth) MatchHTTP(r *http.Request) bool {
	return c.auth.MatchHTTP(r)
}

func (c *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.auth.ServeHTTP(w, r)
	u := c.auth.CurrentUser(r)
	if u != nil {
		c.User = *u
	} else {
		c.User = auth.User{}
	}
	c.node.Sync()
}

type WebServer struct {
	Addr     string
	Handlers []web.Handler `com:"extpoint" json:"-"`
	Logger   logapi.Logger `com:"singleton" json:"-"`

	server *web.Component
}

func (c *WebServer) InitializeDaemon() error {
	log.Println("setting up web")
	c.server = &web.Component{}
	c.server.Log = c.Logger
	c.server.ListenAddr = c.Addr
	c.server.Handlers = c.Handlers
	return nil
}

func (c *WebServer) Serve() {
	log.Printf("starting web on %s", c.Addr)
	c.server.Serve()
}

func (c *WebServer) Stop() {
	c.server.Stop()
}

type SSHServer struct {
	Addr    string
	Handler SSHHandler `com:"singleton" json:"-"`

	server *ssh.Server
}

type SSHHandler interface {
	HandleSSH(sess ssh.Session)
}

func (c *SSHServer) Hello() {
	log.Println("Hello!")
}

func (c *SSHServer) InspectorButtons() []frontend.Button {
	return []frontend.Button{{
		Name: "Hello",
	}}
}

func (c *SSHServer) InitializeDaemon() error {
	log.Println("setting up ssh")
	c.server = &ssh.Server{}
	c.server.Addr = c.Addr
	c.server.Handler = func(s ssh.Session) {
		c.Handler.HandleSSH(s)
	}
	c.server.SetOption(ssh.HostKeyFile("/Users/progrium/.ssh/id_rsa"))
	return nil
}

func (c *SSHServer) Serve() {
	log.Printf("starting ssh on %s", c.Addr)
	log.Fatal(c.server.ListenAndServe())
}

func (c *SSHServer) Stop() {
	c.server.Shutdown(context.Background())
}

type MoreData struct {
	Foobar string
	Baz    bool
}

type PlaceholderComponent struct {
	StringValue string
	IntValue    int
	Object      MoreData
	MapData     map[string]string
	ListData    []string
}

func (c *PlaceholderComponent) Foobar() {
	log.Println("Hello again")
}

func (c *PlaceholderComponent) DifferentButton() {
	log.Println("Hello different")
}

type AnotherComponent struct {
	Foobar  string
	RefTest *PlaceholderComponent
}

func (c *AnotherComponent) PrintRefStringValue() {
	if c.RefTest != nil {
		log.Println(c.RefTest.StringValue)
	}
}

var root = manifold.NewNode("")

func main() {
	var err error
	root, err = workspace.LoadHierarchy()
	if err != nil {
		panic(err)
	}

	registry := &objects.Registry{}

	manifold.Walk(root, func(n *manifold.Node) {
		for _, com := range n.Components {
			registry.Register(objects.New(com.Ref, ""))
		}
	})
	std.Register(registry)
	sessions.Register(registry)
	registry.Reload()

	root.Observe(&manifold.NodeObserver{
		OnChange: func(node *manifold.Node, path string, old, new interface{}) {
			if path == "Name" && node.Dir != "" {
				newDir := filepath.Join(filepath.Dir(node.Dir), new.(string))
				if node.Dir != newDir {
					// TODO: do not break abstraction, have workspace handle this
					if err := os.Rename(node.Dir, newDir); err != nil {
						log.Fatal(err)
					}
				}
			}
			err := workspace.SaveHierarchy(root)
			if err != nil {
				panic(err)
				//log.Println(err)
			}
		},
	})

	go func() {
		daemon.Run(registry, "app")
	}()

	frontend.ListenAndServe(root, addr)
}
