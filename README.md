# gophermap
Golang Struct Relation Diagram

## To-Do

- [ ] walk ast doc for structs in files from dir and create config file `alpha-0.?.?`
    - [ ] create dry-run flag (generate txt output)
        - [ ] refine text output `alpha-0.1.?`
        - [ ] create flag `alpha-0.1.?`
        - [ ] create logic for flag `alpha-0.1.?`
    - [ ] based on package type `alpha-0.1.?`
    - [ ] output to config file (graphviz?) `alpha-0.1.?`
        - [ ] decide on graphing library
    - [ ] structs from other locations `alpha-0.1.?`
    - [ ] allow external non StarExpr types to be recognized `alpha-0.1.2`
    - [x] implement zap logger `alpha-0.1.2`
- [ ] use config to genereate graph  `beta-0.?.?`
- [x] add \*ast.FuncType and \*ast.SelectorExpr to aster.go `alpha-0.1.2`
- [x] refactor processing for aster.go case statements to 'fieldType = someFunction(inputs)' `alpha-0.1.2`
- [x] create aster.go func getUndeterminedType to reduce code redundancy `alpha-0.1.2`
- [ ] fix getAstMapType to support all types for mtv `alpha-0.1.2`
- [ ] fix aster.go to allow support for directly nested structs `alpha-0.1.?`

## Release Requirements

- alpha-0.1.2:
    - [ ] No bugs in running code
    - [ ] Above labeled to-do's finished

## Build

```bash
go build -o ./bin/gophermap
```

## Usage

Print structs + struct properties from all files in a project directory.
```bash
./bin/gophermap -p <path-to-project>
```

Print ast file for debugging.
```bash
./bin/gophermap -p <path-to-project> -a
```

Print detailed debugging logs.
```bash
./bin/gophermap -p <path-to-project> -v
```

Bash function to search output ast file for a specific struct.
```bash
function getStruct {
    if [ -z $GOPHERREPO ]; then
        echo "Set path to gophermap repository to GOPHERREPO"
    elif [ -z $1 ] || [ -z $2 ] || [ -z $3 ]; then
        echo "USAGE: getStruct <struct-name> <output-length> <file-path>"
    else
        cd $GOPHERREPO
        ./bin/gophermap -p $3 -v -a | grep -E -B 3 -A $2 "^[ ].*[0-9]  (\.  ){8}Name: \"$1\""
    fi
}
```
