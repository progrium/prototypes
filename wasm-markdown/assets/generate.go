// +build ignore

package main

import (
	"log"

	"github.com/progrium/prototypes/wasm-markdown/assets"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		Filename:     "assets_generated.go",
		PackageName:  "assets",
		BuildTags:    "!gen",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
