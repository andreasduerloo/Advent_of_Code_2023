package day_10

import (
	"slices"
	"strings"
)

type pipe struct {
	value     rune
	connected []*pipe
	visited   bool
}

func parse(input []byte) (*pipe, []*pipe, int) {
	lines := strings.Split(string(input), "\n")
	var pipes []*pipe
	var start *pipe
	var startIndex int

	// Set the width and height variables - we expect 140 x 140
	width := len(lines[0])
	height := len(lines) - 1 // Trim the empty line

	for _, line := range lines {
		for _, r := range line {
			pipes = append(pipes, &pipe{value: r})
		}
	}

	// Set the connected pointers
	for i, p := range pipes { // Refortmat? We keep recombining the same steps
		switch p.value {
		case '|':
			if i >= width {
				nb := p.connected
				nb = append(nb, pipes[i-width])
				p.connected = nb
			}
			if i < width*height-width {
				nb := p.connected
				nb = append(nb, pipes[i+width])
				p.connected = nb
			}
		case '-':
			if i%width > 0 {
				nb := p.connected
				nb = append(nb, pipes[i-1])
				p.connected = nb
			}
			if i%width != width-1 {
				nb := p.connected
				nb = append(nb, pipes[i+1])
				p.connected = nb
			}
		case 'L':
			if i >= width {
				nb := p.connected
				nb = append(nb, pipes[i-width])
				p.connected = nb
			}
			if i%width != width-1 {
				nb := p.connected
				nb = append(nb, pipes[i+1])
				p.connected = nb
			}
		case 'J':
			if i >= width {
				nb := p.connected
				nb = append(nb, pipes[i-width])
				p.connected = nb
			}
			if i%width > 0 {
				nb := p.connected
				nb = append(nb, pipes[i-1])
				p.connected = nb
			}
		case '7':
			if i < width*height-width {
				nb := p.connected
				nb = append(nb, pipes[i+width])
				p.connected = nb
			}
			if i%width > 0 {
				nb := p.connected
				nb = append(nb, pipes[i-1])
				p.connected = nb
			}
		case 'F':
			if i < width*height-width {
				nb := p.connected
				nb = append(nb, pipes[i+width])
				p.connected = nb
			}
			if i%width != width-1 {
				nb := p.connected
				nb = append(nb, pipes[i+1])
				p.connected = nb
			}
		case 'S':
			start = p
			startIndex = i
		}
	}

	// Set the neighbor pointers for start
	possible := make(map[rune]int)

	if startIndex-width > 0 && slices.Contains([]rune{'|', '7', 'F'}, pipes[startIndex-width].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex-width])
		start.connected = nb

		for _, r := range []rune{'|', 'J', 'L'} {
			if _, present := possible[r]; present {
				possible[r] += 1
			} else {
				possible[r] = 1
			}
		}
	}

	if startIndex+width < width*height-width && slices.Contains([]rune{'|', 'J', 'L'}, pipes[startIndex+width].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex+width])
		start.connected = nb

		for _, r := range []rune{'|', '7', 'F'} {
			if _, present := possible[r]; present {
				possible[r] += 1
			} else {
				possible[r] = 1
			}
		}
	}

	if (startIndex-1)%width != 0 && slices.Contains([]rune{'-', 'F', 'L'}, pipes[startIndex-1].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex-1])
		start.connected = nb

		for _, r := range []rune{'-', '7', 'J'} {
			if _, present := possible[r]; present {
				possible[r] += 1
			} else {
				possible[r] = 1
			}
		}
	}

	if (startIndex+1)%width != width-1 && slices.Contains([]rune{'F', 'J', '7'}, pipes[startIndex+1].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex+1])
		start.connected = nb

		for _, r := range []rune{'-', 'L', 'F'} {
			if _, present := possible[r]; present {
				possible[r] += 1
			} else {
				possible[r] = 1
			}
		}
	}

	for r, count := range possible {
		if count == 2 {
			start.value = r
		}
	}

	return start, pipes, width
}

func steps(start *pipe) int {
	var out int

	// Pick one of the neighbors
	current := start
	current.visited = true
	out += 1
	current = current.connected[0]

	// Go on an adventure
	for !(current.connected[0].visited && current.connected[1].visited) { // One of the neighbors has not been visited
		current.visited = true
		out += 1
		if !current.connected[0].visited {
			current = current.connected[0]
		} else {
			current = current.connected[1]
		}
	}
	current.visited = true

	return out
}

// Tried BFS, didn't work because you can squeeze between the pipes
// Polygon technique: keep track of whether we are inside or outside the shape defined by the loop
func scan(grid []*pipe, width int) int {
	var out int
	var inside bool
	var entryChar rune

	// Kind of regretting working with a single slice now
	for i, p := range grid {
		if i%width == 0 {
			inside = false
		}

		if p.visited && p.value != '-' { // We're on the loop (ignoring horizontals)
			switch p.value {
			case '|':
				inside = !inside // Flip inside, whatever it was
			case 'F': // We're starting a horizontal line
				entryChar = p.value
				inside = !inside
			case 'L': // We're starting a horizontal line
				entryChar = p.value
				inside = !inside
			case 'J': // We're ending a horizontal line
				if entryChar != 'F' { // Flip
					inside = !inside
				}
			case '7': // We're ending a horizontal line
				if entryChar != 'L' { // Flip
					inside = !inside
				}
			}
		} else if !p.visited { // We're not on the loop
			if inside {
				out += 1
			}
		}
	}
	return out
}
