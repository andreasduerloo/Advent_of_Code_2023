package day_11

import (
	"fmt"
	"os"

	"github.com/andreasduerloo/slicetools"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/11.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	image, emptyRows, emptyCols := parse(input)
	image1 := slicetools.MapSlice(image, func(g galaxy) galaxy { return g.trueCoordinates(emptyRows, emptyCols, 2) }) // Woo closure!

	var first int

	for _, galaxy := range image1 {
		for _, otherGalaxy := range image1 {
			first += distance(galaxy, otherGalaxy)
		}
	}

	first = first / 2 // We counted everything twice

	// Second star
	image2 := slicetools.MapSlice(image, func(g galaxy) galaxy { return g.trueCoordinates(emptyRows, emptyCols, 1000000) })

	var second int

	for _, galaxy := range image2 {
		for _, otherGalaxy := range image2 {
			second += distance(galaxy, otherGalaxy)
		}
	}

	second = second / 2

	return first, second
}
