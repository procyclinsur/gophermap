package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

var (
	//fileVar string
	pathVar string
)

func init() {
	//flag.StringVar(&fileVar, "file", "", "Must specify /Path/file")
	flag.StringVar(&pathVar, "path", "", "Must specify /Path")
	flag.Parse()
}

func main() {
	fset := token.NewFileSet()
	//file, err := ioutil.ReadFile(fileVar)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//src := string(file)

	prse, err := parser.ParseDir(fset, pathVar, nil, 0)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	for _, pkgItem := range prse {
		ast.Walk(VisitorFunc(FindTypes), pkgItem)
	}
}

type VisitorFunc func(n ast.Node) ast.Visitor

func (f VisitorFunc) Visit(n ast.Node) ast.Visitor {
	return f(n)
}

func FindTypes(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.Package:
		return VisitorFunc(FindTypes)
	case *ast.File:
		return VisitorFunc(FindTypes)
	case *ast.GenDecl:
		if n.Tok == token.TYPE {
			return VisitorFunc(FindTypes)
		}
	case *ast.TypeSpec:
		switch v := n.Type.(type) {
		case *ast.StructType:
			fmt.Printf("----\n%s\n", n.Name.Name)
			for _, item := range v.Fields.List {
				fmt.Println(item.Names[0].Name, item.Type)
			}
			return VisitorFunc(FindTypes)
		}
	}
	return nil
}
