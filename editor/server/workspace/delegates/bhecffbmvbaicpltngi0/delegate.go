package node

import (
	"io"
	"net/http"

	"github.com/gliderlabs/ssh"
	"github.com/gliderlabs/stdcom/web/auth"
	"github.com/progrium/prototypes/editor/server/manifold"
)

func init() {
	manifold.RegisterDelegate(&Delegate{}, "bhecffbmvbaicpltngi0")
}

type Delegate struct {
	Message string
	Hello   string
	Auth    auth.Requestor `com:"singleton" json:"-"`
}

func (d *Delegate) HandleSSH(sess ssh.Session) {
	io.WriteString(sess, d.Message+"\n")
}

func (d *Delegate) MatchHTTP(r *http.Request) bool {
	return r.URL.Path == "/"
}

func (d *Delegate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// user := d.Auth.CurrentUser(r)
	// if user == nil {
	// 	io.WriteString(w, d.Message)
	// } else {
	// 	io.WriteString(w, fmt.Sprintf("%#v", user))
	// }
	io.WriteString(w, "<script>window.close()</script>")
}
