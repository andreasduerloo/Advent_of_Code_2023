package day_06

import (
	"advent2023/helpers"
	"math"
	"strconv"
	"strings"
)

func parse(input []byte) ([]int, []int) {
	lines := strings.Split(string(input), "\n")

	return helpers.ReGetInts(lines[0]), helpers.ReGetInts(lines[1])
}

// distance = (wait time) * (total time - wait time)
//             | This is the speed       | This is how long you go at that speed

// We need: record < (recordTime * waitTime) - waitTime²
// <=> 0 < -waitTime² + (recordTime * waitTime) - record

// With x = waitTime, these are are quadratic equations of the format -x²+(time * x) - distance. Find the zeroes for the equation, the range in between those is where we beat the record.
// For the zeroes you can either iterate until you find a sign change (lame), or use the quadratic formula: (ab +/- sqrt(b² - 4ac))/2a.

// Returns the number of integers for which we beat the record
func winningRange(time, distance int) int {
	a := float64(-1)
	b := float64(time)
	c := float64(-distance)

	left := ((-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a))
	right := ((-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a))

	return int(math.Floor(right) - math.Floor(left+1) + 1) // Add one to include both ends
}

func concatNum(numbers []int) int {
	var concatString string

	for _, number := range numbers {
		concatString += strconv.Itoa(number)
	}

	num, _ := strconv.Atoi(concatString)
	return num
}
