package camelcards

import (
	"cmp"
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

func histogram(s string) map[byte]int {
	result := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		result[s[i]] += 1
	}

	return result
}

func (g game) WildCard() byte {
	if g.jackIsWild {
		return 'J'
	}
	return '0'
}

func (g game) rankHand(hand string) HandType {
	var wildCard byte = g.WildCard()

	hist := histogram(hand)

	if isFiveOfAKind(hist, wildCard) {
		return 7
	} else if isFourOfAKind(hist, wildCard) {
		return 6
	} else if isFullHouse(hist, wildCard) {
		return 5
	} else if isThreeOfAKind(hist, wildCard) {
		return 4
	} else if isTwoPair(hist, wildCard) {
		return 3
	} else if isOnePair(hist, wildCard) {
		return 2
	}
	return 1
}

func (g game) rankCard(card byte) int {
	ranks := g.Ranks()
	for index, c := range ranks {
		if card == c {
			return len(ranks) - index
		}
	}
	return 0
}

func isFiveOfAKind(hist map[byte]int, wildcard byte) bool {
	for c, v := range hist {
		if c == wildcard && v == 5 {
			return true
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

func (g game) Compare(a, b Hand) int {
	if primary := cmp.Compare(g.rankHand(a.Cards), g.rankHand(b.Cards)); primary != 0 {
		return primary
	}

	for i := 0; i < len(a.Cards); i++ {
		aRank, bRank := g.rankCard(a.Cards[i]), g.rankCard(b.Cards[i])
		if aRank > bRank {
			return 1
		} else if aRank < bRank {
			return -1
		}
	}

	return 0
}

func (g game) Score(hands []Hand) int {
	var score = 0
	for i := len(hands) - 1; i >= 0; i-- {
		score += hands[i].Bid * (i + 1)
	}
	return score
}
