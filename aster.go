package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

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
			npos := fset.Position(n.Name.Pos())
			fmt.Println("----")
			fmt.Printf("line/%d ", npos.Line)
			fmt.Printf("struct:%s\n", n.Name.Name)
			for index, item := range v.Fields.List {
				var fName string
				if item.Names != nil {
					fName = item.Names[0].Name
				} else {
					fName = "nil"
				}
				vpos := fset.Position(v.Fields.Pos())
				fmt.Printf("line:%d ", vpos.Line)
				fmt.Printf("char:%d ", vpos.Column)
				fmt.Printf("index:%d ", index)
				fmt.Printf("parm:%s ", fName)
				fmt.Printf("type:%s\n", item.Type)
			}
			return VisitorFunc(FindTypes)
		}
	}
	return nil
}
