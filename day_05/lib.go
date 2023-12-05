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
	low  int
	high int
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
		out = append(out, seedPair{low: seeds[i], high: seeds[i+1]})
	}

	return out
}
