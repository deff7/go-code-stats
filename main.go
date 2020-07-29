package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func walkScope(scope *types.Scope) {
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		if obj == nil {
			log.Printf("no such name: %v", name)
			continue
		}

		switch v := obj.(type) {
		case *types.Func:
			walkScope(v.Scope())
		case *types.Var:
			log.Println(v.Name(), " -- ", v.Type().String())
		default:
			log.Printf("unknown %#v", v)
		}
	}
}

func main() {
	filename := "./test/simple_go"
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

	conf := types.Config{Importer: importer.Default()}

	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	fmt.Printf("Package  %q\n", pkg.Path())
	fmt.Printf("Name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())
	walkScope(pkg.Scope())
}
