package helpers

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

/////////////////////
// Data structures //
/////////////////////

// A queue can be implemented by using a slice and appending items to the end.
// Dequeue returns the first item in the queue and the rest of the queue (or an empty slice)
func Dequeue[T any](q []T) (T, []T) {
	out := q[0]

	if len(q) == 1 {
		return out, make([]T, 0)
	} else {
		return out, q[1:]
	}
}

// When using map[T]struct{} as a set, this is a somewhat nicer way to check for membership
// Might be overkill (if _, present := map[key]; present)
func Member[T comparable](k T, m map[T]struct{}) bool {
	_, present := m[k]
	return present
}

///////////////////
// Input parsing //
///////////////////

// Returns all the integers found in a string as a slice. Integers do not need to be separated from non-integer runes
// If you know a line will only contain a single int, grab it by indexing [0] in the output
func ReGetInts(s string) []int {
	re := regexp.MustCompile(`[0-9]+`)
	matches := re.FindAllString(s, -1)

	ints := make([]int, 0)

	for _, match := range matches {
		val, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		ints = append(ints, val)
	}

	return ints
}

///////////
// Grids //
///////////

// Four ways to represent a grid
// - A single slice -> calculate the index based on the height, width, row and column
// - A 2D array -> pretty self-explanatory
// - A Map with coordinates as keys
// - A graph of nodes with neighbor pointers

// The first three options all have one issue in common: you have to check for edges whenever you interact with neighbors to avoid going out of bounds
// The last option tackles that problem at creation time, which is why I prefer it

// These implementations are a starting point, additional fields might be needed for individual puzzles. For example:
// - Whether a node has been visited or has otherwise been interacted with (bool)
// - If the location of the neighbors matters, replace the slice of pointers with four pointers: up, down, right, left

// Grids of integers

type inode struct {
	value     int
	neighbors []*inode
}

// Takes a string as input and returns a slice of inodes with populated neighbor pointers, as well as the width and height of the grid
// We assume the input string to consist of multiple lines, representing the rows of the grid
func IGridAsGraph(input string) ([]inode, int, int) {
	out := make([]inode, 0) // []*inode also an option, maybe if inode would have many more fields
	var width, height int

	lines := strings.Split(input, "\n")

	for row, l := range lines {
		if l != "" {
			for col, r := range l {
				val, _ := strconv.Atoi(string(r))
				out = append(out, inode{value: val})

				if col > width {
					width = col
				}
			}
		}
		if row > height {
			height = row
		}
	}

	// Catch the inevitable off-by-one error
	width += 1

	// Populate the neighbor pointers
	for i, n := range out {
		if i%width != 0 { // Check for left edge
			n.neighbors = append(n.neighbors, &out[i-1])
		}
		if i%width != width-1 && i != width*height { // Check for right edge
			n.neighbors = append(n.neighbors, &out[i+1])
		}
		if i >= width { // Check for top edge
			n.neighbors = append(n.neighbors, &out[i-width])
		}
		if i < (width*height)-width { // Check for bottom edge
			n.neighbors = append(n.neighbors, &out[i+width])
		}
	}

	return out, width, height
}

type rnode struct {
	value     rune
	neighbors []*rnode
}

// Takes a string as input and returns a slice of rnodes with populated neighbor pointers, as well as the width and height of the grid
// We assume the input string to consist of multiple lines, representing the rows of the grid
func RGridAsGraph(input string) ([]rnode, int, int) {
	out := make([]rnode, 0) // []*rnode also an option, maybe if rnode would have many more fields
	var width, height int

	lines := strings.Split(input, "\n")

	for row, l := range lines {
		if l != "" {
			for col, r := range l {
				out = append(out, rnode{value: r})

				if col > width {
					width = col
				}
			}
		}
		if row > height {
			height = row
		}
	}

	// Catch the inevitable off-by-one error
	width += 1

	// Populate the neighbor pointers
	for i, n := range out {
		if i%width != 0 { // Check for left edge
			n.neighbors = append(n.neighbors, &out[i-1])
		}
		if i%width != width-1 && i != width*height { // Check for right edge
			n.neighbors = append(n.neighbors, &out[i+1])
		}
		if i >= width { // Check for top edge
			n.neighbors = append(n.neighbors, &out[i-width])
		}
		if i < (width*height)-width { // Check for bottom edge
			n.neighbors = append(n.neighbors, &out[i+width])
		}
	}

	return out, width, height
}

type point struct {
	x int
	y int
}

// Returns the manhattan distance between two points
func ManhattanDistance(a, b point) int {
	return slices.Max([]int{a.x, b.x}) - slices.Min([]int{a.x, b.x}) + slices.Max([]int{a.y, b.y}) - slices.Min([]int{a.y, b.y})
}
