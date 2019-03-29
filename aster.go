package main

import (
	"go/ast"
	"go/token"
	"strings"
)

var structMap = StructMap{}
var typeList = TypeList{}

//TypeList v1.0
type TypeList []string

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

func getWalkOutput() (tl TypeList, sm StructMap) {
	tl = typeList
	sm = structMap
	return
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
		return walkTypeSpec(n)
	}
	return nil
}

func walkTypeSpec(n *ast.TypeSpec) ast.Visitor {
	typeList = append(typeList, n.Name.Name)
	switch v := n.Type.(type) {
	case *ast.StructType:
		sugar.Debugf("Struct: %s", n.Name.Name)
		return walkStructSpec(n, v)
	}
	return nil
}

func walkStructSpec(n *ast.TypeSpec, v *ast.StructType) ast.Visitor {
	recordAstStructType(n.Name.Name, v)
	return VisitorFunc(FindTypes)
}

func getUndeterminedType(fn string, fi *ast.Field) (rv string) {
	switch s := fi.Type.(type) {
	case *ast.StructType:
		rv = recordAstStructType(fn, s)
	case *ast.StarExpr:
		rv = getAstStarExpr(s)
	case *ast.MapType:
		rv = getAstMapType(s)
	case *ast.ArrayType:
		rv = getAstArrayType(s)
	case *ast.SelectorExpr:
		rv = getAstSelectorExpr(s)
	case *ast.FuncType:
		rv = getAstFuncType(s)
	case *ast.ChanType:
		rv = getAstChanType(s)
	case *ast.Ident:
		rv = getAstIdent(s)
	}
	return
}

func recordAstStructType(fn string, s *ast.StructType) (rv string) {
	structMap[fn] = StructDef{
		fn,
		map[string]string{},
		[]string{},
	}
	for _, item := range s.Fields.List {
		var fieldName string

		if item.Names != nil {
			fieldName = item.Names[0].Name
		}

		sugar.Debugf("    Field: %s", fieldName)

		fieldType := getUndeterminedType(fieldName, item)

		sugar.Debugf("        Type: %s", fieldType)

		if fieldName == "" {
			fieldName = fieldType
		}

		structMap[fn].Properties[fieldName] = fieldType
	}
	return "struct"
}

func getAstChanType(s *ast.ChanType) (rv string) {
	var tv string

	switch se := s.Value.(type) {
	case *ast.SelectorExpr:
		tv = getAstSelectorExpr(se)
	}

	if s.Dir == ast.SEND {
		rv = "chan" + "<-" + tv
	} else if s.Dir == ast.RECV {
		panic("Check AST Document Dir was not 1")
	} else {
		logger.Fatal("Unknown Error Invalid Chan Type")
	}

	return
}

func getAstFuncType(s *ast.FuncType) (rv string) {
	var tfn []string
	var trn []string

	// get parameters
	for _, p := range s.Params.List {
		tfn = append(tfn, getUndeterminedType("func", p))
	}

	// get return values
	for _, r := range s.Results.List {
		trn = append(trn, getUndeterminedType("func", r))
	}

	fn := strings.Join(tfn, " ")
	rn := strings.Join(trn, " ")

	rv = "func(" + fn + ")" + "(" + rn + ")"

	return
}

func getAstStarExpr(s *ast.StarExpr) (rv string) {
	switch se := s.X.(type) {
	case *ast.SelectorExpr:
		rv = "*" + getAstSelectorExpr(se)
	case *ast.Ident:
		rv = "*" + getAstIdent(se)
	}
	return
}

func getAstMapType(s *ast.MapType) (rv string) {
	var mKey string
	switch mtk := s.Key.(type) {
	case *ast.Ident:
		mKey = getAstIdent(mtk)
	}
	switch mtv := s.Value.(type) {
	case *ast.Ident:
		rv = "map[" + mKey + "]" + getAstIdent(mtv)
	case *ast.ArrayType:
		rv = "map[" + mKey + "]" + getAstArrayType(mtv)
	case *ast.SelectorExpr:
		rv = "map[" + mKey + "]" + getAstSelectorExpr(mtv)
	}
	return
}

func getAstArrayType(s *ast.ArrayType) (rv string) {
	switch at := s.Elt.(type) {
	case *ast.StarExpr:
		rv = "[]" + getAstStarExpr(at)
	case *ast.MapType:
		rv = "[]" + getAstMapType(at)
	case *ast.ArrayType:
		rv = "[]" + getAstArrayType(at)
	case *ast.SelectorExpr:
		rv = "[]" + getAstSelectorExpr(at)
	case *ast.FuncType:
		rv = "[]" + getAstFuncType(at)
	case *ast.Ident:
		rv = "[]" + getAstIdent(at)
	}
	return
}

func getAstSelectorExpr(s *ast.SelectorExpr) (rv string) {
	Sel := s.Sel.Name
	switch se := s.X.(type) {
	case *ast.Ident:
		rv = se.Name + "." + Sel
	}
	return
}

func getAstIdent(s *ast.Ident) string {
	return s.Name
}
