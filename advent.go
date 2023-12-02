package main

import (
	"advent2023/day_01"
	"advent2023/day_02"
	"advent2023/day_03"
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

	solved := []func() (int, int){day_01.Solve, day_02.Solve, day_03.Solve}

	if day <= len(solved) {
		fmt.Println("Solutions for day", day)
		first, second := solved[day-1]()
		fmt.Println(first, second)
	} else {
		fmt.Println("That's either not a valid day, or it has not been solved (yet!)")
	}
}
