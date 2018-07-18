package main

import (
	"fmt"
	"time"

	"github.com/progrium/prototypes/electrond/client/electrond"
	"github.com/progrium/prototypes/qrpc"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	api := qrpc.NewAPI()
	om := qrpc.NewObjectManager()
	om.Mount(api, "objects")

	electron, err := electrond.Dial("localhost:4242", api)
	must(err)

	must(electron.GlobalShortcut.Register(electrond.GlobalShortcutRegisterParams{
		Accelerator: "CommandOrControl+Shift+T",
		Callback: om.Handle(func() {
			must(electron.Clipboard.WriteText(electrond.ClipboardWriteTextParams{
				Text: fmt.Sprint(time.Now()),
			}, nil))
		}),
	}, nil))

	electron.ServeAPI()
}
