package day_13

import (
	"strings"
)

func parse(input []byte) [][]string {
	var out [][]string
	for _, block := range strings.Split(string(input), "\n\n") {
		if block != "" {
			thisBlock := strings.Split(block, "\n")
			out = append(out, thisBlock)
		}
	}

	return out
}

func value(block []string) int {
	// Check horizontally (i.e., axis is VERTICAL)
	for i := 0; i < len(block[0])-1; i++ {
		if checkVertical(block, i, 0) {
			return i + 1
		}
	}

	// Check vertically (i.e. axis is HORIZONTAL)
	for i := 0; i < len(block)-1; i++ {
		if checkHorizontal(block, i, 0) {
			return (i + 1) * 100
		}
	}
	return 0
}

func checkVertical(block []string, col, dist int) bool {
	left := col - dist
	right := (col + 1) + dist

	mirror := true
	for _, row := range block {
		if row == "" {
			continue
		}

		mirror = mirror && (row[left] == row[right])
		if !mirror {
			return false
		}
	}
	if mirror {
		if left == 0 || right == len(block[0])-1 {
			return true
		} else {
			return checkVertical(block, col, dist+1)
		}
	}

	return false
}

func checkHorizontal(block []string, row, dist int) bool {
	top := row - dist
	bottom := (row + 1) + dist

	mirror := true
	for col := 0; col < len(block[0]); col++ {
		mirror = mirror && (block[top][col] == block[bottom][col])
		if !mirror {
			return false
		}
	}
	if mirror {
		if top == 0 || bottom == len(block)-1 {
			return true
		} else {
			return checkHorizontal(block, row, dist+1)
		}
	}
	return false
}
