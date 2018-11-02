package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "test.go", nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("parser.ParseFile error(%v)\n", err)
		return
	}

	for _, importSpec := range f.Imports {
		var importName string

		if importSpec.Name == nil {
			importName = ""
		} else {
			importName = importSpec.Name.Name
		}

		importPath := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
		pkg, err := build.Import(importPath, "", 0)

		if err != nil {
			fmt.Printf("build.Import error(%v)\n", err)
			return
		}

		if importName == "" {
			importName = pkg.Name
		}

		fmt.Printf("importName(%+v)\n", importName)

	}

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)

		if !ok {
			continue
		}

		if genDecl.Tok != token.TYPE {
			continue
		}

		typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)

		if !ok {
			continue
		}

		fmt.Printf("typeSpec(%+v) \n", interfaceType)

		for _, method := range interfaceType.Methods.List {
			fmt.Printf("method(%+v)\n", method.Names)
		}
	}
}
