package day_09

import (
	"advent2023/helpers"
	"strings"

	"github.com/andreasduerloo/slicetools"
)

func parse(input []byte) [][]int {
	lines := strings.Split(string(input), "\n")
	var out [][]int

	for _, line := range lines {
		if line != "" {
			out = append(out, helpers.ReGetInts(line))
		}
	}

	return out
}

func nextValue(line []int) int {
	var allDifferences [][]int
	current := line

	for !slicetools.AllSlice(current, func(num int) bool { return num == 0 }) {
		current = differences(current)
		allDifferences = append(allDifferences, current)
	}

	out := line[len(line)-1]

	for _, diff := range allDifferences {
		out += diff[len(diff)-1]
	}

	return out
}

// Feels like we're deriving
func differences(line []int) []int {
	var out []int

	/*
		if len(line) == 1 {
			return []int{0}
		}
	*/

	for i := 0; i < len(line)-1; i++ {
		out = append(out, (line[i+1] - line[i]))
	}

	return out
}

// Second star
func previousValue(line []int) int {
	var allDifferences [][]int
	current := line

	for !slicetools.AllSlice(current, func(num int) bool { return num == 0 }) {
		current = differences(current)
		allDifferences = append(allDifferences, current)
	}

	return line[0] - backtrack(allDifferences)
}

func backtrack(differences [][]int) int { // Recursive
	if len(differences) == 2 {
		return differences[0][0] - differences[1][0]
	}
	return differences[0][0] - backtrack(differences[1:])
}
