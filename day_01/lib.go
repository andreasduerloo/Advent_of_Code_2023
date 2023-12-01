package day_01

import (
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

func findIndex(s string, list []string, first bool) (int, error) {
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
