package day06

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/input"
	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day06 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day06{}
}

type Race struct {
	TimeMS   int
	RecordMM int
}

func parseInts(s string) []int {
	scanner := input.NewIntScanner(strings.NewReader(strings.TrimSpace(s)))
	scanner.Split(bufio.ScanWords)
	result := make([]int, 0)
	for scanner.Scan() {
		result = append(result, scanner.Int())
	}
	return result
}

func parseRacePart2(s string) Race {
	lines := strings.Split(s, "\n")

	timesRaw := strings.Split(lines[0], ":")
	distancesRaw := strings.Split(lines[1], ":")

	time, err := strconv.Atoi(strings.ReplaceAll(timesRaw[1], " ", ""))
	ui.Die(err)

	distance, err := strconv.Atoi(strings.ReplaceAll(distancesRaw[1], " ", ""))
	ui.Die(err)

	return Race{
		TimeMS:   time,
		RecordMM: distance,
	}
}

func parseRacesPart1(s string) []Race {
	lines := strings.Split(s, "\n")

	times := parseInts(strings.Split(lines[0], ":")[1])
	distances := parseInts(strings.Split(lines[1], ":")[1])

	result := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		result[i] = Race{
			TimeMS:   times[i],
			RecordMM: distances[i],
		}
	}

	return result
}

func (race Race) WaysToWin() int {
	ways := 0
	for hold := 1; hold < race.TimeMS; hold++ {
		if (race.TimeMS-hold)*hold > race.RecordMM {
			ways += 1
		}
	}
	return ways
}

func (d day06) Solve(reader io.Reader) (any, any) {
	input, err := io.ReadAll(reader)
	ui.Die(err)

	races := parseRacesPart1(string(input))

	part1 := 1
	for _, race := range races {
		part1 *= race.WaysToWin()
	}

	part2Race := parseRacePart2(string(input))

	return part1, part2Race.WaysToWin()
}
