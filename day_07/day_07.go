package day_07

import (
	"fmt"
	"os"
	"slices"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/07.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	hands := parse(input)

	slices.SortFunc(hands, compare)
	first := score(hands)

	return first, 0
}
