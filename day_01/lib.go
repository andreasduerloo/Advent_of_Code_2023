package day_01

import (
	"advent2023/helpers"
	"regexp"
	"strconv"
	"strings"
)

func parse(input []byte, second bool) int {
	lines := strings.Split(string(input), "\n")
	digits := make([][]int, 0)

	for _, line := range lines {
		if line != "" {
			digits = append(digits, findDigits(line, second))
		}
	}

	values := mapInts(digits, func(d []int) int { return 10*d[0] + d[len(d)-1] }) // Higher-order function, anonymous function, woo!
	return helpers.SumSlice(values)
}

func findDigits(s string, second bool) []int {
	out := make([]int, 0)

	for i, r := range s {
		// Regular digits
		val, err := strconv.Atoi(string(r))
		if err == nil {
			out = append(out, val)
			continue
		}

		if second {
			digits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

			// Three-letter digits
			if i+3 <= len(s) {
				nextThree := s[i : i+3]
				if val, present := digits[nextThree]; present {
					out = append(out, val)
					continue
				}
			}

			// Four-letter digits
			if i+4 <= len(s) {
				nextFour := s[i : i+4]
				if val, present := digits[nextFour]; present {
					out = append(out, val)
					continue
				}
			}

			// Five-letter digits
			if i+5 <= len(s) {
				nextFive := s[i : i+5]
				if val, present := digits[nextFive]; present {
					out = append(out, val)
				}
			}
		}
	}

	return out
}

func mapInts(slice [][]int, mapping func([]int) int) []int {
	var result []int
	for _, v := range slice {
		result = append(result, mapping(v))
	}
	return result
}

// Here is my original solution. It's good for the first star, but horrible for the second one. I rewrote my solution with the advantage of hindsight.

func parse1(input []byte) []int {
	lines := strings.Split(string(input), "\n")
	out := make([]int, 0)

	for _, line := range lines {
		if line != "" {
			reDigits := regexp.MustCompile(`\d{1}`)
			matches := reDigits.FindAllString(line, -1)

			first, _ := strconv.Atoi(matches[0])
			second, _ := strconv.Atoi(matches[len(matches)-1])

			out = append(out, (10*first + second))
		}
	}

	return out
}
