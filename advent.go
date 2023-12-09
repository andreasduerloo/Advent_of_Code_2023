package main

import (
	"advent2023/day_01"
	"advent2023/day_02"
	"advent2023/day_03"
	"advent2023/day_04"
	"advent2023/day_05"
	"advent2023/day_06"
	"advent2023/day_07"
	"advent2023/day_08"
	"advent2023/day_09"
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

	solved := []func() (int, int){day_01.Solve, day_02.Solve, day_03.Solve, day_04.Solve, day_05.Solve, day_06.Solve, day_07.Solve, day_08.Solve, day_09.Solve}

	if day <= len(solved) {
		fmt.Println("Solutions for day", day)
		first, second := solved[day-1]()
		fmt.Println(first, second)
	} else {
		fmt.Println("That's either not a valid day, or it has not been solved (yet!)")
	}
}
