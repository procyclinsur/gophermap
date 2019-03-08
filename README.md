# gophermap
Golang Struct Relation Diagram

## To-Do

- [ ] walk ast doc for structs in files from dir and create config file `alpha-0.?.?`
    - [ ] create dry-run flag (generate txt output)
        - [ ] refine text output `alpha-0.1.?`
        - [ ] create flag `alpha-0.1.1`
        - [ ] create logic for flag `alpha-0.1.?`
    - [ ] based on package type `alpha-0.1.?`
    - [ ] output to config file (graphviz?) `alpha-0.1.?`
        - [ ] decide on graphing library
    - [ ] structs from other locations `alpha-0.1.1`
    - [x] exclude golang test files `alpha-0.1.1`
    - [ ] create inter struct relations map `alpha-0.1.1`
    - [ ] decode parameter type names `alpha-0.1.1`
- [ ] use config to genereate graph  `beta-0.?.?`

## Release Requirements

- alpha-0.1.1:
    - [ ] No bugs in running code
    - [ ] Above labeled to-do's finished

## Build

```bash
go build -o ./bin/gophermap
```
