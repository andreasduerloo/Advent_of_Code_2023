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

	slices.SortFunc(hands, sorter(false))
	first := score(hands)

	for _, hand := range hands {
		hand.reidentify()
	}

	slices.SortFunc(hands, sorter(true))
	second := score(hands)

	return first, second
}
