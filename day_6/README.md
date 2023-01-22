# Solution of Day 6 of Advent of Code 20222

## Description
Implemented here generic set data structures in go.  
It has several concrete sets:
- ByteSet
- IntSet
- RuneSet
- StringSet

To generate concrete set using generic set following should be done:

- `//go:generate genny -in=set.go -out=set-string.go gen "Item=string Value=int"` - insert this line to the core file (set.go) 
    - _Note! double backslash (//) is mandatory_
- `export PATH=$PATH:$(go env GOPATH)/bin` - add the GOHOME/bin folder to the PATH
- `go generate` - generate concrete set data structures 

_Note! Link: https://flaviocopes.com/golang-data-structure-set/_
## Execution   
To run tasks execute following commands:  

- `go run task.go -d -f input_file -task task1`  - run task 1
- `go run task.go -d -f input_file -task task2`  - run task 2

>  By Huseyn Gambarov. 2022
