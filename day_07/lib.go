package day_07

import (
	"strconv"
	"strings"
)

type hand struct {
	cards    string
	bet      int
	handType string
}

func parse(input []byte) []*hand {
	var out []*hand
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		cards := fields[0]
		bet, _ := strconv.Atoi(fields[1])

		out = append(out, &hand{cards: cards, bet: bet, handType: identify(cards)})
	}

	return out
}

func identify(cards string) string {
	localmap := make(map[rune]int)

	for _, r := range cards {
		if val, present := localmap[r]; present {
			localmap[r] = val + 1
		} else {
			localmap[r] = 1
		}
	}

	switch len(localmap) { // How many different cards do we have
	case 1:
		return "five of a kind"
	case 2: // Full house or four of a kind
		for _, count := range localmap {
			if count == 1 || count == 4 {
				return "four of a kind"
			} else {
				return "full house"
			}
		}
	case 3: // Two pair or three of a kind
		for _, count := range localmap {
			if count == 3 {
				return "three of a kind"
			}
		}
		return "two pair"
	case 4:
		return "one pair"
	case 5:
		return "high card"
	}
	return "not a type"
}

func compareHands(left, right *hand) int {
	hands := map[string]int{
		"five of a kind":  7,
		"four of a kind":  6,
		"full house":      5,
		"three of a kind": 4,
		"two pair":        3,
		"one pair":        2,
		"high card":       1,
	}

	return hands[left.handType] - hands[right.handType]
}

func compareCards(left, right *hand, second bool) int {
	var cards map[rune]int

	if !second {
		cards = map[rune]int{
			'A': 13,
			'K': 12,
			'Q': 11,
			'J': 10,
			'T': 9,
			'9': 8,
			'8': 7,
			'7': 6,
			'6': 5,
			'5': 4,
			'4': 3,
			'3': 2,
			'2': 1,
		}
	} else {
		cards = map[rune]int{
			'A': 13,
			'K': 12,
			'Q': 11,
			'T': 10,
			'9': 9,
			'8': 8,
			'7': 7,
			'6': 6,
			'5': 5,
			'4': 4,
			'3': 3,
			'2': 2,
			'J': 1,
		}
	}

	for i, r := range left.cards {
		if cards[r]-cards[rune(right.cards[i])] == 0 {
			continue
		} else {
			return cards[r] - cards[rune(right.cards[i])]
		}
	}
	return 0
}

func sorter(second bool) func(*hand, *hand) int { // Returns the sorting function
	return func(left, right *hand) int {
		if compareHands(left, right) != 0 {
			return compareHands(left, right)
		} else {
			return compareCards(left, right, second)
		}
	}
}

func score(hands []*hand) int {
	var out int

	for i, hand := range hands {
		out += (i + 1) * hand.bet
	}

	return out
}

func (h *hand) reidentify() {
	localmap := make(map[rune]int)

	for _, r := range h.cards {
		if val, present := localmap[r]; present {
			localmap[r] = val + 1
		} else {
			localmap[r] = 1
		}
	}

	if localmap['J'] == 0 || localmap['J'] == 5 { // No jacks or all jacks, no change.
		return
	}

	// We have at least one joker
	// Making the best hand always means adding it to the most common card (or one of them in case of a draw), because it will instantly transform it into a higher hand
	highestCount := 0
	var highestCard rune
	for card, count := range localmap {
		if card != 'J' {
			if count > highestCount {
				highestCount = count
				highestCard = card
			}
		}
	}

	// The jacks count as whatever that card was
	localmap[highestCard] += localmap['J']
	delete(localmap, 'J')

	switch len(localmap) { // How many different cards do we have
	case 1:
		h.handType = "five of a kind"
	case 2: // XXJYY or XXXJY
		product := 1
		for _, count := range localmap {
			product *= count
		}
		if product == 4 {
			h.handType = "four of a kind"
		} else if product == 6 {
			h.handType = "full house"
		}
	case 3: // Always three of a kind, we never use a jack to make a second pair
		h.handType = "three of a kind"
	case 4:
		h.handType = "one pair"
	}
}
