package day_08

import (
	"advent2023/helpers"
	"regexp"
	"strings"
)

type node struct {
	name  string
	left  *node
	right *node
}

func parse(input []byte) (string, map[string]*node) {
	lines := strings.Split(string(input), "\n")

	instructions := lines[0]
	reNodes := regexp.MustCompile(`[A-Z]{3}`)

	// Create all the nodes (without pointers) and store them in a map
	nodes := make(map[string]*node)

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		matches := reNodes.FindAllString(line, -1)
		nodes[matches[0]] = &node{name: matches[0]}
	}

	// Populate the neighbor pointers
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		matches := reNodes.FindAllString(line, -1)
		node := nodes[matches[0]]
		node.left = nodes[matches[1]]
		node.right = nodes[matches[2]]
	}

	return instructions, nodes
}

func steps(start, end *node, instructions string) int {
	var steps int
	current := start

	instructionQueue := []rune(instructions)
	var instruction rune

	for current.name != end.name {
		instruction, instructionQueue = helpers.Dequeue(instructionQueue)

		switch instruction {
		case 'L':
			current = current.left
		case 'R':
			current = current.right
		}
		steps += 1
		instructionQueue = append(instructionQueue, instruction) // Put the instruction back at the end of the queue
	}

	return steps
}

// Second star
func startingPositions(nodes map[string]*node) []*node {
	reStart := regexp.MustCompile(`[A-Z]{2}A`)
	var out []*node

	for name, node := range nodes {
		if reStart.Match([]byte(name)) {
			out = append(out, node)
		}
	}

	return out
}

func stepsUntilZ(start *node, instructions string) int {
	var steps int
	current := start
	stopRe := regexp.MustCompile(`[A-Z]{2}Z`)

	instructionQueue := []rune(instructions)
	var instruction rune

	for !stopRe.Match([]byte(current.name)) {
		instruction, instructionQueue = helpers.Dequeue(instructionQueue)

		switch instruction {
		case 'L':
			current = current.left
		case 'R':
			current = current.right
		}
		steps += 1
		instructionQueue = append(instructionQueue, instruction) // Put the instruction back at the end of the queue
	}

	return steps
}

func greatestCommonDivisor(i, j int) int { // Euclid's algorithm
	if j == 0 {
		return i
	}

	if j > i {
		return greatestCommonDivisor(j, i)
	}

	return greatestCommonDivisor(j, i%j)
}

func leastCommonMultiple(i, j int) int {
	return (i * j) / greatestCommonDivisor(i, j)
}

func findLCM(nums []int) int {
	lcm := leastCommonMultiple(nums[0], nums[1])

	for i := 0; i < len(nums); i++ {
		lcm = leastCommonMultiple(lcm, nums[i])
	}

	return lcm
}

// Chinese remainder theorem
