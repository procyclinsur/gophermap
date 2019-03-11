package main

import (
	"go/ast"
	"go/token"
)

var structMap = StructMap{}

//StructMap v1.0
type StructMap map[string]StructDef

//StructDef v1.0
type StructDef struct {
	//Name of struct
	Name string
	//Map of property-name:property-type
	Properties map[string]string
	//List of structs contained
	Contains []string
}

//VisitorFunc function v1.0
type VisitorFunc func(n ast.Node) ast.Visitor

//Visit function v1.0
func (f VisitorFunc) Visit(n ast.Node) ast.Visitor {
	return f(n)
}

func getStructMap() StructMap {
	return structMap
}

//FindTypes function v1.0
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
			switch s := item.Type.(type) {
			case *ast.StarExpr:
				var xSel string
				var xxName string
				switch p := s.X.(type) {
				case *ast.SelectorExpr:
					xSel = p.Sel.Name
					switch x := p.X.(type) {
					case *ast.Ident:
						xxName = x.Name
					}
				}
				fieldType = "*" + xxName + "." + xSel
			case *ast.MapType:
				var mKey string
				var mVal string
				switch k := s.Key.(type) {
				case *ast.Ident:
					mKey = k.Name
				}
				switch v := s.Value.(type) {
				case *ast.Ident:
					mVal = v.Name
				}
				fieldType = "map[" + mKey + "]" + mVal
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
