package day_03

import (
	"regexp"
	"strconv"
	"strings"
)

type number struct {
	begin int
	end   int
	value int
}

type grid struct {
	width  int
	height int
	data   []rune
}

type gear struct {
	neighbors []*number
}

func parse(input []byte) (grid, []number, map[int]gear) {
	var gridvalue []rune
	var numbers []number
	gears := make(map[int]gear)

	lines := strings.Split(string(input), "\n")
	numRe := regexp.MustCompile(`\d+`)

	// Set the width and height variables - we expect 140 x 140
	width := len(lines[0])
	height := len(lines) - 1 // Trim the empty line

	for i, line := range lines {
		numIndexes := numRe.FindAllIndex([]byte(line), -1)
		numMatches := numRe.FindAllString(line, -1)

		for j, match := range numMatches {
			val, _ := strconv.Atoi(match)

			numbers = append(numbers, number{
				begin: (i * width) + numIndexes[j][0],
				end:   (i * width) + numIndexes[j][1],
				value: val,
			})
		}

		for j, r := range line {
			gridvalue = append(gridvalue, r)
			if r == '*' {
				gears[i*width+j] = gear{}
			}
		}
	}

	outgrid := grid{
		width:  width,
		height: height,
		data:   gridvalue,
	}

	return outgrid, numbers, gears
}

func checkNumber(num number, schem grid, gears map[int]gear) (bool, int) {
	// Build a slice of neighbor indexes
	var candidates []int

	// Look at the row above
	if num.begin >= schem.width {

		// Can we look up and left?
		if num.begin%schem.width != 0 {
			candidates = append(candidates, num.begin-(schem.width+1))
		}

		// Can we look up and right?
		if num.end%schem.width != schem.width {
			candidates = append(candidates, num.end-schem.width)
		}

		// Look up
		for i := num.begin - schem.width; i < num.end-schem.width; i++ {
			candidates = append(candidates, i)
		}
	}

	// Look at the same row
	if num.begin > 0 {
		candidates = append(candidates, num.begin-1)
	}

	if num.end%schem.width != schem.width-1 {
		candidates = append(candidates, num.end)
	}

	// Look at the row below
	if num.begin < (schem.width*schem.height)-schem.width {

		// Can we look down and left?
		if num.begin%schem.width != 0 {
			candidates = append(candidates, num.begin+(schem.width-1))
		}

		// Can we look down and right?
		if num.end%schem.width != schem.width {
			candidates = append(candidates, num.end+schem.width)
		}

		// Look down
		for i := num.begin + schem.width; i < num.end+schem.width; i++ {
			candidates = append(candidates, i)
		}
	}

	for _, cand := range candidates {
		if isSymbol(schem.data[cand]) {
			if schem.data[cand] == '*' {
				// This is a gear
				gear, _ := gears[cand]
				gear.neighbors = append(gear.neighbors, &num)
				gears[cand] = gear
			}
			return true, num.value
		}
	}

	return false, 0
}

func isSymbol(r rune) bool {
	symbols := regexp.MustCompile(`[^0-9.]`)
	return symbols.Match([]byte(string(r)))
}

func (g gear) ratio() int {
	if len(g.neighbors) == 2 {
		return g.neighbors[0].value * g.neighbors[1].value
	}
	return 0
}
