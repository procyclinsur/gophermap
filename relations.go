package main

var relationList = RelationList{}

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
	return relationList
}
