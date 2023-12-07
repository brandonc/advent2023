package day07

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/solutions/day07/camelcards"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day07 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day07{}
}

func parseInput(reader io.Reader) []camelcards.Hand {
	scanner := bufio.NewScanner(reader)
	hands := make([]camelcards.Hand, 0)

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, camelcards.Hand{
			Cards: fields[0],
			Bid:   bid,
		})
	}

	return hands
}

func (d day07) Solve(reader io.Reader) (any, any) {
	hands := parseInput(reader)

	part1 := camelcards.Game()
	slices.SortFunc(hands, part1.Compare)

	part1Score := part1.Score(hands)

	part2 := camelcards.GameWithJacksWild()
	slices.SortFunc(hands, part2.Compare)

	return part1Score, part2.Score(hands)
}
