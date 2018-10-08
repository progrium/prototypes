package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dop251/goja"
	"github.com/krpors/dom"
)

func Uint8Array(call goja.ConstructorCall) *goja.Object {
	// call.This contains the newly created object as per http://www.ecma-international.org/ecma-262/5.1/index.html#sec-13.2.2
	// call.Arguments contain arguments passed to the function

	//call.This.Set("method", method)

	// If return value is a non-nil *Object, it will be used instead of call.This
	// This way it is possible to return a Go struct or a map converted
	// into goja.Value using runtime.ToValue(), however in this case
	// instanceof will not work as expected.
	return nil
}

func main() {
	d, err := ioutil.ReadFile("./less.min.js")
	if err != nil {
		log.Fatal(err)
	}
	builder := dom.NewParser(bytes.NewBufferString("<html><head></head><body></body></html>"))
	doc, err := builder.Parse()
	if err != nil {
		panic(err)
	}
	vm := goja.New()
	var console = vm.NewObject()
	console.Set("log", fmt.Println)
	vm.Set("console", console)
	var window = vm.NewObject()
	window.Set("location", map[string]string{
		"protocol": "http",
	})
	var document = vm.NewObject()
	var currentScript = vm.NewObject()
	currentScript.Set("dataset", vm.NewObject())
	document.Set("currentScript", currentScript)
	document.Set("getElementsByTagName", doc.GetElementsByTagName)
	document.Set("createElement", doc.CreateElement)
	document.Set("createTextNode", doc.CreateText)
	document.Set("appendChild", doc.AppendChild)
	window.Set("document", document)
	vm.Set("window", window)
	vm.Set("document", document)
	vm.Set("Uint8Array", Uint8Array)
	_, err = vm.RunString(string(d))
	//_, err = vm.RunString(`console.log(window.document.getElementsByTagName("body")[0])`)
	if err != nil {
		panic(err)
	}
}
