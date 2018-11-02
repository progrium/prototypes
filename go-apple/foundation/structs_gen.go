// +build ignore

package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/progrium/prototypes/go-apple/bridgesupport"
)

func main() {
	sigs, err := bridgesupport.LoadSignatures("./Foundation.bridgesupport")
	if err != nil {
		panic(err)
	}
	f := jen.NewFile("foundation")
	for _, s := range sigs.Structs {
		var fields []jen.Code
		for i, f := range s.Type.Names {
			t := s.Type.Types[i].String()
			if s.Type.Types[i].Kind() == reflect.Struct {
				t = strings.Trim(s.Type.Types[i].Name, "_")
			}
			fields = append(fields, jen.Id(strings.Title(f)).Id(t))
		}
		f.Type().Id(s.Name).Struct(fields...)
	}
	fmt.Printf("%#v", f)
}
