package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"set/set"
)

var debug bool
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

func debug_print_runes(msgs []rune)  {
	if debug {
		for _, msg := range msgs {
			fmt.Printf("[+] %c\n", msg)
		}
	}
}

func debug_print_rune(msg rune)  {
	if debug {
		fmt.Printf("[+] %c\n", msg)
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


func main()  {
	
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

		for _, line := range lines {

			var package_len int = len(line)
			var signal string = line
			var start_marker int = 0
	
			for i := 0; i < package_len-3; i++  {
				
				var markers set.RuneSet = set.RuneSet{}
				var success_marker bool = true
	
				for _, marker := range signal[i : i + 4] {
	
					if markers.Has(marker) {
						start_marker++
						success_marker = false
						break
					}
	
					markers = *markers.Add(marker)
				}
				
				if success_marker {
					start_marker += 4
					break
				}
	
				debug_print("Printing runes")
				debug_print_runes(markers.Runes())
			}
	
			fmt.Printf("The start-of-packer marker is at: %d\n", start_marker)
		}

	} else if *task == "task2" {
		
		for _, line := range lines {

			var package_len int = len(line)
			var signal string = line
			var start_marker int = 0
	
			for i := 0; i < package_len-13; i++  {
				
				var markers set.RuneSet = set.RuneSet{}
				var success_marker bool = true
	
				for _, marker := range signal[i : i + 14] {
	
					if markers.Has(marker) {
						start_marker++
						success_marker = false
						break
					}
	
					markers = *markers.Add(marker)
				}
				
				if success_marker {
					start_marker += 14
					break
				}
	
				debug_print("Printing runes")
				debug_print_runes(markers.Runes())
			}
	
			fmt.Printf("The start-of-message marker is at: %d\n", start_marker)
		}
	}
}