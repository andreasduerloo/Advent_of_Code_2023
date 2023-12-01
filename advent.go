package main

import (
	"advent2023/day_01"
	"advent2023/day_02"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument was passed - exiting.")
		return
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The argument is not an integer - exiting.")
	}

	fmt.Println("Solutions for day", day)

	switch day {
	case 1:
		first, second := day_01.Solve()
		fmt.Println(first, second)
	case 2:
		first, second := day_02.Solve()
		fmt.Println(first, second)
	default:
		fmt.Println("That's either not a valid day, or it has not been solved (yet!)")
	}

}
