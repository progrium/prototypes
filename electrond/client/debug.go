package main

import (
	"flag"
	"fmt"

	"github.com/progrium/prototypes/qrpc"

	"github.com/progrium/prototypes/electrond/client/electrond"
)

const addr = "localhost:4242"

func main() {
	api := qrpc.NewAPI()
	om := qrpc.NewObjectManager()
	om.Mount(api, "objects")

	cb := om.Register(func() {
		fmt.Println("CALLBACK")
	})

	electron, err := electrond.Dial(addr, api)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	err = electron.Shell.Beep(nil)
	if err != nil {
		panic(err)
	}

	err = electron.GlobalShortcut.Register(electrond.GlobalShortcutRegisterParams{Accelerator: "CommandOrControl+B", Callback: cb.Handle()}, nil)
	if err != nil {
		panic(err)
	}

	electron.ServeAPI()

	// err = electron.App.SetBadgeCount(electrond.AppSetBadgeCountParams{2}, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// err = electron.AppDock.Bounce(electrond.DockBounceParams{"critical"}, nil)
	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Printf("resp: %#v\n", resp2)
}
