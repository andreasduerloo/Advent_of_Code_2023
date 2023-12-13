package day_12

import (
	"advent2023/helpers"
	"strings"
)

type row struct {
	springs string
	damaged []int
}

func parse(input []byte) []row {
	lines := strings.Split(string(input), "\n")
	var out []row

	for _, line := range lines {
		if line == "" {
			continue
		}

		out = append(out, row{springs: strings.Split(line, " ")[0], damaged: helpers.ReGetInts(line)})
	}

	return out
}

/*
func (r row) possibilites() int {
	// The total possibilites is the product of possibilities for the sub-problems
	// A sub-problem is any connected combination of #'s and ?'s with one or more numbers
	out := 1
	subRegex := regexp.MustCompile(`[·]+`)

	subProblems := subRegex.FindAllString(r.springs, -1)

	// If longer than next number -> take it
	// Try to add next number as well
}
*/
