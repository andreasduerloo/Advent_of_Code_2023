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
	neighbors []*number // Pointers to the rescue
}

///////////////////
// Input parsing //
///////////////////

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

////////////////
// First star //
////////////////

func checkNumber(num number, schem grid, gears map[int]gear) (bool, int) {
	// Check if the number is adjacent to a symbol

	// Look at the row above
	if num.begin >= schem.width {

		// Can we look up and left?
		if num.begin%schem.width != 0 {
			if isSymbol(schem.data[num.begin-(schem.width+1)]) {
				if schem.data[num.begin-(schem.width+1)] == '*' {
					// This is a gear
					gear, _ := gears[num.begin-(schem.width+1)]
					gear.neighbors = append(gear.neighbors, &num)
					gears[num.begin-(schem.width+1)] = gear
				}
				return true, num.value
			}
		}

		// Can we look up and right?
		if num.end%schem.width != schem.width {
			if isSymbol(schem.data[num.end-schem.width]) {
				if schem.data[num.end-schem.width] == '*' {
					gear, _ := gears[num.end-schem.width]
					gear.neighbors = append(gear.neighbors, &num)
					gears[num.end-schem.width] = gear
				}
				return true, num.value
			}
		}

		// Look up
		for i := num.begin - schem.width; i < num.end-schem.width; i++ {
			if isSymbol(schem.data[i]) {
				if schem.data[i] == '*' {
					gear, _ := gears[i]
					gear.neighbors = append(gear.neighbors, &num)
					gears[i] = gear
				}
				return true, num.value
			}
		}
	}

	// Look at the same row
	if num.begin > 0 {
		if isSymbol(schem.data[num.begin-1]) {
			if schem.data[num.begin-1] == '*' {
				gear, _ := gears[num.begin-1]
				gear.neighbors = append(gear.neighbors, &num)
				gears[num.begin-1] = gear
			}
			return true, num.value
		}
	}

	if num.end%schem.width != schem.width-1 {
		if isSymbol(schem.data[num.end]) {
			if schem.data[num.end] == '*' {
				gear, _ := gears[num.end]
				gear.neighbors = append(gear.neighbors, &num)
				gears[num.end] = gear
			}
			return true, num.value
		}
	}

	// Look at the row below
	if num.begin < (schem.width*schem.height)-schem.width {

		// Can we look down and left?
		if num.begin%schem.width != 0 {
			if isSymbol(schem.data[num.begin+(schem.width-1)]) {
				if schem.data[num.begin+(schem.width-1)] == '*' {
					gear, _ := gears[num.begin+(schem.width-1)]
					gear.neighbors = append(gear.neighbors, &num)
					gears[num.begin+(schem.width-1)] = gear
				}
				return true, num.value
			}
		}

		// Can we look down and right?
		if num.end%schem.width != schem.width {
			if isSymbol(schem.data[num.end+schem.width]) {
				if schem.data[num.end+schem.width] == '*' {
					gear, _ := gears[num.end+schem.width]
					gear.neighbors = append(gear.neighbors, &num)
					gears[num.end+schem.width] = gear
				}
				return true, num.value
			}
		}

		// Look down
		for i := num.begin + schem.width; i < num.end+schem.width; i++ {
			if isSymbol(schem.data[i]) {
				if schem.data[i] == '*' {
					gear, _ := gears[i]
					gear.neighbors = append(gear.neighbors, &num)
					gears[i] = gear
				}
				return true, num.value
			}
		}
	}

	return false, 0
}

func isSymbol(r rune) bool {
	symbols := regexp.MustCompile(`[^0-9.]`)
	return symbols.Match([]byte(string(r)))
}

/////////////////
// Second star //
/////////////////

func gearRatio(g gear) int {
	if len(g.neighbors) == 2 {
		return g.neighbors[0].value * g.neighbors[1].value
	}
	return 0
}
