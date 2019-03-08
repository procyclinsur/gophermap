package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getPathList(path string, walkFN filepath.WalkFunc) error {
	if err := filepath.Walk(path, walkFN); err != nil {
		return err
	}
	return nil
}

func visit(path string, f os.FileInfo, err error) error {
	r, _ := regexp.Compile(`(^.*/\.[^\.].*$)`)
	if f.IsDir() {
		match := r.MatchString(path)
		if match != true {
			pathList = append(pathList, path)
			return nil
		}
	}
	return nil
}

func fileFilter(f os.FileInfo) (rtrn bool) {
	match := strings.Contains(f.Name(), "test.go")
	if match != true {
		rtrn = true
	} else {
		rtrn = false
	}
	fmt.Printf("%s : %t\n", f.Name(), rtrn)
	return rtrn
}
