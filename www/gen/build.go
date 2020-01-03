package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var App http.FileSystem = http.Dir("build")

func main() {
	if err := vfsgen.Generate(App, vfsgen.Options{
		PackageName: "www",
		// BuildTags:    "",
		VariableName: "App",
		Filename:     "assets.gen.go",
	}); err != nil {
		log.Fatalln(err)
	}
}
