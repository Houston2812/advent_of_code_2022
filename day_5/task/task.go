package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	// custom modules
	"queue/queue"
	"stack/stack"
)

// debug functions

var debug bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// debug function to print string
func debug_print(msg string) {
	if debug {
		fmt.Printf("[+] %s\n", msg)
	}
}

// debug function to print element of type ItemType
func debug_print_stack_element(msg stack.ItemType) {
	if debug {
		fmt.Printf("[+] Stack element: %c\n", msg)
	}
}

// debug function to print integer array
func debug_print_array(msgs []int) {
	if debug {
		fmt.Printf("[+] ")
		for _, msg := range msgs {
			fmt.Printf("%d, ", msg)
		}
		fmt.Println()
	}
}

// debug function to print stack
func debug_print_stack(msg stack.Stack) {
	if debug {
		fmt.Print("[+][+] ")
		for _, el := range msg.All() {
			fmt.Printf("%c ", el)
		}
		fmt.Println()
	}
}

// debug function to print queue
func debug_print_queue(msg queue.Queue) {
	if debug {
		fmt.Print("[+][+] ")
		for _, el := range msg.All() {
			fmt.Printf("%c ", el)
		}
		fmt.Println()
	}
}

// debug function to print element of type ItemType
func debug_print_queue_element(msg queue.ItemType) {
	if debug {
		fmt.Printf("[+] Queue element: %c\n", msg)
	}
}


// function to print a single stack
func print_stack(msg stack.Stack, id int) {
	fmt.Printf("Printing the stack %d: \n\t", id + 1)
	for _, el := range msg.All() {
		fmt.Printf("%c ", el)
	}
	fmt.Println()
}

// function to print elements of type ItemType
func print_stack_elements(msgs []stack.ItemType) {
	for _, msg := range msgs {
		fmt.Printf("%c", msg)
	}
	fmt.Println()
}

// function to print a single queue
func print_queue(msg queue.Queue, id int) {
	fmt.Printf("Printing the stack %d: \n\t", id + 1)
	for _, el := range msg.All() {
		fmt.Printf("%c ", el)
	}
	fmt.Println()
}

// function to print elements of type ItemType
func print_queue_elements(msgs []queue.ItemType) {
	for _, msg := range msgs {
		fmt.Printf("%c", msg)
	}
	fmt.Println()
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

// function to create initial state of the stack
func createtInitialStack(initial_state []string,) []stack.Stack {

	stacks := []stack.Stack{}

	for i := 0; i < len(initial_state); i++ {
		debug_print(initial_state[i])
	}

	debug_print("Stack elements")

	elements := strings.Fields(initial_state[len(initial_state) - 1])

	for _, element := range elements {
		value, err := strconv.Atoi(element)
		check(err)
		
		stacks = append(stacks, stack.Stack{})

		for i := len(initial_state) - 2; i >= 0; i-- {
			item := initial_state[i][1 + 4 * (value - 1)]

			// check if the single stack is finished. 
			if item != 32 {
				debug_print(fmt.Sprintf("Item: %d (bottom up)", i))
				debug_print(fmt.Sprintf("\tType of item: %T", item))
				debug_print(fmt.Sprintf("\tValue of item: %c", item))
				stacks[value - 1].Push(item)
			}
		}

		debug_print(fmt.Sprintf("Stack ID: %d", value ))
		debug_print_stack(stacks[value-1])
	}

	return stacks
}

// get regex instructions from the line of data
func getInstructions(line string) []int {
	r, _ := regexp.Compile("[0-9]+")

	// grep the instructions
	var tmp []string = r.FindAllString(line, 3)
	var instructions []int

	for _, value := range tmp {

		// convert to int
		int_val, err := strconv.Atoi(value)
		check(err)

		instructions = append(instructions, int_val)
	}

	// return array of instructions
	return instructions
}

func createtInitialQueue(initial_state []string) []queue.Queue{
	queues := []queue.Queue{}

	for i := 0; i < len(initial_state); i++ {
		debug_print(initial_state[i])
	}

	debug_print("Queue elements")

	elements := strings.Fields(initial_state[len(initial_state) - 1])

	for _, element := range elements {
		value, err := strconv.Atoi(element)
		check(err)
		
		queues = append(queues, queue.Queue{})

		for i := len(initial_state) - 2; i >= 0; i-- {
			item := initial_state[i][1 + 4 * (value - 1)]

			// check if the single stack is finished. 
			if item != 32 {
				debug_print(fmt.Sprintf("Item: %d (bottom up)", i))
				debug_print(fmt.Sprintf("\tType of item: %T", item))
				debug_print(fmt.Sprintf("\tValue of item: %c", item))
				queues[value - 1].PushTop(item)
			}
		}

		debug_print(fmt.Sprintf("Queue ID: %d", value ))
		debug_print_queue(queues[value-1])
	}

	return queues
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
		// create initial stack
		var stacks []stack.Stack
		var lineNum int

		for index, line := range lines {
			if len(line) == 0 {
				stacks = createtInitialStack(lines[:index])
				lineNum = index
				break
			}
		}
		
		// print the stack
		for index, stack := range stacks {
			print_stack(stack, index)
		}
		
		debug_print("Actions")
		
		// regex to match instructions
		
		for i := lineNum+1; i < len(lines); i++ {
			debug_print(lines[i])
			
			var instructions []int = getInstructions(lines[i])
			
			// instructions[0] = amount
			// instructions[1] = source
			// instructions[2] = destionation

			for instructions[0] > 0 {
				
				el := stacks[instructions[1]-1].Pop()
				debug_print_stack_element(*el)
				stacks[instructions[2] - 1].Push(*el)
				
				debug_print(fmt.Sprintf("Stack %d: ", instructions[1]))
				debug_print_stack(stacks[instructions[1] - 1])

				debug_print(fmt.Sprintf("Stack %d: ", instructions[2]))
				debug_print_stack(stacks[instructions[2] - 1])
				
				instructions[0]--
			}

			debug_print_array(instructions)
		}

		fmt.Printf("Final stack")
		for index, stack := range stacks {
			print_stack(stack, index)
		}

		// combine top elements of each stack in one array
		var answer []stack.ItemType
		for i := 0; i < len(stacks); i++ {
			answer = append(answer, *stacks[i].Pop())
		}

		fmt.Printf("Answer is: ")
		print_stack_elements(answer)

	} else if *task == "task2" {
		var queues []queue.Queue
		var lineNum int

		for index, line := range lines {
			if len(line) == 0 {
				queues = createtInitialQueue(lines[:index])
				lineNum = index
				break
			}
		}
		
		debug_print(fmt.Sprintf("Line num queue: %d", lineNum))
		// print the queue

		fmt.Printf("Initial state of queues\n\n")
		for index, queue := range queues {
			print_queue(queue, index)
		}

		debug_print("Actions")

		for i := lineNum+1; i < len(lines); i++ {
			debug_print(lines[i])
			
			var instructions []int = getInstructions(lines[i])
			
			// instructions[0] = amount
			// instructions[1] = source
			// instructions[2] = destionation

			crates := queue.Queue{}

			// collect the crates in the additional queue
			for instructions[0] > 0 {
				
				crate := queues[instructions[1] - 1].PopTop()

				debug_print("Crate")
				debug_print_queue_element(*crate)
				
				crates.PushTop(*crate)
				instructions[0]--
			}

			// take elements from the additional queue and add them to the destination queue
			for crates.Size() > 0 {
				queues[instructions[2]-1].PushTop(*crates.PopTop())
			}

			debug_print("Source queue")
			debug_print_queue(queues[instructions[1] - 1])
			debug_print("Destination queue")
			debug_print_queue(queues[instructions[2] - 1])
		}

		fmt.Printf("\nFinal state of queues\n\n")
		for index, queue := range queues {
			print_queue(queue, index)
		}

		// combine top elements of each queue in one array
		var answer []queue.ItemType
		for i := 0; i < len(queues); i++ {
			top_el := *queues[i].PopTop()
			if top_el != nil {
				answer = append(answer, top_el)
			}
		}

		fmt.Printf("\nAnswer is: ")
		print_queue_elements(answer)
	}
}
