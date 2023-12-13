package day_13

import (
	"fmt"
	"os"

	"github.com/andreasduerloo/slicetools"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/13a.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	blocks := parse(input)
	var first int

	for _, block := range blocks {
		val, _, _ := value(block)
		first += val
	}

	smudgedMirrors := slicetools.MapSlice(blocks, smudge)

	var second int

	for _, mirror := range smudgedMirrors {
		second += mirror
	}

	return first, second
}
