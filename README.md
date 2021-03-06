# gophermap
Golang Struct Relation Diagram

## To-Do

- [ ] walk ast doc for structs in files from dir and create config file `alpha-0.?.?`
    - [ ] create dry-run flag (generate txt output) `alpha-0.1.3`
        - [ ] refine text output `alpha-0.1.3`
        - [ ] create flag `alpha-0.1.3`
        - [ ] create logic for flag `alpha-0.1.3``
    - [ ] based on package type `alpha-0.1.?`
    - [ ] output to config file (graphviz) `alpha-0.1.3`
        - [ ] decide on graphing library `alpha-0.1.3`
        - [ ] fix empty structs `alpha-0.1.3`
        - [x] fix "chan<-" line ~168 `alpha-0.1.3`
        - [ ] fix possible type issues in relations `alpha-0.1.3`
        - [x] fix remove asterisk from relation list items `alpha-0.1.3`
    - [ ] structs from other locations `alpha-0.1.?`
- [ ] use config to genereate graph  `alpha-0.1.?`
- [ ] finish fixing getAstMapType to support all types for mtv `alpha-0.1.?`
- [x] fix type list -> type map `alpha-0.1.3`
- [ ] fix type map processing per type and default in aster.go `alpha-0.1.3`
- [ ] fix template object labeling [ struct, interface, etc. ] `alpha-0.1.3`
- [ ] add comments to dot.tmpl `alpha-0.1.3`


## Release Requirements

- alpha-0.1.3:
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
