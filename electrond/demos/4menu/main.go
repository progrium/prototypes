package main

import (
	"github.com/progrium/prototypes/qrpc"

	"github.com/progrium/prototypes/electrond/client/electrond"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func objPtr(h qrpc.ObjectHandle) *qrpc.ObjectHandle {
	return &h
}

func main() {
	api := qrpc.NewAPI()
	om := qrpc.NewObjectManager()
	om.Mount(api, "objects")

	electron, err := electrond.Dial("localhost:4242", api)
	must(err)

	var icon qrpc.ObjectHandle
	must(electron.NativeImage.CreateFromPath(electrond.NativeImageCreateFromPathParams{
		Path: "/Users/progrium/src/github.com/progrium/fluxbot/assets/Icon.png",
	}, &icon))

	var tray qrpc.ObjectHandle
	must(electron.Call("Tray.make", electrond.TrayParams{
		Image: icon,
	}, &tray))

	var notify qrpc.ObjectHandle
	must(electron.Call("Notification.make", electrond.NotificationParams{
		Options: electrond.NotificationParamsOptions{
			Title: "Golang ATX",
			Body:  "Hello world",
		},
	}, &notify))

	var menu qrpc.ObjectHandle
	must(electron.Menu.BuildFromTemplate(electrond.MenuBuildFromTemplateParams{
		Template: []*electrond.MenuItemConstructorOptions{
			&electrond.MenuItemConstructorOptions{
				Label: "Hello world",
				Click: objPtr(om.Handle(func() {
					electron.Call(notify.ObjectPath+"/show", nil, nil)
				})),
			},
		},
	}, &menu))

	must(electron.Call(tray.ObjectPath+"/setContextMenu", map[string]interface{}{
		"menu": menu,
	}, nil))

	electron.ServeAPI()
}
