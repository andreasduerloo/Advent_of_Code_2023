package day_14

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/14.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	rocks := parse(input)

	rollNorth(&rocks)
	first := calculateLoad(rocks)

	rep := cycle(1_000_000_000, 1, &rocks)
	rocks = parse(input)

	fmt.Println(rep)

	cycle2(rep+30, 0, &rocks)
	second := calculateLoad(rocks)

	return first, second
}
