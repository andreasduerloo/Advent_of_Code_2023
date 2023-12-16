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

	/*
		for _, line := range rocks {
			fmt.Println(string(line))
		}
	*/

	return first, 0
}
