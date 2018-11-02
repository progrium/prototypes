// +build ignore

package main

import (
	"fmt"
	"strconv"
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
	for _, s := range sigs.Enums {
		v := s.Value64
		if v == "" {
			v = s.Value
		}
		t := "int64"
		var val interface{}
		var err error
		if strings.Contains(v, ".") {
			t = "float64"
			if val, err = strconv.ParseFloat(v, 64); err != nil {
				fmt.Println(s.Name)
				panic(err)
			}
		} else {
			if val, err = strconv.ParseInt(v, 10, 64); err != nil {
				if val, err = strconv.ParseUint(v, 10, 64); err != nil {
					fmt.Println(s.Name)
					panic(err)
				}
			}
		}
		if s.Name == "NSTimeIntervalSince1970" {
			// otherwise Lit will turn it into 9.783072e+08.0
			f.Const().Id(s.Name).Id(t).Op("=").Id(v)
		} else {
			f.Const().Id(s.Name).Id(t).Op("=").Lit(val)
		}

	}
	fmt.Printf("%#v", f)
}
