package main

import (
	"flag"
	"os"

	"github.com/progrium/prototypes/supervisor"
)

func main() {
	flag.Parse()
	s, err := supervisor.NewSupervisor(os.Stdout)
	if err != nil {
		panic(err)
	}
	go s.Watch()
	err = s.LoadDir(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	s.Wait()
}
