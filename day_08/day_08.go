package day_08

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/08.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	instructions, nodes := parse(input)

	first := steps(nodes["AAA"], nodes["ZZZ"], instructions)

	return first, 0
}
