package day_03

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/03.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	schematic, numbers, gears := parse(input)

	var first int

	for _, number := range numbers {
		valid, value := checkNumber(number, schematic, gears)

		if valid {
			first += value
		}
	}

	var second int

	for _, gear := range gears {
		second += gear.ratio()
	}

	return first, second
}
