package day_01

import (
	"advent2023/helpers"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parse(input []byte) []int {
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

// What makes this problem difficult is the fact that regexes (in go) don't find overlapping matches.
// In other words: twone will match 'two' and not 'one'. That's fine if 'twone' is the first digit, but it's problematic
// if the 'one' in 'twone' happens to be the last digit.
// I wonder if a more naive solution isn't better - see below.
func parse2(input []byte) []int {
	lines := strings.Split(string(input), "\n")
	out := make([]int, 0)

	for _, line := range lines {
		if line != "" {
			reDigits := regexp.MustCompile(`\d{1}|zero|oneight|one|twone|two|threeight|three|four|fiveight|five|six|sevenine|seven|eightwo|eighthree|eight|nineight|nine`)
			matches := reDigits.FindAllString(line, -1)

			digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			var outNum int

			if len(matches[0]) == 1 {
				first, _ := strconv.Atoi(matches[0])
				outNum += 10 * first
			} else {
				first, _ := findIndex(matches[0], digits, true)
				outNum += 10 * first
			}

			if len(matches[len(matches)-1]) == 1 {
				second, _ := strconv.Atoi(matches[len(matches)-1])
				outNum += second
			} else {
				second, _ := findIndex(matches[len(matches)-1], digits, false)
				outNum += second
			}

			out = append(out, outNum)
		}
	}

	return out
}

func findIndex(s string, list []string, first bool) (int, error) { // This looks horrible, could be replaced with a hashmap
	for i := 0; i < len(list); i++ {
		if list[i] == s {
			return i, nil
		} else if s == "oneight" {
			if first {
				return 1, nil
			} else {
				return 8, nil
			}
		} else if s == "twone" {
			if first {
				return 2, nil
			} else {
				return 1, nil
			}
		} else if s == "threeight" {
			if first {
				return 3, nil
			} else {
				return 8, nil
			}
		} else if s == "fiveight" {
			if first {
				return 5, nil
			} else {
				return 8, nil
			}
		} else if s == "sevenine" {
			if first {
				return 7, nil
			} else {
				return 9, nil
			}
		} else if s == "eightwo" {
			if first {
				return 8, nil
			} else {
				return 2, nil
			}
		} else if s == "eighthree" {
			if first {
				return 8, nil
			} else {
				return 3, nil
			}
		} else if s == "nineight" {
			if first {
				return 9, nil
			} else {
				return 8, nil
			}
		}
	}

	return 0, errors.New("Not found!")
}

// Here is the naive solution, where I just iterate throug the entire string
func findDigits(s string) []int {
	out := make([]int, 0)

	for i, r := range s {
		// Regular digits
		val, err := strconv.Atoi(string(r))
		if err == nil {
			out = append(out, val)
			continue
		}

		// Three-letter digits
		if i+3 <= len(s) {
			nextThree := s[i : i+3]
			switch nextThree {
			case "one":
				out = append(out, 1)
				continue
			case "two":
				out = append(out, 2)
				continue
			case "six":
				out = append(out, 6)
				continue
			}
		}

		// Four-letter digits
		if i+4 <= len(s) {
			nextFour := s[i : i+4]
			switch nextFour {
			case "four":
				out = append(out, 4)
				continue
			case "five":
				out = append(out, 5)
				continue
			case "nine":
				out = append(out, 9)
				continue
			}
		}

		// Five-letter digits
		if i+5 <= len(s) {
			nextFive := s[i : i+5]
			switch nextFive {
			case "three":
				out = append(out, 3)
				continue
			case "seven":
				out = append(out, 7)
				continue
			case "eight":
				out = append(out, 8)
				continue
			}
		}
	}

	return out
}

func naive(input []byte) int {
	lines := strings.Split(string(input), "\n")
	digits := make([][]int, 0)

	for _, line := range lines {
		if line != "" {
			digits = append(digits, findDigits(line))
		}
	}

	values := mapInts(digits, func(d []int) int { return 10*d[0] + d[len(d)-1] }) // Higher-order function, anonymous function, woo!
	return helpers.SumSlice(values)
}

func mapInts(slice [][]int, mapping func([]int) int) []int {
	var result []int
	for _, v := range slice {
		result = append(result, mapping(v))
	}
	return result
}
