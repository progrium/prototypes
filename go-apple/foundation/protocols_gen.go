// +build ignore

package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/progrium/prototypes/go-apple/bridgesupport"
)

func main() {
	sigs, err := bridgesupport.LoadSignatures("./Foundation.bridgesupport")
	if err != nil {
		panic(err)
	}
	f := jen.NewFile("foundation")
	for _, p := range sigs.InformalProtocols {
		f.Type().Id(p.Name).Interface()
	}
	fmt.Printf("%#v", f)
}
