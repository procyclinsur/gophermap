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
    - [ ] structs from other locations `alpha-0.1.2`
    - [x] exclude golang test files `alpha-0.1.1`
    - [x] create inter struct relations map `alpha-0.1.1`
        - [x] internal structs
        - [x] external structs
            - [x] create a list of types using `*ast.TypeSpec` `alpha-0.1.1`
                - break walkStructSpec into walkTypeSpec and walkStructSpec
                - parse all types into list for reference by relations.go
            - [x] parse external stucts on ?(StarExpr)
    - [ ] allow external non StarExpr types to be recognized `alpha-0.1.2`
    - [x] organize data into structs `alpha-0.1.1`
    - [x] decode parameter type names `alpha-0.1.1`
    - [x] changed flag parser to require arguments `alpha-0.1.1`
    - [x] need to parse for `*ast.MapType` on parameter values `alpha-0.1.1`
        - `./bin/gophermap -p ../../atp/line-sender/ -a | grep -C 30 words | head -60`
    - [ ] implement zap logger `alpha-0.1.1`
- [ ] use config to genereate graph  `beta-0.?.?`
- [x] add usage instructions `alpha-0.1.1`

## Release Requirements

- alpha-0.1.1:
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
