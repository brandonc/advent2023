package day04

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/ds"
	"github.com/brandonc/advent2023/internal/input"
	"github.com/brandonc/advent2023/internal/maths"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day04 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day04{}
}

func scanCardNumbers(s string) []int {
	result := make([]int, 0)
	scanner := input.NewIntScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result = append(result, scanner.Int())
	}
	return result
}

func (d day04) Solve(reader io.Reader) (any, any) {
	scanner := bufio.NewScanner(reader)

	var cards, part1 = 0, 0
	// Ideally, you could use an array to keep track of the copy count, but I don't
	// yet know how many cards there are. So use a map of ID -> copy count
	copies := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		// String parsing using delimeters
		cardFieldSet := strings.Split(line, ":")
		cardID, _ := strconv.Atoi(strings.TrimSpace(cardFieldSet[0][len("Card "):]))
		numbersSets := strings.Split(cardFieldSet[1], "|")

		matching := ds.NewIntSet(
			scanCardNumbers(numbersSets[1]),
		).Intersect(
			ds.NewIntSet(
				scanCardNumbers(numbersSets[0]),
			),
		)

		// Add copies for each match
		for copy := 1; copy <= 1+copies[cardID]; copy++ {
			for n := cardID + 1; n <= cardID+len(matching); n++ {
				copies[n]++
			}
		}

		// Score doubles for each matching, which is 2^m when m >= 1
		if len(matching) >= 1 {
			part1 += maths.IntPow(2, len(matching)-1)
		}

		cards += 1
	}

	// Count original cards + copies
	var part2 = cards
	for _, c := range copies {
		part2 += c
	}

	return part1, part2
}
