package day_05

import (
	"advent2023/helpers"
	"slices"
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

// Second star

func pair(seeds []int) []seedPair {
	var out []seedPair

	for i := 0; i < (len(seeds))-1; i += 2 {
		out = append(out, seedPair{low: seeds[i], count: seeds[i+1]})
	}

	return out
}

// Ranges and range operations

type intrange [][]int

func intersection(left, right intrange) intrange {
	var out intrange

	for _, leftpair := range left {
		for _, rightpair := range right {
			start := slices.Max([]int{leftpair[0], rightpair[0]})
			end := slices.Min([]int{leftpair[1], rightpair[1]})

			if start <= end { // If this is false, the ranges do not overlap at all.
				out = append(out, []int{start, end})
			}
		}
	}

	return mergeRanges(out)
}

func union(left, right intrange) intrange {
	allRanges := append(left, right...) // Put all the subranges together
	return mergeRanges(allRanges)
}

func sortRanges(left, right []int) int {
	return left[0] - right[0]
}

func mergeRanges(ranges intrange) intrange {
	if len(ranges) <= 1 { // Already as simple as it will ever be
		return ranges
	}

	slices.SortFunc(ranges, sortRanges)

	var mergedRanges intrange
	mergedRanges = append(mergedRanges, ranges[0])

	for i := 1; i < len(ranges); i++ {
		currentRange := ranges[i]
		previousRange := mergedRanges[len(mergedRanges)-1]

		if currentRange[0] <= previousRange[1] { // These ranges overlap
			previousRange[1] = slices.Max([]int{previousRange[1], currentRange[1]})
		} else { // These don't
			mergedRanges = append(mergedRanges, currentRange)
		}
	}

	return mergedRanges
}

func destToIntRange(fm farmMap) intrange {
	var out intrange

	for _, r := range fm.maps {
		out = append(out, []int{r.dest, r.dest + r.len})
	}

	return out
}

func sourceToIntRange(fm farmMap) intrange {
	var out intrange

	for _, r := range fm.maps {
		out = append(out, []int{r.source, r.source + r.len})
	}

	return out
}

// Another option: reverse transform
func reverse(outcome int, transformation farmMap) int {
	// Check whether the value is in any of the output ranges
	for _, mapRange := range transformation.maps {
		if outcome >= mapRange.dest && outcome <= mapRange.dest+mapRange.len {
			return mapRange.source + (outcome - mapRange.dest)
		}
	}

	// It's not - just return what we got
	return outcome
}

func contains(val int, pairs []seedPair) bool {
	for _, pair := range pairs {
		if val >= pair.low && val <= pair.low+pair.count {
			return true
		}
	}
	return false
}

// Third time is the charm!
// We need a function that maps an outcome (intrange) to an input (also intrange), based on a farmMap
// We know that the solution is in the outcome range [0, first], so we work back from there until we hit the seed pairs

func findInput(outcome intrange, fmap farmMap) intrange { // Does not handle passthrough values
	var out intrange

	for _, rd := range fmap.maps {
		common := intersection([][]int{{rd.dest, rd.dest + rd.len}}, outcome)
		if len(common) != 0 {
			for _, c := range common { // Individual pairs, subtract the offset to get to the unput
				c[0] = c[0] + (rd.source - rd.dest)
				c[1] = c[1] + (rd.source - rd.dest)
			}

			out = union(common, out)
		}
	}

	return out
}

func toIntRange(seeds []int) intrange {
	var out intrange

	for i := 0; i < (len(seeds))-1; i += 2 {
		out = append(out, []int{seeds[i], seeds[i] + seeds[i+1]})
	}

	return out
}
