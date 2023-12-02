package day_02

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/02.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	games := parse(input)

	max := game{
		blue:  14,
		red:   12,
		green: 13,
	}

	first := filterReduce(games, max, possible, reduceId)

	second := mapReduce(games, gamePower)

	return first, second
}
