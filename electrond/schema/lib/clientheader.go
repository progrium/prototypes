package electrond

import (
	"github.com/progrium/prototypes/libmux/mux"
	"github.com/progrium/prototypes/qrpc"
)

func Dial(addr string, api qrpc.API) (*Client, error) {
	sess, err := mux.DialWebsocket(addr)
	if err != nil {
		return nil, err
	}
	return (&Client{
		Client: &qrpc.Client{
			Session: sess,
			API:     api,
		},
	}).setup(), nil
}
