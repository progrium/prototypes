package main

import (
	"fmt"
	"log"

	"github.com/progrium/prototypes/go-obs-websocket"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "ws://192.168.1.200:4444/"
	client, err := obs.DialWebsocket(url, "12345")
	fatal(err)
	defer client.Close()

	resp, err := client.Request("GetSceneList", nil)
	fatal(err)

	fmt.Println(resp)

	for event := range client.Events {
		fmt.Println(event)
	}
}
