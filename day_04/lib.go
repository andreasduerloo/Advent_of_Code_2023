package day_04

import (
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id      int
	numbers []int
	winning map[int]struct{}
}

func parse(input []byte) []game {
	var out []game

	lines := strings.Split(string(input), "\n")
	re := regexp.MustCompile(`[0-9|]+`)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		var id int
		var numbers []int
		winning := make(map[int]struct{})

		var separator bool
		for i, match := range matches {

			if i == 0 {
				id, _ = strconv.Atoi(match)
				continue
			}

			if !separator { // Before line
				if match == "|" {
					separator = true
				} else {
					val, _ := strconv.Atoi(match)
					numbers = append(numbers, val)
				}
			} else {
				val, _ := strconv.Atoi(match)
				winning[val] = struct{}{}
			}
		}
		out = append(out, game{
			id:      id,
			numbers: numbers,
			winning: winning,
		})
	}
	return out
}

func (g game) score() int {
	var score int

	for _, num := range g.numbers {
		if _, present := g.winning[num]; present {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}
