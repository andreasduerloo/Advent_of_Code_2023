package day_05

import (
	"fmt"
	"os"
	"slices"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/05.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	seeds, farmMaps := parse(input)
	seedPairs := pair(seeds)

	for _, fmap := range farmMaps { // Do this with a map function?
		var newVals []int
		for _, seed := range seeds {
			newVals = append(newVals, transform(seed, fmap))
		}
		seeds = newVals
	}

	first := slices.Min(seeds)

	/* Second star
	I brute forced the second star (using multithreading - thank you goroutines).
	It works and it computes in a reasonable time (about two minutes), but it's obviously not the optimal solution.
	I think I know what I need to do: work from the back and determine what input range produces the right output range at every step.
	This is pretty complicated: there's the overlaps between the source and destination ranges to figure out, and the fact that numbers can bypass a map entirely.
	I'll come back to this problem later. In the meantime, here's the brute force solution.
	*/

	lowestChan := make(chan int, 10)
	var lowestVals []int

	for i := 0; i < 10; i++ {
		go func(s seedPair, f []farmMap, lowChan chan<- int) {
			fmt.Println("I have started planting seeds!")
			lowest := first
			for i := s.low; i < s.low+s.count; i++ {
				seed := i
				for _, fmap := range f {
					seed = transform(seed, fmap)
				}
				if seed < lowest {
					lowest = seed
				}
			}
			lowChan <- lowest
		}(seedPairs[i], farmMaps, lowestChan)
	}

	for i := 0; i < 10; i++ {
		received := <-lowestChan
		fmt.Println("I received a lowest value!")
		lowestVals = append(lowestVals, received)
	}

	second := slices.Min(lowestVals)

	return first, second
}
