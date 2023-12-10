package day_10

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/10.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	start, pipes, width := parse(input)

	stepCount := steps(start)
	first := (stepCount + 1) / 2

	second := scan(pipes, width)

	return first, second
}
