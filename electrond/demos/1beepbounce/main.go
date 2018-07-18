package main

import (
	"github.com/progrium/prototypes/electrond/client/electrond"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	electron, err := electrond.Dial("localhost:4242", nil)
	must(err)

	must(electron.Shell.Beep(nil))

	must(electron.AppDock.Bounce(electrond.DockBounceParams{Type: "critical"}, nil))

}
