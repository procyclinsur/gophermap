package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
)

var (
	//fileVar string
	rootPath string
	pathList []string
	debug    bool
)

//StructDef v1.0
type StructDef struct {
	//Name of struct
	Name string
	//Map of property-name:property-type
	Properties map[string]string
	//List of structs contained
	Contains []string
}

func init() {
	//flag.StringVar(&fileVar, "file", "", "Must specify /Path/file")
	flag.StringVar(&rootPath, "path", "", "Must specify /Path")
	flag.BoolVar(&debug, "debug", false, "For AST output set to true.")
	flag.Parse()
}

func main() {
	if err := filepath.Walk(rootPath, visit); err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return
	}

	if debug != true {
		parseDirs()
	} else {
		debugParseDirs()
	}
}

func parseDirs() {
	fset := token.NewFileSet()

	for _, pathVar := range pathList {
		// fmt.Println("DEBUG: ", pathVar)
		prse, err := parser.ParseDir(fset, pathVar, nil, 0)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		for _, pkgItem := range prse {
			ast.Walk(VisitorFunc(FindTypes), pkgItem)
		}
	}
}

func debugParseDirs() {
	fset := token.NewFileSet()

	for _, pathVar := range pathList {
		// fmt.Println("DEBUG: ", pathVar)
		prse, err := parser.ParseDir(fset, pathVar, nil, 0)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		for _, pkgItem := range prse {
			ast.Fprint(os.Stdout, fset, pkgItem, func(name string, value reflect.Value) bool {
				if ast.NotNilFilter(name, value) {
					return value.Type().String() != "*ast.Object"
				}
				return false
			})
		}
	}
}

func visit(path string, f os.FileInfo, err error) error {
	r, _ := regexp.Compile(`(^.*/\.[^\.].*$)`)
	if f.IsDir() {
		match := r.MatchString(path)
		if match != true {
			pathList = append(pathList, path)
			return nil
		}
	}
	return nil
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
