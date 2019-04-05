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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	opts      Options
	fset      *token.FileSet
	pathList  []string
	logger    *zap.Logger
	sugar     *zap.SugaredLogger
	zapConfig zap.Config
)

//Options : Command Line Options
type Options struct {
	Path     string `short:"p" long:"path" description:"Project directory path" required:"true"`
	Debug    bool   `short:"v" long:"verbose" description:"Print debug messages"`
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

	var logLevel zap.AtomicLevel
	if opts.Debug != true {
		logLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	} else {
		logLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	var initialFields map[string]interface{}

	zapConfig = zap.Config{
		Level:             logLevel,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "type",
			TimeKey:        "",
			NameKey:        "",
			CallerKey:      "src",
			StacktraceKey:  "",
			LineEnding:     "\n",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.EpochMillisTimeEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths: []string{
			"stdout",
			"/tmp/logs",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: initialFields,
	}
}

func main() {
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar = logger.Sugar()

	logger.Info("Parsing project directory for golang files...")
	if err := getPathList(opts.Path, visit); err != nil {
		sugar.Errorf("filepath.Walk() returned %v\n", err)
	}

	fset = token.NewFileSet()

	if opts.AstDebug != true {
		logger.Info("Beginning AST Analysis...")
		tl, sm := parseDirFiles(fset)
		logger.Info("TYPE_LIST: ")
		logger.Debug(spew.Sdump(tl))
		logger.Info("STRUCT_DEFS: ")
		logger.Debug(spew.Sdump(sm))
		rl := relationMapper(sm, tl)
		logger.Info("RELATIONSHIPS: ")
		logger.Debug(spew.Sdump(rl))
		buildTemplate(sm, rl)
	} else {
		logger.Info("AST Debug Enabled!")
		debugParseDirFiles(fset)
	}
}

func parseDirFiles(f *token.FileSet) (TypeList, StructMap) {
	for _, pathVar := range pathList {
		prse, err := parser.ParseDir(f, pathVar, fileFilter, 0)
		if err != nil {
			sugar.Errorf("Error: ", err)
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
			sugar.Errorf("Error: ", err)
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
