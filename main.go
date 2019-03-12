package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	flags "github.com/jessevdk/go-flags"
)

var (
	opts     Options
	fset     *token.FileSet
	pathList []string
	//debug bool
)

//Options : Command Line Options
type Options struct {
	Path     string `short:"p" long:"path" description:"Project directory path" required:"true"`
	AstDebug bool   `short:"a" long:"astdebug" description:"Print AST file"`
	HelpFlag bool   `short:"h" long:"help" description:"Print this help message"`
}

func init() {
	prsr := flags.NewParser(&opts, flags.Default)
	if _, err := prsr.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			panic("Input required parameters.")
		} else {
			errMsg := fmt.Sprintf("%s\n\tUse the -h or --help flag for more options.", err)
			panic(errMsg)
		}

	} else if err != nil {
		panic(err)
	}
}

func main() {
	if err := getPathList(opts.Path, visit); err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	fset = token.NewFileSet()

	if opts.AstDebug != true {
		tl, sm := parseDirFiles(fset)
		rl := relationMapper(sm, tl)
		fmt.Println("#####STRUCT_DEFS#####")
		spew.Dump(sm)
		fmt.Println("######TYPE_LIST######")
		spew.Dump(tl)
		fmt.Println("####RELATIONSHIPS####")
		spew.Dump(rl)
		//_ = rl
	} else {
		debugParseDirFiles(fset)
	}
}

func parseDirFiles(f *token.FileSet) (TypeList, StructMap) {
	for _, pathVar := range pathList {
		prse, err := parser.ParseDir(f, pathVar, fileFilter, 0)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		for _, pkgItem := range prse {
			ast.Walk(VisitorFunc(FindTypes), pkgItem)
		}
	}
	return getWalkOutput()
}

func debugParseDirFiles(f *token.FileSet) {
	for _, pathVar := range pathList {
		prse, err := parser.ParseDir(f, pathVar, fileFilter, 0)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		for _, pkgItem := range prse {
			ast.Fprint(os.Stdout, f, pkgItem, func(name string, value reflect.Value) bool {
				if ast.NotNilFilter(name, value) {
					return value.Type().String() != "*ast.Object"
				}
				return false
			})
		}
	}
}
