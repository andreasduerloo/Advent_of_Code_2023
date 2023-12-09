package day_09

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/09.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	lines := parse(input)

	var first int
	for _, line := range lines {
		first += nextValue(line)
	}

	var second int
	for _, line := range lines {
		second += previousValue(line)
	}

	return first, second
}
