package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	n := 0
	for {
		fmt.Println(os.Getpid(), n)
		n += 2
		time.Sleep(2 * time.Second)
	}
}
