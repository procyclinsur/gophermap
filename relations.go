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
	Children map[string]string
}

func relationMapper(sm StructMap, tl TypesMap) RelationList {
	var relationList = RelationList{}
	for ps := range sm {
		ptList := map[string]string{}
		sugar.Debugf("Parent Struct: %s", ps)
		for ppn, ppt := range sm[ps].Properties {
			sugar.Debugf("    Checking Property: %s", ppn)
			sugar.Debugf("        Type: %s", ppt)
			cn, ct, ok := apptInTypesMap(ppt, tl)
			//sugar.Debugf("CN: %s", cn)
			//sugar.Debugf("CT: %s", ct)
			//sugar.Debugf("OK: %s", ok)
			if ok != false {
				sugar.Debugf("            Type '%s' matches requirement adding to list", ppt)
				ptList[cn] = ct
			}
		}
		relationList = append(relationList, Relation{
			ps,
			ptList,
		})
	}
	return relationList
}

func apptInTypesMap(p string, tl TypesMap) (cn, ct string, ok bool) {
	for tn, tt := range tl {
		ap := strings.ToLower(p)
		if string(ap[0]) == "*" {
			ap = ap[1:]
		}
		atn := strings.ToLower(tn)
		if ap == atn {
			cn, ct = tn, tt
		}
		if cn != "" {
			ok = true
			return
		}
	}
	cn, ct = "", ""
	ok = false
	return
}
