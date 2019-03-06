package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func visit(path string, f os.FileInfo, err error) error {
	r, _ := regexp.Compile(`(^.*/\..*$)`)
	if f.IsDir() {
		match := r.MatchString(path)
		if match != true {
			fmt.Printf("%s\n", path)
			return nil
		}
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
