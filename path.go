package main

import (
	"os"
	"path/filepath"
	"regexp"
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
