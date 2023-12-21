package day21

import (
	"bufio"
	"io"
	"strings"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day21 struct {
	testSteps int
}

func Factory() solution.Solver {
	return day21{}
}

type Coords struct {
	MapY, MapX int
	Y, X       int
}

type Coords2D struct {
	Y, X int
}

type Garden struct {
	Map             [][]byte
	InitialPosition Coords
	positions       map[Coords]struct{}
	Rocks           int
}

func parseGarden(reader io.Reader) *Garden {
	result := Garden{
		Map: make([][]byte, 0),
	}

	scanner := bufio.NewScanner(reader)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if pos := strings.IndexByte(line, 'S'); pos >= 0 {
			result.InitialPosition = Coords{0, 0, y, pos}
			result.positions = map[Coords]struct{}{
				{0, 0, y, pos}: {},
			}
		}
		result.Rocks += strings.Count(line, "#")
		result.Map = append(result.Map, []byte(line))
		y += 1
	}

	return &result
}

func (g *Garden) CountPossiblePositions(stepsRemaining int) int {
	if stepsRemaining == 0 {
		return len(g.positions)
	}

	nextPositions := make(map[Coords]struct{}, len(g.positions)*2)
	for pos := range g.positions {
		for _, nextPos := range []Coords{
			{0, 0, pos.Y - 1, pos.X},
			{0, 0, pos.Y + 1, pos.X},
			{0, 0, pos.Y, pos.X - 1},
			{0, 0, pos.Y, pos.X + 1},
		} {
			if nextPos.Y < 0 || nextPos.Y >= len(g.Map) {
				continue
			}

			if nextPos.X < 0 || nextPos.X >= len(g.Map[nextPos.Y]) {
				continue
			}

			if g.Map[nextPos.Y][nextPos.X] == '#' {
				continue
			}

			nextPositions[nextPos] = struct{}{}
		}
	}

	g.positions = nextPositions
	return g.CountPossiblePositions(stepsRemaining - 1)
}

// Doesn't duplicate the map, but tracks all positions, which is way too slow.
func (g *Garden) CountPossiblePositionsInfinite(stepsRemaining int) int {
	if stepsRemaining == 0 {
		return len(g.positions)
	}

	nextPositions := make(map[Coords]struct{}, len(g.positions))
	for pos := range g.positions {
		for _, delta := range []Coords2D{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			nextPos := Coords{
				MapY: pos.MapY,
				MapX: pos.MapX,
				Y:    pos.Y + delta.Y,
				X:    pos.X + delta.X,
			}
			if nextPos.Y < 0 {
				nextPos.MapY -= 1
				nextPos.Y = len(g.Map) - 1
			} else if nextPos.Y >= len(g.Map) {
				nextPos.MapY += 1
				nextPos.Y = 0
			} else if nextPos.X < 0 {
				nextPos.MapX -= 1
				nextPos.X = len(g.Map[pos.Y]) - 1
			} else if nextPos.X >= len(g.Map[pos.Y]) {
				nextPos.MapX += 1
				nextPos.X = 0
			}

			if g.Map[nextPos.Y][nextPos.X] == '#' {
				continue
			}

			nextPositions[nextPos] = struct{}{}
		}
	}

	g.positions = nextPositions
	return g.CountPossiblePositionsInfinite(stepsRemaining - 1)
}

func (d day21) Part1(reader io.Reader) int {
	gardenMap := parseGarden(reader)
	steps := 64
	if d.testSteps > 0 {
		steps = d.testSteps
	}
	return gardenMap.CountPossiblePositions(steps)
}

func (d day21) Part2(reader io.Reader) int {
	gardenMap := parseGarden(reader)
	steps := 64
	if d.testSteps > 0 {
		steps = d.testSteps
	}
	return gardenMap.CountPossiblePositionsInfinite(steps)
}
