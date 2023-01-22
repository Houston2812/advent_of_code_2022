package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go

func check(e error)  {
	if e != nil {
		panic(e)
	}
}
func main()  {

	debug := flag.Bool("d", false, "Print debug messages")
	file_input := flag.String("f", "", "Path to file that should be read")
	flag.Parse()

	if len(*file_input) == 0 {
		os.Exit(0)
	}

	data, err := os.Open(*file_input)
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	counter := 1
	calorie := 0
	var calories []int
	calories = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			calories = append(calories, calorie)
			calorie = 0
			continue
		}

		tmp, _ := strconv.Atoi(line)
		calorie += tmp

		// fmt.Printf("Line %04d -> %d\n", counter, calorie)
		counter++
	}
	check(scanner.Err())

	sort.Slice(calories, func(p, q int) bool {
		return calories[p] > calories[q]
	})

	if *debug == true {
		for i := 0; i < len(calories); i++ {
			fmt.Printf("%04d -> %d\n", i, calories[i])
		}

	}

	// task 1 
	fmt.Printf("The most calories carried by elf - %d\n", calories[0])

	// task 2
	top_three_max := calories[0] + calories[1] + calories[2]
	fmt.Printf("The most calories carried by top 3 elfes - %d\n", top_three_max)

}