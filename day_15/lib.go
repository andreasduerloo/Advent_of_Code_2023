package day_15

import (
	"strings"
)

func hash(input []byte) int {
	var out int

	steps := strings.Split(string(input), ",")

	for _, step := range steps {
		var hash int

		for _, r := range step {
			hash += int(r)
			hash *= 17
			hash = hash % 256
		}

		out += hash
	}

	return out
}
