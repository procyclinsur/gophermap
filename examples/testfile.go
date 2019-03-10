package main

import "fmt"

type test struct {
	a string
	b int
	c string
}

type math struct {
	x float32
	y float64
	z int64
}

type mixed struct {
	obj1 *test
	obj2 *math
}

type empty struct {
}

type structMap struct {
	StructList map[string]StructDef
}

//StructDef v1.0
type StructDef struct {
	//Name of struct
	Name string
	//Map of property-name:property-type
	Properties map[string]string
	//List of structs contained
	Contains []string
}

func main() {
	struct1 := StructDef{
		"fakeStruct1",
		map[string]string{
			"fakeProperty1": "int",
			"fakeProperty2": "string",
			"fakeProperty3": "mixed",
		},
		[]string{"mixed"},
	}
	var struct2 structMap
	struct2.StructList = make(map[string]StructDef)
	struct2.StructList["fakeStruct1"] = struct1
	fmt.Println(struct2)
	struct3 := StructDef{
		"fakeStruct2",
		map[string]string{
			"fakeProperty1": "float32",
			"fakeProperty2": "math",
			"fakeProperty3": "empty",
		},
		[]string{"math", "empty"},
	}
	struct2.StructList["fakeStruct2"] = struct3
	fmt.Println(struct2)
}

//{
//	map[
//	        fakeStruct1:{
//			fakeStruct1
//			map[
//			        fakeProperty1:int
//				fakeProperty2:string
//				fakeProperty3:mixed
//			]
//			[
//			        mixed
//			]
//		}
//		fakeStruct2:{
//			fakeStruct2
//			map[
//			        fakeProperty1:float32
//				fakeProperty2:math
//				fakeProperty3:empty
//			]
//			[
//			        math
//				empty
//			]
//		}
//	]
//}
