# gophermap
Golang Struct Relation Diagram

## To-Do

- [x] find paths in directory (paths.go)
- [x] display ast document for debug (aster.go)
    - [x] single-file
    - [ ] git directory
    - [ ] integrate paths.go code
- [ ] walk ast document for files in dir for structs and create config (main.go)
    - [x] top-level structs
    - [ ] based on package type
    - [ ] structs from other locations
    - [ ] integrate paths.go code
    - [ ] Create inter struct relations
- [ ] use config to genereate graph
