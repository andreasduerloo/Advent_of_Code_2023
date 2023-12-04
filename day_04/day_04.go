package day_04

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/04.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	games := parse(input)

	// First star

	var first int
	for _, game := range games {
		first += game.score()
	}

	// Second star

	cache := make(map[int]int)

	var second int
	for i := len(games); i >= 1; i-- { // Working backwards, see readme
		second += games[i-1].newTickets(len(games), cache)
	}

	return first, second
}
