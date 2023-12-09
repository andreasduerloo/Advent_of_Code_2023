package day_09

import (
	"advent2023/helpers"
	"strings"
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

	for current[0] != 0 {
		current = differences(line)
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

	for i := 0; i < len(line)-1; i++ {
		out = append(out, (line[i+1] - line[i]))
	}

	return out
}
