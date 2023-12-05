package day_05

import (
	"fmt"
	"os"
	"slices"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/05.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	seeds, farmMaps := parse(input)

	for _, fmap := range farmMaps { // Do this with a map function?
		var newVals []int
		for _, seed := range seeds {
			newVals = append(newVals, transform(seed, fmap))
		}
		seeds = newVals
	}

	first := slices.Min(seeds)

	return first, 0
}
