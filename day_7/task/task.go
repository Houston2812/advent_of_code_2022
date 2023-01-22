package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tree/tree"
)


var debug bool
var fileSystem tree.FileSystem
var parents []tree.EntityDescriptor
var sizes []int 

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func debug_print(msg string)  {
	if debug {
		fmt.Printf("[+] %s\n", msg)
	}
}

func debug_print_parents(parents []tree.EntityDescriptor)  {
	for _, parent := range parents {
		tree.PrintEntity(parent)
	}
}
// function to read the lines
func readLines(data *os.File) ([]string, error) {
    var lines []string
    scanner := bufio.NewScanner(data)

    for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
    }

    return lines, scanner.Err()
}

func createTree(line string) {
	args := strings.Fields(line)

	if debug {
		fmt.Print("[+] Args: ")
		for _, arg := range args {
			fmt.Printf("%s; ", arg)
		}
		fmt.Printf("\n")
	}

	
	if args[0] == "$" {
		if args[1] == "cd" {
			if args[2] == "/" {
				parents = append(parents, tree.EntityDescriptor{"/", "dir", 0})
				fileSystem.Insert(parents[0])
				debug_print("Done /")
				
			} else if args[2] == ".." {
				parents = parents[:len(parents) - 1]
				debug_print("Done ..")
			} else {
				parents = append(parents, tree.EntityDescriptor{args[2], "dir", 0})
				debug_print("Printing parents")
				debug_print_parents(parents)
				debug_print("Done !?")

			}
		} else if args[1] == "ls" {
			debug_print("Done ls")
			// fmt.Printf("Performed ls")
		} 
	} else {
		tmp_parents := make([]tree.EntityDescriptor, 0)
		tmp_parents = append(tmp_parents, parents[len(parents) - 1])
		
		debug_print_parents(tmp_parents)
		if args[0] == "dir" {
			fileSystem.Insert(tree.EntityDescriptor{args[1], args[0], 0}, tmp_parents...)

			debug_print("Done dir")

		} else {
			size, err := strconv.Atoi(args[0])
			sizes = append(sizes, size)

			check(err)
			fileSystem.Insert(tree.EntityDescriptor{args[1], "file", size}, tmp_parents...)
			
			debug_print("Done file")

		}
	}
	

}
func main() {

	// define CLI flags
	flag.BoolVar(&debug, "d", false, "Print debug messages")
	file_input := flag.String("f", "", "Path to file that should be read")
	task := flag.String("task", "", "The number of task (task1 or task2)")

	flag.Parse()

	// validate input
	if len(*file_input) == 0 || len(*task) == 0 {
		os.Exit(0)
	}

	// read the file
	data, err := os.Open(*file_input)
	check(err)
	defer data.Close()

	debug_print("Debug started")

	lines, err := readLines(data)
	check(err)
	
	if *task == "task1" {
		
	
		// fileSystem.Insert(tree.EntityDescriptor{"/", "dir", 0})
	
		
		for _, line := range lines {
			// if index == 8 {
			// 	break
			// }
			debug_print(line)
			createTree(line)
		}
		

		// fmt.Printf("Printing the tree\n")
		fileSystem.PrintInorder(fileSystem.Root, 0)

		// var size int = fileSystem.GetTotalSize(fileSystem.Root)
		var size int = 0
		for _, sz := range sizes {
			size += sz
		}
		fmt.Printf("Total size is: %d\n", size)

	}
}
