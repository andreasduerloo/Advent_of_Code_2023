package day_02

import (
	"regexp"
	"strconv"
	"strings"
)

// Input parsing
type game struct {
	id    int
	blue  int
	red   int
	green int
}

func parse(input []byte) []game {
	var out []game

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line != "" {
			out = append(out, newGame(line))
		}
	}

	return out
}

func newGame(input string) game {
	var out game

	idRe := regexp.MustCompile(`\d+`)
	id, _ := strconv.Atoi(idRe.FindString(input))

	out.id = id

	blueRe := regexp.MustCompile(`\d+ b`)
	redRe := regexp.MustCompile(`\d+ r`)
	greenRe := regexp.MustCompile(`\d+ g`)

	blueMatches := blueRe.FindAllString(input, -1)
	redMatches := redRe.FindAllString(input, -1)
	greenMatches := greenRe.FindAllString(input, -1)

	for _, blue := range blueMatches {
		val, _ := strconv.Atoi(idRe.FindString(blue))
		if val > out.blue {
			out.blue = val
		}
	}

	for _, red := range redMatches {
		val, _ := strconv.Atoi(idRe.FindString(red))
		if val > out.red {
			out.red = val
		}
	}

	for _, green := range greenMatches {
		val, _ := strconv.Atoi(idRe.FindString(green))
		if val > out.green {
			out.green = val
		}
	}

	return out
}

// First star
func possible(g, max game) bool {
	return g.blue <= max.blue && g.red <= max.red && g.green <= max.green
}

func reduceId(g game) int {
	return g.id
}

func filterReduce(s []game, max game, filter func(game, game) bool, reduce func(game) int) int {
	var out int

	for _, elem := range s {
		if filter(elem, max) {
			out += reduce(elem)
		}
	}

	return out
}

// Second star
func gamePower(g game) int {
	return g.blue * g.red * g.green
}

func mapReduce(s []game, reduce func(game) int) int {
	var out int

	for _, elem := range s {
		out += reduce(elem)
	}

	return out
}
