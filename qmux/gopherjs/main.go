package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/progrium/prototypes/qmux"
)

// Dumb experiment. Ignore

func main() {
	js.Global.Set("qmux", map[string]interface{}{
		"NewSession": qmux.NewSession,
	})
}
