package main

import (
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

	return fName
}

func walkStructSpec(n *ast.TypeSpec) ast.Visitor {
	switch v := n.Type.(type) {
	case *ast.StructType:
		structMap[n.Name.Name] = StructDef{
			n.Name.Name,
			map[string]string{},
			[]string{},
		}

		for _, item := range v.Fields.List {
			var fieldType string
			var fieldName string
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
				fieldType = "*" + xxName + "." + xSel
			case *ast.Ident:
				fieldType = s.Name
			}
			if item.Names != nil {
				fieldName = item.Names[0].Name
			} else {
				fieldName = fieldType
			}
			structMap[n.Name.Name].Properties[fieldName] = fieldType
		}
		return VisitorFunc(FindTypes)
	}
	return nil
}
