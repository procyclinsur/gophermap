package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
)

var (
	//fileVar string
	rootPath string
	pathList []string
	astDebug bool
	fset     *token.FileSet
	//debug bool
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
	flag.BoolVar(&astDebug, "astdebug", false, "For AST output set to true.")
	flag.Parse()
}

func main() {
	if err := getPathList(rootPath, visit); err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	fset = token.NewFileSet()

	if astDebug != true {
		parseDirFiles(fset)
	} else {
		debugParseDirFiles(fset)
	}
}

func parseDirFiles(f *token.FileSet) {
	for _, pathVar := range pathList {
		//fmt.Println("DEBUG: ", pathVar)
		prse, err := parser.ParseDir(f, pathVar, fileFilter, 0)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		for _, pkgItem := range prse {
			ast.Walk(VisitorFunc(FindTypes), pkgItem)
		}
	}
}

func debugParseDirFiles(f *token.FileSet) {
	for _, pathVar := range pathList {
		// fmt.Println("DEBUG: ", pathVar)
		prse, err := parser.ParseDir(f, pathVar, fileFilter, 0)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		for _, pkgItem := range prse {
			ast.Fprint(os.Stdout, f, pkgItem, func(name string, value reflect.Value) bool {
				if ast.NotNilFilter(name, value) {
					return value.Type().String() != "*ast.Object"
				}
				return false
			})
		}
	}
}
