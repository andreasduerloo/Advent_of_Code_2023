package day_14

import (
	"fmt"
	"hash/fnv"
	"strings"
)

const (
	NORTH = iota
	WEST
	SOUTH
	EAST
)

func parse(input []byte) [][]rune {
	var out [][]rune

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		out = append(out, []rune(line))
	}

	return out
}

func roll(direction int) func(*[][]rune) {
	funcs := []func(*[][]rune){
		rollNorth,
		rollWest,
		rollSouth,
		rollEast,
	}

	return funcs[direction%4]
}

func cycle(times, start int, rocks *[][]rune) int {
	hashes := make([]uint64, 0)
	hashes = append(hashes, summarize(*rocks))

	detected := 0
	hare := 0

	for i := start; i < times; i++ {
		roll(i)(rocks)
		hashes = append(hashes, summarize(*rocks))

		cycle, hareNow := detectCycle(hashes)
		if cycle {
			fmt.Println("Found a cycle at i =", i, "collecting some more values...")
			detected = i
			hare = hareNow
			break
		}
	}

	// Add another i cycles to give the algorithm enough time to find the start and length of the cycle
	for i := detected + 1; i < detected*2; i++ {
		roll(i)(rocks)
		hashes = append(hashes, summarize(*rocks))
	}

	// Now find the start and length of the cycle
	start, length := cycleStartAndLength(hashes, hare)
	fmt.Println("The cycle starts at", start, "and has a length of", length)

	// Now find the representative value for 'times'
	return ((times - start) % length) + start
}

func cycle2(times, start int, rocks *[][]rune) {
	for i := start; i < times; i++ {
		roll(i)(rocks)
		fmt.Println(calculateLoad(*rocks), summarize(*rocks), i)
	}
}

func rollNorth(rocks *[][]rune) {
	for row := 0; row < len(*rocks); row++ {
		for col, r := range (*rocks)[row] {
			if r == 'O' { // Found a ball
				check := row - 1 // Start looking north

				for check >= 0 && (*rocks)[check][col] == '.' { // Move the ball up until blocked or at the top
					check-- // Keep looking upwards for an open space
				} // The ball will stop at check+1 (just below a block/ball or the top edge)

				if check+1 != row { // Only move if the ball is not already at the highest valid position
					(*rocks)[row][col] = '.'     // Clear the original position
					(*rocks)[check+1][col] = 'O' // Move the ball up to the final valid position
				}
			}
		}
	}
}

func rollWest(rocks *[][]rune) {
	for col := 0; col < len((*rocks)[0]); col++ {
		for row := 0; row < len(*rocks); row++ {
			if (*rocks)[row][col] == 'O' {
				check := col - 1

				for check >= 0 && (*rocks)[row][check] == '.' {
					check--
				}

				if check-1 != col {
					(*rocks)[row][col] = '.'
					(*rocks)[row][check+1] = 'O'
				}
			}
		}
	}
}

func rollSouth(rocks *[][]rune) {
	for row := len(*rocks) - 1; row >= 0; row-- {
		for col, r := range (*rocks)[row] {
			if r == 'O' {
				check := row + 1

				for check < len(*rocks) && (*rocks)[check][col] == '.' {
					check++
				}

				if check-1 != row {
					(*rocks)[row][col] = '.'
					(*rocks)[check-1][col] = 'O'
				}
			}
		}
	}
}

func rollEast(rocks *[][]rune) {
	for col := len((*rocks)[0]) - 1; col >= 0; col-- {
		for row := 0; row < len(*rocks); row++ {
			if (*rocks)[row][col] == 'O' {
				check := col + 1

				for check < len((*rocks)[0]) && (*rocks)[row][check] == '.' {
					check++
				}

				if check+1 != col {
					(*rocks)[row][col] = '.'
					(*rocks)[row][check-1] = 'O'
				}
			}
		}
	}
}

func calculateLoad(rocks [][]rune) int {
	var out int
	max := len(rocks)

	for row, line := range rocks {
		for _, rock := range line {
			if rock == 'O' {
				out += (max - row)
			}
		}
	}

	return out
}

func summarize(rocks [][]rune) uint64 {
	hasher := fnv.New64a()

	for _, row := range rocks {
		for _, r := range row {
			hasher.Write([]byte(string(r)))
		}
	}

	return hasher.Sum64()
}

func detectCycle(hashes []uint64) (bool, int) {
	max := len(hashes)

	tortoise, hare := 0, 0

	for {
		if tortoise+1 >= max || hare+2 >= max {
			return false, 0
		}

		tortoise++
		hare += 2

		if hashes[tortoise] == hashes[hare] {
			return true, hare
		}
	}
}

func cycleStartAndLength(hashes []uint64, hare int) (int, int) {
	start, length := 0, 0

	// Find the start
	tortoise := 0
	for hashes[tortoise] != hashes[hare] {
		tortoise++
		hare++
	}

	start = tortoise

	// Find the length
	tortoise++
	compareVal := hashes[hare]
	for hashes[tortoise] != compareVal {
		tortoise++
	}

	length = tortoise - start

	return start, length
}

func rollAndCount(rocks [][]rune) int {
	var out int
	max := len(rocks)

	// Column by column
	for col := 0; col < len(rocks[0]); col++ {
		lastFree := 0

		for row := 0; row < len(rocks); row++ {
			switch rocks[row][col] {
			case 'O':
				out += (max - (lastFree))
				lastFree = (lastFree + 1)
			case '#':
				lastFree = (row + 1)
			}
		}
	}

	return out
}
