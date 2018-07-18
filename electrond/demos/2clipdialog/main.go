package main

import (
	"fmt"

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

	var clipboard string
	must(electron.Clipboard.ReadText(electrond.ClipboardReadTextParams{}, &clipboard))

	fmt.Println(clipboard)

	must(electron.Dialog.ShowMessageBox(electrond.DialogShowMessageBoxParams{
		Options: electrond.DialogShowMessageBoxParamsOptions{
			Message: clipboard,
		},
	}, nil))

}
