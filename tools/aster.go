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
	rootPath string
	pathList []string
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
	flag.StringVar(&rootPath, "path", "", "Must specify /Path")
	flag.Parse()
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

func main() {
	// fmt.Println("DEBUG: ", rootPath)
	if err := filepath.Walk(rootPath, visit); err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return
	}

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
