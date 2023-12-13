package day_13

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/13.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	blocks := parse(input)
	var first int

	for _, block := range blocks {
		first += value(block)
	}

	return first, 0
}
