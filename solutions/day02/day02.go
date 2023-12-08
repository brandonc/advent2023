package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/maths"
	"github.com/brandonc/advent2023/solutions/solution"
)

var impossibleAmounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type day02 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return &day02{}
}

func (d day02) Part1(reader io.Reader) int {
	result, _ := d.solve(reader)
	return result.(int)
}

func (d day02) Part2(reader io.Reader) int {
	_, result := d.solve(reader)
	return result.(int)
}

func (d day02) solve(reader io.Reader) (any, any) {
	var part1, part2 = 0, 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		idSplit := strings.SplitN(line, ":", 2)
		idRaw := strings.TrimSpace(idSplit[0][len("Game "):])
		id, _ := strconv.Atoi(idRaw)

		state := struct {
			id         int
			impossible bool
			minimums   map[string]int
		}{
			id: id,
			// A game is impossible if a set contains more than 12 red cubes,
			// 13 green cubes, or 14 blue cubes seen at once.
			impossible: false,

			// Track the minimum number of each color cubes in each game, based
			// on the number of each seen at once.
			minimums: map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			},
		}

		// Each set is separated by a semicolon
		setSplit := strings.Split(idSplit[1], ";")
		for _, set := range setSplit {
			// Each color amount is separated by a comma
			colorSplit := strings.Split(set, ",")
			for _, color := range colorSplit {
				// The color and amount are separated by a space
				fieldSplit := strings.Split(strings.TrimSpace(color), " ")
				numColor, _ := strconv.Atoi(fieldSplit[0])

				if numColor > impossibleAmounts[fieldSplit[1]] {
					state.impossible = true
				}

				state.minimums[fieldSplit[1]] = maths.Max(state.minimums[fieldSplit[1]], numColor)
			}
		}

		if !state.impossible {
			part1 += state.id
		}

		part2 += state.minimums["red"] * state.minimums["green"] * state.minimums["blue"]
	}
	return part1, part2
}
