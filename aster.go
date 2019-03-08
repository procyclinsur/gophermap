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
		return walkStructSpec(n)
	}
	return nil
}

func fieldNameNilString(n []*ast.Ident) (fName string) {
	if n != nil {
		fName = n[0].Name
	} else {
		fName = "nil"
	}
	return fName
}

func walkStructSpec(n *ast.TypeSpec) ast.Visitor {
	switch v := n.Type.(type) {
	case *ast.StructType:
		npos := fset.Position(n.Name.Pos())
		fmt.Println("----")
		fmt.Printf("line/%d ", npos.Line)
		fmt.Printf("struct:%s\n", n.Name.Name)
		for index, item := range v.Fields.List {
			vpos := fset.Position(v.Fields.Pos())
			fName := fieldNameNilString(item.Names)
			var fType string
			var xSel string
			var xxName string
			switch s := item.Type.(type) {
			case *ast.StarExpr:
				switch p := s.X.(type) {
				case *ast.SelectorExpr:
					xSel = p.Sel.Name
					switch x := p.X.(type) {
					case *ast.Ident:
						xxName = x.Name
					}
				}
				fType = "*" + xxName + "." + xSel
			case *ast.Ident:
				fType = s.Name
			}
			fmt.Printf("line:%d ", vpos.Line)
			fmt.Printf("char:%d ", vpos.Column)
			fmt.Printf("index:%d ", index)
			fmt.Printf("parm:%s ", fName)
			fmt.Printf("type:%s\n", fType)
		}
		return VisitorFunc(FindTypes)
	}
	return nil
}
