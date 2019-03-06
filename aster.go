package main

import (
  "flag"
  "go/ast"
  "go/parser"
  "go/token"
  "os"
  "reflect"
  "log"
)

var(
  pathVar string
)

func init() {
	flag.StringVar(&pathVar, "path", "", "Must specify /Path")
	flag.Parse()
}

func main() {
	fset := token.NewFileSet()

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
