package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var debug bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convert(el byte) int {
	priority := 0
	if int(el) > 64 && int(el) < 91 {
		priority = int(el) - 65 + 27
	} else {
		priority = int(el) - 97 + 1
	}
	return priority
}

func debug_print(msg string) {
	if debug {
		fmt.Printf("[+] %s\n", msg)
	}
}

func debug_print_rune(msg rune) {
	if debug {
		fmt.Printf("[+] %c\n", msg)
	}
}

func debug_print_map(msg map[byte]int) {
	if debug{
		for key, element := range msg {
			fmt.Printf("Key: %c => Element: %d\n", key, element)
		}
	}
}


func main() {

	flag.BoolVar(&debug, "d", false, "Print debug messages")
	file_input := flag.String("f", "", "Path to file that should be read")
	task := flag.String("task", "", "The number of task")
	flag.Parse()

	// if len(*file_input) == 0 || len(*task) == 0
	if len(*file_input) == 0 {
		os.Exit(0)
	}

	data, err := os.Open(*file_input)
	check(err)

	debug_print("Debug started")

	scanner := bufio.NewScanner(data)
	sum := 0
	counter := 0

	if *task == "task1" {

		for scanner.Scan() {
			// read content of a rucksack
			line := scanner.Text()
			counter ++
			// divide content of rucksack into 2 parts
			first_half := line[:int(len(line)/2)]
			second_half := line[int(len(line)/2):]
	
			// debug messages
			if len(first_half) == len(second_half) {
				debug_print("Length is not equal")
			}
			debug_print(fmt.Sprintf("first_half: %s; second_half: %s", first_half, second_half))
			debug_print(line)
			
			is_found := false
	
			for i := 0; i < len(first_half); i++ {
				for j := 0; j < len(second_half); j++ {
					if (first_half[i] == second_half[j]) {
						fmt.Printf("The mistaken element is: %c\n", second_half[j])
						is_found = true
						sum += convert(second_half[j])
						break
					}	
				}
				if is_found {
					break
				}
			}			
		}
	
		fmt.Printf("Sum is the %d\n; Counter: %d\n", sum, counter)

	} else if *task == "task2" {
		for {
			scanner.Scan()
			line1 := scanner.Text()


			if len(line1) == 0 {
				break
			}

			scanner.Scan()
			line2 := scanner.Text()

			if len(line2) == 0 {
				break
			}

			scanner.Scan()
			line3 := scanner.Text()
			
			if len(line3) == 0 {
				break
			}

			counts := make(map[byte]int)

			for i := 0; i < len(line1); i++ {
				counts[line1[i]]=1
			}
			// debug_print_map(counts)


			for i := 0; i < len(line2); i++ {
				if counts[line2[i]] == 1 {
					counts[line2[i]]=2
				}
			}
			// debug_print_map(counts)


			for i := 0; i < len(line3); i++ {
				if counts[line3[i]] == 2 {
					counts[line3[i]]=3
				}
			}

			for key, element := range counts {
				if element == 3 {
					sum += convert(key)
				}
			}
			debug_print_map(counts)
			// debug_print(fmt.Sprintf("%s : %s : %s", line1, line2, line3))
		}
		fmt.Printf("The sum is %d\n", sum)
	}

}
