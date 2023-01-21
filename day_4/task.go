package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var debug bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func debug_print(msg string) {
	if debug {
		fmt.Printf("[+] %s\n", msg)
	}
}

func main()  {
	flag.BoolVar(&debug, "d", false, "Print debug messages")
	file_input := flag.String("f", "", "Path to file that should be read")
	task := flag.String("task", "", "The number of task")
	flag.Parse()

	if len(*file_input) == 0 || len(*task) == 0 {
		os.Exit(0)
	}

	data, err := os.Open(*file_input)
	check(err)

	debug_print("Debug started")

	scanner := bufio.NewScanner(data)
	counter := 0

	if *task == "task1" {
		for scanner.Scan() {
			line := scanner.Text()
			ranges := strings.Split(line, ",")
			first_assignment := strings.Split(ranges[0], "-")
			second_assignment := strings.Split(ranges[1], "-")
			
			debug_print(first_assignment[0])

			// Get integer values of the ranges 
			a0, err := strconv.Atoi(first_assignment[0])
			check(err)
			b0, err := strconv.Atoi(first_assignment[1])
			check(err)

			a1, err := strconv.Atoi(second_assignment[0])
			check(err)
			b1, err := strconv.Atoi(second_assignment[1])
			check(err)

			debug_print(line)
			
			// First case 
			// ----a0----b0----
			// -a1---------b1--
			if a0 >= a1 && b0 <= b1 {
				counter++
				debug_print("First in second")
				continue
			}
			
			// Second case
			// --a0-------b0--
			// -----a1--b1----
			if a1 >= a0 && b1 <= b0 {
				counter++
				debug_print("Second in first")
				continue
			}

			debug_print("No overlapping")
		}

		fmt.Printf("The amount of inclusive assignments: %d\n", counter)
	} else if *task == "task2" {
		line_number := 0
		for scanner.Scan() {
			line_number++
			line := scanner.Text()
			ranges := strings.Split(line, ",")
			first_assignment := strings.Split(ranges[0], "-")
			second_assignment := strings.Split(ranges[1], "-")
			
			debug_print(first_assignment[0])

			// Get integer values of the ranges 
			a0, err := strconv.Atoi(first_assignment[0])
			check(err)
			b0, err := strconv.Atoi(first_assignment[1])
			check(err)

			a1, err := strconv.Atoi(second_assignment[0])
			check(err)
			b1, err := strconv.Atoi(second_assignment[1])
			check(err)

			debug_print(line)
		
			// First case 
			// -a0--b0----------
			// --------a1---b1--
			if a0 > a1 && b0 > a1 && a0 > b1 && b0 > b1 {
				counter++
				continue
			}
			
			// First case 
			// --------a0---b0--
			// -a1--b1----------
			if a1 > a0 && b1 > a0 && a1 > b0 && b1 > b0 {
				counter++
				continue
			}

			debug_print("Overlapping")

 		}
		fmt.Printf("The amount of overlapping assignments: %d\n", line_number - counter)

	}
}