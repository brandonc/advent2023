package camelcards

import (
	"cmp"
	"fmt"

	"github.com/brandonc/advent2023/internal/ui"
)

type game struct {
	jackIsWild bool
}

func Game() game {
	return game{}
}

func GameWithJacksWild() game {
	return game{
		jackIsWild: true,
	}
}

func (g game) Ranks() []byte {
	if g.jackIsWild {
		return []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	} else {
		return []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	}
}

type HandType int

type Hand struct {
	Cards string
	Bid   int
}

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPair      HandType = 3
	OnePair      HandType = 2
	HighCard     HandType = 1
)

func histogram(s string) map[byte]int {
	result := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		result[s[i]] += 1
	}

	return result
}

func (g game) Type(hand string) HandType {
	var wildCard byte = 0
	if g.jackIsWild {
		wildCard = 'J'
	}

	hist := histogram(hand)

	if isFiveOfAKind(hist, wildCard) {
		return FiveOfAKind
	} else if isFourOfAKind(hist, wildCard) {
		return FourOfAKind
	} else if isFullHouse(hist, wildCard) {
		return FullHouse
	} else if isThreeOfAKind(hist, wildCard) {
		return ThreeOfAKind
	} else if isTwoPair(hist, wildCard) {
		return TwoPair
	} else if isOnePair(hist, wildCard) {
		return OnePair
	}
	return HighCard
}

func (g game) Score(hands []Hand) int {
	var score = 0
	for i := len(hands) - 1; i >= 0; i-- {
		score += hands[i].Bid * (i + 1)
	}
	return score
}

func isFiveOfAKind(hist map[byte]int, wildcard byte) bool {
	for c, v := range hist {
		if c == wildcard {
			if v == 5 {
				return true
			}
		}

		if v+hist[wildcard] == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(hist map[byte]int, wildcard byte) bool {
	for c, v := range hist {
		if c == wildcard {
			continue
		}

		if v+hist[wildcard] == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(hist map[byte]int, wildcard byte) bool {
	var three, two byte = 0, 0
	wildcardsUsed := 0

	for c, v := range hist {
		if c == wildcard {
			continue
		}

		if v == 3 {
			three = c
			break
		} else if v+hist[wildcard] == 3 {
			three = c
			wildcardsUsed = 3 - v
			break
		}
	}

	for c, v := range hist {
		if c == three || c == wildcard {
			continue
		}
		if v+hist[wildcard]-wildcardsUsed == 2 {
			two = c
			break
		}
	}

	return (three > 0 && two > 0)
}

func isThreeOfAKind(hist map[byte]int, wildcard byte) bool {
	for c, v := range hist {
		if c == wildcard {
			continue
		}

		if v+hist[wildcard] == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(hist map[byte]int, wildcard byte) bool {
	var a, b byte = 0, 0
	wildcardsUsed := 0
	for c, v := range hist {
		if c == wildcard {
			continue
		}

		if v+hist[wildcard]-wildcardsUsed == 2 {
			if a == 0 {
				a = c
				wildcardsUsed = 2 - v
			} else {
				b = c
			}
		}
	}
	return (a > 0 && b > 0)
}

func isOnePair(hist map[byte]int, wildcard byte) bool {
	for c, v := range hist {
		if c == wildcard {
			continue
		}

		if v+hist[wildcard] == 2 {
			return true
		}
	}
	return false
}

func (g game) rank(card byte) int {
	ranks := g.Ranks()
	for index, c := range ranks {
		if card == c {
			return len(ranks) - index
		}
	}
	return 0
}

func (g game) Compare(a, b Hand) int {
	primary := cmp.Compare(g.Type(a.Cards), g.Type(b.Cards))
	if primary != 0 {
		return primary
	}

	for i := 0; i < 5; i++ {
		aRank, bRank := g.rank(a.Cards[i]), g.rank(b.Cards[i])
		if aRank > bRank {
			return 1
		} else if aRank < bRank {
			return -1
		}
	}

	ui.Die(fmt.Errorf("two hands are identical: %q", a.Cards))
	return 0
}
