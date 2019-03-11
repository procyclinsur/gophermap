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

func relationMapper(sm StructMap) RelationList {
	var relationList = RelationList{}
	for ps := range sm {
		var ptList []string
		for _, ppt := range sm[ps].Properties {
			old := string(ppt[0])
			new := strings.ToUpper(old)
			appt := strings.Replace(ppt, old, new, 1)
			if child := apptInStructMap(appt, sm); child != false {
				ptList = append(ptList, appt)
			}
		}
		relationList = append(relationList, Relation{
			ps,
			ptList,
		})
	}
	return relationList
}

func apptInStructMap(p string, sm StructMap) bool {
	for ps := range sm {
		if p == ps {
			return true
		}
	}
	return false
}
