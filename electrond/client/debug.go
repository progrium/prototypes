package main

import (
	"flag"

	"github.com/progrium/prototypes/qrpc"

	"github.com/progrium/prototypes/electrond/client/electrond"
)

const addr = "localhost:4242"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	api := qrpc.NewAPI()
	om := qrpc.NewObjectManager()
	om.Mount(api, "objects")

	electron, err := electrond.Dial(addr, api)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	err = electron.Shell.Beep(nil)
	if err != nil {
		panic(err)
	}

	var icon qrpc.ObjectHandle
	must(electron.NativeImage.CreateFromPath(electrond.NativeImageCreateFromPathParams{
		Path: "/Users/progrium/src/github.com/progrium/fluxbot/assets/Icon.png",
	}, &icon))

	var tray qrpc.ObjectHandle
	must(electron.Call("Tray.make", electrond.TrayParams{
		Image: icon,
	}, &tray))

	var menu qrpc.ObjectHandle
	must(electron.Menu.BuildFromTemplate(electrond.MenuBuildFromTemplateParams{
		Template: []*electrond.MenuItemConstructorOptions{
			&electrond.MenuItemConstructorOptions{
				Label: "Hello world",
			},
		},
	}, &menu))

	must(electron.Call(tray.ObjectPath+"/setContextMenu", map[string]interface{}{
		"menu": menu,
	}, nil))

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
