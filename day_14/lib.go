package day_14

import (
	"strings"
)

func parse(input []byte) [][]rune {
	var out [][]rune

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		out = append(out, []rune(line))
	}

	return out
}

func rollNorth(rocks *[][]rune) {
	for row, line := range *rocks {
		for col, r := range line {
			if r == 'O' { // Look north as far as we can go
				check := row - 1
				for check >= 0 {
					switch (*rocks)[check][col] {
					case '.':
						if check == 0 {
							(*rocks)[row][col] = '.'
							(*rocks)[check][col] = 'O'
						} else {
							check--
						}
					case '#':
						if check+1 == row {
							check = -1
						} else {
							(*rocks)[row][col] = '.'
							(*rocks)[check+1][col] = 'O'
							check = -1
						}
					case 'O':
						if check+1 == row {
							check = -1
						} else {
							(*rocks)[row][col] = '.'
							(*rocks)[check+1][col] = 'O'
							check = -1
						}
					}
				}
			}
		}
	}
}

func calculateLoad(rocks [][]rune) int {
	var out int
	max := len(rocks)

	for row, line := range rocks {
		for _, rock := range line {
			if rock == 'O' {
				out += (max - row)
			}
		}
	}

	return out
}
