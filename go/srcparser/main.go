package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	ast, err := parser.ParseDir(
		token.NewFileSet(),
		os.ExpandEnv("$GOPATH")+"/src/github.com/sadlil/experiments/go/srcparser/src",
		nil,
		parser.ParseComments,
	)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range ast {
		fmt.Println(k)
		d := doc.New(v, "", 0)
		fmt.Println(d.Doc)

		for _, t := range d.Types {
			fmt.Println(t.Name)
			fmt.Println(t.Doc)
		}

		for _, f := range v.Files {
			fmt.Println(f.Comments)
			fmt.Println(f.Doc.Text())
		}

	}
}
