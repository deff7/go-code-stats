package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	filename := "./test/example_go"
	mode := parser.Mode(0)
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filename, nil, mode)
	if err != nil {
		log.Fatal(err)
	}

	for _, imp := range f.Imports {
		if imp.Name != nil {
			fmt.Print(imp.Name.Name, " ")
		}
		fmt.Println(imp.Path.Value)
	}
}
