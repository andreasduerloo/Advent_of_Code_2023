package day_06

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/06.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	times, distances := parse(input)

	first := 1

	for i, time := range times {
		first *= winningRange(time, distances[i])
	}

	time := concatNum(times)
	distance := concatNum(distances)

	second := winningRange(time, distance)

	return first, second
}
