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
		for _, ppt := range sm[ps].Properties {
			//old := string(ppt[0])
			//new := strings.ToUpper(old)
			//appt := strings.Replace(ppt, old, new, 1)
			if child, ok := apptInTypeList(ppt, tl); ok != false {
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
