package day_13

import (
	"fmt"
	"strings"
)

type mirror [][]rune

func parse(input []byte) []mirror {
	var out []mirror
	for _, block := range strings.Split(string(input), "\n\n") {
		if block != "" {
			lines := strings.Split(block, "\n")
			var thisBlock mirror

			for _, line := range lines {
				if line != "" {
					thisBlock = append(thisBlock, []rune(line))
				}
			}

			out = append(out, thisBlock)
		}
	}

	return out
}

func value(m mirror) (int, bool, int) {
	// Check horizontally (i.e., axis is VERTICAL)
	for i := 0; i < len(m[0])-1; i++ {
		if checkVertical(m, i, 0) {
			return i + 1, true, i
		}
	}

	// Check vertically (i.e. axis is HORIZONTAL)
	for i := 0; i < len(m)-1; i++ {
		if checkHorizontal(m, i, 0) {
			return (i + 1) * 100, false, i
		}
	}
	return 0, false, 0
}

func checkVertical(m mirror, col, dist int) bool {
	left := col - dist
	right := (col + 1) + dist

	mirror := true
	for _, row := range m {
		if string(row) == "" {
			continue
		}

		mirror = mirror && (row[left] == row[right])
		if !mirror {
			return false
		}
	}
	if mirror {
		if left == 0 || right == len(m[0])-1 {
			return true
		} else {
			return checkVertical(m, col, dist+1)
		}
	}

	return false
}

func checkHorizontal(m mirror, row, dist int) bool {
	top := row - dist
	bottom := (row + 1) + dist

	mirror := true
	for col := 0; col < len(m[0]); col++ {
		mirror = mirror && (m[top][col] == m[bottom][col])
		if !mirror {
			return false
		}
	}
	if mirror {
		if top == 0 || bottom == len(m)-1 {
			return true
		} else {
			return checkHorizontal(m, row, dist+1)
		}
	}
	return false
}

func smudge(m mirror) int {
	_, origBool, origLine := value(m)
	fmt.Println(origBool, origLine)

	for i := 0; i < len(m[0])*(len(m)-1); i++ {
		if m[i/len(m[0])][i%len(m[0])] == '.' {
			m[i/len(m[0])][i%len(m[0])] = '#'
		} else {
			m[i/len(m[0])][i%len(m[0])] = '.'
		}

		for i := 0; i < len(m)-1; i++ {
			if checkHorizontal(m, i, 0) {
				val, newBool, newLine := value(m)
				if origBool != newBool || origLine != newLine {
					return val
				}
			}
		}

		for i := 0; i < len(m[0])-1; i++ {
			if checkVertical(m, i, 0) {
				_, newBool, newLine := value(m)
				fmt.Println(newBool, newLine)
				if origBool != newBool || origLine != newLine {
					return (i + 1) * 100
				}
			}
		}

		if m[i/len(m[0])][i%len(m[0])] == '.' {
			m[i/len(m[0])][i%len(m[0])] = '#'
		} else {
			m[i/len(m[0])][i%len(m[0])] = '.'
		}
	}
	return 0
}
