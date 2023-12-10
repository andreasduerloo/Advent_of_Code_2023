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

	start, pipes := parse(input)

	stepCount := steps(start)
	first := (stepCount + 1) / 2

	second := len(pipes) - (bfs(pipes[0]) + stepCount)
	// The problem is that squeezing between the pipes is allowed - think of a way to work with that

	return first, second
}
