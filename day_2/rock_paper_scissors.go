package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

func main()  {
	
	flag.BoolVar(&debug, "d", false, "Print debug messages")
	file_input := flag.String("f", "", "Path to file that should be read")
	task := flag.String("task", "", "The number of task")
	flag.Parse()

	if len(*file_input) == 0 || len(*task) == 0{
		os.Exit(0)
	}

	data, err := os.Open(*file_input)
	check(err)

	debug_print("[+]Debug started")

	scanner := bufio.NewScanner(data)
	counter := 1
	
	scores := make(map[string]int)
	score := 0

	scores["A"] = 1
	scores["B"] = 2
	scores["C"] = 3
	scores["X"] = 1
	scores["Y"] = 2
	scores["Z"] = 3

	if *task == "task1" {

		for scanner.Scan() {
			line := scanner.Text()
			
			plays := strings.Fields(line)
	
			score += scores[plays[1]]
	
			if plays[0] == "A" {
				if plays[1] == "X" {
	
					score += 3 
					debug_print("Draw")
	
				} else if plays[1] == "Y" {
	
					score += 6
					debug_print("Win")
				
				} else if plays[1] == "Z" {
	
					debug_print("Lose")
	
				}
			 } else if plays[0] == "B" {
				if plays[1] == "X" {
	
					debug_print("Lose")
	
				} else if plays[1] == "Y" {
					
					score += 3
					debug_print("Draw")
	
				} else if plays[1] == "Z" {
	
					score += 6
					debug_print("Win")
	
				}
			} else if plays[0] == "C" {
				if plays[1] == "X" {
	
					score += 6
					debug_print("Win")
	
				} else if plays[1] == "Y" {
	
					debug_print("Lose")
	
				} else if plays[1] == "Z" {
					
					score += 3
					debug_print("Draw")
	
				}
			}
	
			counter++
		}

	}
	
	if *task == "task2" {
		for scanner.Scan() {
			line := scanner.Text()
			
			plays := strings.Fields(line)
	
			if plays[1] == "X" {
				
				debug_print("Lose")

				if plays[0] == "A" {
					score += scores["Z"]
				} else if plays[0] == "B" {
					score += scores["X"]
				} else if plays[0] == "C" {
					score += scores["Y"]
				}

			} else if plays[1] == "Y" {
				
				debug_print("Draw")

				score += 3
				score += scores[plays[0]]

			} else if plays[1] == "Z" {
				
				debug_print("Win")

				score += 6

				if plays[0] == "A" {
					score += scores["Y"]
				} else if plays[0] == "B" {
					score += scores["Z"]
				} else if plays[0] == "C" {
					score += scores["X"]
				}
			}
		}
	}

	fmt.Printf("The score for game is %d\n", score)
}