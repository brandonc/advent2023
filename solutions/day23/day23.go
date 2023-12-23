package day23

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day23 struct{}

func Factory() solution.Solver {
	return day23{}
}

type Coords struct {
	Y, X int
}

type Dir int

const (
	None  Dir = 0
	North Dir = 1
	West  Dir = 2
	South Dir = 3
	East  Dir = 4
)

type Map struct {
	Cells          [][]byte
	StartPosition  Coords
	finishPosition Coords
}

func (m Map) String() string {
	sb := strings.Builder{}
	for _, row := range m.Cells {
		sb.Write(row)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func parseMap(reader io.Reader) *Map {
	result := &Map{
		Cells: make([][]byte, 0),
	}

	scanner := bufio.NewScanner(reader)
	firstRow := true
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if firstRow {
			result.StartPosition = Coords{
				Y: 0,
				X: strings.IndexByte(line, '.'),
			}
			firstRow = false
		}
		result.Cells = append(result.Cells, []byte(line))
	}
	result.finishPosition = Coords{
		Y: len(result.Cells) - 1,
		X: strings.IndexByte(line, '.'),
	}

	return result
}

func (c Coords) Equals(other Coords) bool {
	return c.Y == other.Y && c.X == other.X
}

func (c Coords) String() string {
	return fmt.Sprintf("(%d, %d)", c.Y, c.X)
}

func (m *Map) CountMaxStepsNoSlopes(start Coords) (int, bool) {
	if start.Equals(m.finishPosition) {
		return 0, true
	}

	neighbors := [4]struct {
		C  Coords
		OK byte // Something besides '.' that indicates a valid neighbor
	}{
		{C: Coords{Y: start.Y - 1, X: start.X}},          // North
		{C: Coords{Y: start.Y, X: start.X - 1}},          // West
		{OK: 'v', C: Coords{Y: start.Y + 1, X: start.X}}, // South
		{OK: '>', C: Coords{Y: start.Y, X: start.X + 1}}, // East
	}

	var steps [4]int
	var exits [4]bool
	previous := m.Cells[start.Y][start.X]

	// DFS
	for i := 0; i < 4; i++ {
		try := neighbors[i].C
		if try.Y >= 0 && try.Y < len(m.Cells) && try.X >= 0 && try.X < len(m.Cells[0]) && (m.Cells[try.Y][try.X] == '.' || m.Cells[try.Y][try.X] == neighbors[i].OK) {
			// Use the map to mark the cell as visited for the current path
			m.Cells[start.Y][start.X] = '#'
			steps[i], exits[i] = m.CountMaxStepsNoSlopes(try)
			m.Cells[start.Y][start.X] = previous

			if !exits[i] {
				steps[i] = 0
			}
		}
	}

	return 1 + max(steps[0], steps[1], steps[2], steps[3]), exits[0] || exits[1] || exits[2] || exits[3]
}

func (m *Map) CountMaxSteps(start Coords) (int, bool) {
	if start.Equals(m.finishPosition) {
		return 0, true
	}

	neighbors := [4]Coords{
		{Y: start.Y - 1, X: start.X}, // North
		{Y: start.Y, X: start.X - 1}, // West
		{Y: start.Y + 1, X: start.X}, // South
		{Y: start.Y, X: start.X + 1}, // East
	}

	var steps [4]int
	var exits [4]bool
	previous := m.Cells[start.Y][start.X]

	// DFS
	for i := 0; i < 4; i++ {
		if neighbors[i].Y >= 0 && neighbors[i].Y < len(m.Cells) && neighbors[i].X >= 0 && neighbors[i].X < len(m.Cells[0]) && m.Cells[neighbors[i].Y][neighbors[i].X] != '#' {
			// Use the map to mark the cell as visited for the current path
			m.Cells[start.Y][start.X] = '#'
			steps[i], exits[i] = m.CountMaxSteps(neighbors[i])
			m.Cells[start.Y][start.X] = previous

			if !exits[i] {
				steps[i] = 0
			}
		}
	}

	return 1 + max(steps[0], steps[1], steps[2], steps[3]), exits[0] || exits[1] || exits[2] || exits[3]
}

func (d day23) Part1(reader io.Reader) int {
	theMap := parseMap(reader)

	steps, exit := theMap.CountMaxStepsNoSlopes(theMap.StartPosition)
	if !exit {
		ui.Die(errors.New("no exit found"))
	}
	return steps
}

func (d day23) Part2(reader io.Reader) int {
	theMap := parseMap(reader)

	steps, exit := theMap.CountMaxSteps(theMap.StartPosition)
	if !exit {
		ui.Die(errors.New("no exit found"))
	}
	return steps
}
