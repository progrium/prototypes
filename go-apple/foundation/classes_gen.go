// +build ignore

package main

import (
	"fmt"
	"reflect"
	"regexp"
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

var selectorRegex = regexp.MustCompile(`:[[:lower:]]`)

func selectorToFuncName(s string) string {
	sel := selectorRegex.ReplaceAllStringFunc(s, func(ss string) string {
		return strings.Title(ss[1:])
	})
	return strings.TrimRight(strings.Title(sel), ":")
}

func main() {
	sigs, err := bridgesupport.LoadSignatures("./Foundation.bridgesupport")
	if err != nil {
		panic(err)
	}
	f := jen.NewFile("foundation")
	for _, c := range sigs.Classes {
		f.Type().Id(c.Name).Struct(
			jen.Qual("github.com/progrium/objc", "Object"),
		)
		for _, m := range c.Methods {
			if !m.ClassMethod {
				continue
			}
			var args []jen.Code
			for i, arg := range m.Args {
				if arg.Type.Code == "" && arg.Type64 == nil {
					// TODO: handle
					goto skip
				}
				rt := arg.Type
				if arg.Type64 != nil {
					rt = *arg.Type64
				}
				varName := string(byte(97 + i))
				args = append(args, jen.Id(varName).Id(toID(rt)))
			}
			if m.RetVal != nil {
				rt := m.RetVal.Type
				if m.RetVal.Type64 != nil {
					rt = *m.RetVal.Type64
				}
				f.Func().Id(selectorToFuncName(m.Selector)).Params(args...).Id(toID(rt)).Block()
			} else {
				f.Func().Id(selectorToFuncName(m.Selector)).Params(args...).Block()
			}
		skip:
		}
	}
	fmt.Printf("%#v", f)
}
