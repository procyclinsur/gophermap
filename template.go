package main

import (
	"os"
	"text/template"
)

type erdDiagram struct {
	StructMap
	RelationList
}

func buildTemplate(sm StructMap, rl RelationList) {
	erd := erdDiagram{
		sm,
		rl,
	}

	hdr, err := template.New("dot.tmpl").ParseGlob("./templates/*.tmpl")
	if err != nil {
		sugar.Errorf("Failed to create template: %s\n", err)
	}

	err = hdr.Execute(os.Stdout, erd)
	if err != nil {
		sugar.Errorf("Failed to parse structMap: %s\n", err)
	}
}
