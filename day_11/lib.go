package day_11

import (
	"slices"
	"strings"

	"github.com/andreasduerloo/slicetools"
)

type galaxy struct {
	row int
	col int
}

func parse(input []byte) ([]galaxy, []int, []int) {
	lines := strings.Split(string(input), "\n")
	var out []galaxy
	var emptyRows, emptyCols []int

	for row, line := range lines {
		if !slicetools.AnySlice([]rune(line), func(r rune) bool { return r == '#' }) {
			emptyRows = append(emptyRows, row)
		}
		for column, r := range line {
			if r == '#' {
				out = append(out, galaxy{row: row, col: column})
			}
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		count := 0

		for _, row := range lines {
			if row != "" {
				if row[i] == '#' {
					count += 1
				}
			}
		}

		if count == 0 {
			emptyCols = append(emptyCols, i)
		}
	}

	return out, emptyRows, emptyCols
}

func (g galaxy) trueCoordinates(emptyRows, emptyCols []int) galaxy {
	rowsBefore := len(slicetools.FilterSlice(emptyRows, func(i int) bool { return i < g.row }))
	colsBefore := len(slicetools.FilterSlice(emptyCols, func(i int) bool { return i < g.col }))

	return galaxy{
		row: g.row + rowsBefore,
		col: g.col + colsBefore,
	}
}

func (g galaxy) bigCoordinates(emptyRows, emptyCols []int) galaxy {
	rowsBefore := len(slicetools.FilterSlice(emptyRows, func(i int) bool { return i < g.row }))
	colsBefore := len(slicetools.FilterSlice(emptyCols, func(i int) bool { return i < g.col }))

	return galaxy{
		row: g.row + ((1000000 - 1) * rowsBefore), // It's not row + factor * rows, but row + (factor - 1) * rows, it just worked out the same for 2.
		col: g.col + ((1000000 - 1) * colsBefore), // Because we already count the empty row/col in the original coordinate!
	}
}

func distance(g1, g2 galaxy) int {
	return (slices.Max([]int{g1.row, g2.row}) - slices.Min([]int{g1.row, g2.row})) + (slices.Max([]int{g1.col, g2.col}) - slices.Min([]int{g1.col, g2.col}))
}
