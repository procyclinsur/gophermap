package main

import (
	"strings"
)

//RelationList v1.0
type RelationList []Relation

//Relation v1.0
type Relation struct {
	//Containing Struct
	Parent string
	//Contained Structs
	Children []string
}

func relationMapper(sm StructMap, tl TypeList) RelationList {
	var relationList = RelationList{}
	for ps := range sm {
		var ptList []string
		sugar.Debugf("Parent Struct: %s", ps)
		for ppn, ppt := range sm[ps].Properties {
			sugar.Debugf("    Checking Property: %s", ppn)
			sugar.Debugf("        Type: %s", ppt)
			if child, ok := apptInTypeList(ppt, tl); ok != false {
				sugar.Debugf("            Type '%s' matches requirement adding to list", ppt)
				ptList = append(ptList, child)
			}
		}
		relationList = append(relationList, Relation{
			ps,
			ptList,
		})
	}
	return relationList
}

func apptInTypeList(p string, tl TypeList) (child string, ok bool) {
	for _, ps := range tl {
		ap := strings.ToLower(p)
		aps := strings.ToLower(ps)
		if ap == aps {
			child = ps
		} else if string(ap[0]) == "*" {
			child = p
		}
		if child != "" {
			ok = true
			return
		}
	}
	child = ""
	ok = false
	return
}
