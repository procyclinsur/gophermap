# gophermap
Golang Struct Relation Diagram

## To-Do

- [x] find paths in directory (paths.go)  `alpha-0.1.0 release`
- [x] display ast document for debug (aster.go)  `alpha-0.1.0 release`
    - [x] single-file  `alpha-0.1.0 release`
    - [x] git directory  `alpha-0.1.0 release`
    - [x] integrate paths.go code  `alpha-0.1.0 release`
- [ ] walk ast document for files in dir for structs and create config (main.go)
    - [x] top-level structs `alpha-0.1.0 release`
    - [ ] based on package type
    - [ ] structs from other locations
    - [ ] exclude golang test files
    - [x] integrate paths.go code  `alpha-0.1.0 release`
    - [x] integrate aster.go debug code  `alpha-0.1.0 release`
    - [ ] Create inter struct relations
- [ ] use config to genereate graph

## Release Requirements

- alpha-0.1.0:
    - [ ] No bugs in running code
    - [x] Above labeled to-do's finished
