package day_01

import (
	"advent2023/helpers"
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/01.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	list := parse(input)
	first := helpers.SumSlice(list)

	list = parse2(input)
	second := helpers.SumSlice(list)

	// Test the naive solution
	fmt.Println(naive(input))

	return first, second
}
