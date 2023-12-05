package day_05

import (
	"advent2023/helpers"
	"strings"
)

type farmMap struct {
	name string
	maps []*rangeDef
}

type rangeDef struct {
	dest   int
	source int
	len    int
}

type seedPair struct {
	low   int
	count int
}

func parse(input []byte) ([]int, []farmMap) {
	blocks := strings.Split(string(input), "\n\n")

	seeds := helpers.ReGetInts(blocks[0])
	var allMaps []farmMap

	for _, block := range blocks[1:] {
		if block == "" {
			continue
		}

		lines := strings.Split(block, "\n")
		name := lines[0]
		var ranges []*rangeDef

		for _, rangeLine := range lines[1:] {
			if rangeLine != "" {
				lineInts := helpers.ReGetInts(rangeLine)
				ranges = append(ranges, &rangeDef{dest: lineInts[0], source: lineInts[1], len: lineInts[2]})
			}
		}
		allMaps = append(allMaps, farmMap{name: name, maps: ranges})
	}

	return seeds, allMaps
}

func transform(seed int, transformation farmMap) int {
	// Check whether the value is in any of the ranges
	for _, mapRange := range transformation.maps {
		if seed >= mapRange.source && seed < mapRange.source+mapRange.len {
			return mapRange.dest + (seed - mapRange.source)
		}
	}

	// It's not - just return what we got
	return seed
}

func pair(seeds []int) []seedPair {
	var out []seedPair

	for i := 0; i < (len(seeds))-1; i += 2 {
		out = append(out, seedPair{low: seeds[i], count: seeds[i+1]})
	}

	return out
}

////////////////////
// Code Graveyard //
////////////////////

func (f farmMap) lowestDestination() *rangeDef {
	lowest := f.maps[0]

	for _, fmap := range f.maps {
		if fmap.dest < lowest.dest {
			lowest = fmap
		}
	}

	return lowest
}

func (f farmMap) findSource(dest *rangeDef) *rangeDef {
	for _, fmap := range f.maps {
		if canProduce(dest, fmap) {
			var newSource, newLen int

			if dest.source > fmap.dest {
				newSource = fmap.source + (dest.source - fmap.dest)

				if dest.source+dest.len > fmap.dest+fmap.len {
					newLen = (fmap.dest + fmap.len) - newSource
				} else {
					newLen = (dest.source + dest.len) - newSource
				}
			} else {
				newSource = fmap.source

				if dest.source+dest.len > fmap.dest+fmap.len {
					newLen = (fmap.dest + fmap.len) - newSource
				} else {
					newLen = (dest.source + dest.len) - newSource
				}
			}

			return &rangeDef{
				source: newSource,
				dest:   0,
				len:    newLen,
			}
		}
	}

	return &rangeDef{source: 0, dest: 0, len: 0}
}

func canProduce(wanted, given *rangeDef) bool {
	if given.dest <= wanted.source && given.dest+given.len >= wanted.source {
		return true
	} else if given.dest >= wanted.source && given.dest+given.len <= wanted.source+wanted.len {
		return true
	}

	return false
}

/*

	// Second star
	// Work backwards: what is the lowest range for the last map?
	wanted := farmMaps[len(farmMaps)-1].lowestDestination() // We want something in that range (or lower, but we'll get there)
	var previous farmMap

	for i := 2; i <= len(farmMaps); i++ {
		previous = farmMaps[len(farmMaps)-i]
		Wanted := previous.findSource(wanted)
		fmt.Println(Wanted)
	}
*/
