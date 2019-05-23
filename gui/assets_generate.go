// +build generate

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var Assets http.FileSystem = http.Dir("dist")
	err := vfsgen.Generate(Assets, vfsgen.Options{
		PackageName:  "gui",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
