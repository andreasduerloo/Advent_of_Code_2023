package day_10

import (
	"advent2023/helpers"
	"slices"
	"strings"
)

type pipe struct {
	value     rune
	connected []*pipe
	neighbors []*pipe
	visited   bool
	set       int
}

func parse(input []byte) (*pipe, []*pipe) {
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

	// Set the neighbor pointers (second star) - WRAP AROUND THE EDGES
	for i, p := range pipes {
		if i >= width {
			nb := p.neighbors
			nb = append(nb, pipes[i-width])
			p.neighbors = nb
		} else {
			nb := p.neighbors
			nb = append(nb, pipes[len(pipes)-width-i])
			p.neighbors = nb
		}

		if i < width*height-width {
			nb := p.neighbors
			nb = append(nb, pipes[i+width])
			p.neighbors = nb
		} else {
			nb := p.neighbors
			nb = append(nb, pipes[i-(width*height-width)])
			p.neighbors = nb
		}

		if i%width > 0 {
			nb := p.neighbors
			nb = append(nb, pipes[i-1])
			p.neighbors = nb
		} else {
			nb := p.neighbors
			nb = append(nb, pipes[i+(width-1)])
			p.neighbors = nb
		}

		if i%width != width-1 {
			nb := p.neighbors
			nb = append(nb, pipes[i+1])
			p.neighbors = nb
		} else {
			nb := p.neighbors
			nb = append(nb, pipes[i-(width)+1])
			p.neighbors = nb
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
	if startIndex-width > 0 && slices.Contains([]rune{'|', '7', 'F'}, pipes[startIndex-width].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex-width])
		start.connected = nb
	}

	if startIndex+width < width*height-width && slices.Contains([]rune{'|', 'J', 'L'}, pipes[startIndex+width].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex+width])
		start.connected = nb
	}

	if (startIndex-1)%width != 0 && slices.Contains([]rune{'-', 'F', 'L'}, pipes[startIndex-1].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex-1])
		start.connected = nb
	}

	if (startIndex+1)%width != width-1 && slices.Contains([]rune{'F', 'J', '7'}, pipes[startIndex+1].value) {
		nb := start.connected
		nb = append(nb, pipes[startIndex+1])
		start.connected = nb
	}

	return start, pipes
}

func steps(start *pipe) int {
	var out int

	// Pick one of hte neighbors
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

	return out
}

// Second star
// Breadth-first search from the top left corner re-using the 'visited' field. Anything we can't reach is enclosed by the loop
// So second = total - (staps + bfs)
func bfs(start *pipe) int {
	var queue []*pipe
	var out int

	current := start
	current.visited = true
	out += 1

	for _, nb := range current.neighbors {
		if !nb.visited {
			nb.visited = true
			queue = append(queue, nb)
		}
	}

	for len(queue) > 0 {
		current, queue = helpers.Dequeue(queue)
		out += 1

		for _, nb := range current.neighbors {
			if !nb.visited {
				nb.visited = true
				queue = append(queue, nb)
			}
		}
	}

	return out
}
