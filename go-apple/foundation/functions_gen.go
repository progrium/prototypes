// +build ignore

package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/progrium/prototypes/go-apple/bridgesupport"
)

func toID(ti bridgesupport.TypeInfo) string {
	switch ti.Kind() {
	case reflect.Struct:
		return strings.Trim(ti.Name, "_")
	default:
		return ti.String()
	}
}

func main() {
	sigs, err := bridgesupport.LoadSignatures("./Foundation.bridgesupport")
	if err != nil {
		panic(err)
	}
	f := jen.NewFile("foundation")
	for _, fn := range sigs.Functions {
		if fn.Inline {
			continue
		}
		var args []jen.Code
		for i, arg := range fn.Args {
			rt := arg.Type
			if arg.Type64 != nil {
				rt = *arg.Type64
			}
			varName := string(byte(97 + i))
			args = append(args, jen.Id(varName).Id(toID(rt)))
		}
		if fn.RetVal != nil {
			rt := fn.RetVal.Type
			if fn.RetVal.Type64 != nil {
				rt = *fn.RetVal.Type64
			}
			f.Func().Id(fn.Name).Params(args...).Id(toID(rt)).Block()
		} else {
			f.Func().Id(fn.Name).Params(args...).Block()
		}

	}
	fmt.Printf("%#v", f)
}
