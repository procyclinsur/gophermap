# gophermap
Golang Struct Relation Diagram

## To-Do

- [x] find paths in directory (paths.go)
- [x] display ast document for debug (aster.go)
    - [x] single-file
    - [x] git directory
    - [x] integrate paths.go code
- [ ] walk ast document for files in dir for structs and create config (main.go)
    - [x] top-level structs
    - [ ] based on package type
    - [ ] structs from other locations
    - [x] integrate paths.go code
    - [x] integrate aster.go debug code
    - [ ] Create inter struct relations
- [ ] use config to genereate graph
