package day16

import (
	"bufio"
	"fmt"
	"io"

	"github.com/brandonc/advent2023/internal/maths"
	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day16 struct{}

func Factory() solution.Solver {
	return day16{}
}

type Coords struct{ y, x int }
type CoordsVec struct{ y, x, dy, dx int }

type Grid struct {
	Rows [][]byte
	// A copy of the grid with the energized tiles marked with a #
	Energized map[Coords]struct{}
	seen      map[CoordsVec]struct{}
}

func parseGrid(reader io.Reader) *Grid {
	var result = Grid{
		Rows:      make([][]byte, 0),
		Energized: make(map[Coords]struct{}),
		seen:      make(map[CoordsVec]struct{}),
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		result.Rows = append(result.Rows, []byte(line))
	}
	return &result
}

func (c CoordsVec) Northbound() bool {
	return c.dy == -1
}

func (c CoordsVec) Southbound() bool {
	return c.dy == 1
}

func (c CoordsVec) Eastbound() bool {
	return c.dx == 1
}

func (c CoordsVec) Westbound() bool {
	return c.dx == -1
}

func (c CoordsVec) Continue() CoordsVec {
	return CoordsVec{c.y + c.dy, c.x + c.dx, c.dy, c.dx}
}

func (c CoordsVec) MirrorForward() CoordsVec {
	switch {
	case c.Northbound():
		return CoordsVec{c.y, c.x + 1, 0, 1}
	case c.Southbound():
		return CoordsVec{c.y, c.x - 1, 0, -1}
	case c.Eastbound():
		return CoordsVec{c.y - 1, c.x, -1, 0}
	case c.Westbound():
		return CoordsVec{c.y + 1, c.x, 1, 0}
	default:
		return c
	}
}

func (c CoordsVec) MirrorReverse() CoordsVec {
	switch {
	case c.Northbound():
		return CoordsVec{c.y, c.x - 1, 0, -1}
	case c.Southbound():
		return CoordsVec{c.y, c.x + 1, 0, 1}
	case c.Eastbound():
		return CoordsVec{c.y + 1, c.x, 1, 0}
	case c.Westbound():
		return CoordsVec{c.y - 1, c.x, -1, 0}
	default:
		return c
	}
}

func (g *Grid) laser(c CoordsVec) {
	y, x, dy, dx := c.y, c.x, c.dy, c.dx
	if y < 0 || y >= len(g.Rows) || x < 0 || x >= len(g.Rows[y]) {
		return
	}

	// If laser was already called with these parameters, that is a cycle
	// that should be ignored.
	if _, ok := g.seen[c]; ok {
		return
	}
	g.seen[c] = struct{}{}

	// Mark the tile as energized
	g.Energized[Coords{y, x}] = struct{}{}

	switch g.Rows[y][x] {
	case '.':
		// continue in direction
		g.laser(c.Continue())
		return
	case '/':
		// 90 degrees left if going e/w, 90 degrees right if going n/s
		g.laser(c.MirrorForward())
		return
	case '\\':
		// 90 degrees right if going e/w, 90 degrees left if going n/s
		g.laser(c.MirrorReverse())
		return
	case '|':
		// continue in direction if going n/s
		if dx == 0 {
			g.laser(c.Continue())
			return
		}
		// split if going e/w
		g.laser(CoordsVec{y + 1, x, 1, 0})
		g.laser(CoordsVec{y - 1, x, -1, 0})
		return
	case '-':
		// continue in direction if going e/w
		if dy == 0 {
			g.laser(c.Continue())
			return
		}
		// split if going n/s
		g.laser(CoordsVec{y, x + 1, 0, 1})
		g.laser(CoordsVec{y, x - 1, 0, -1})
		return
	default:
		ui.Die(fmt.Errorf("Unknown character in grid: %c at %d, %d", g.Rows[y][x], y, x))
	}
}

func (g *Grid) TurnOn(init CoordsVec) {
	g.laser(init)
}

func (g *Grid) TurnOff() {
	for c := range g.Energized {
		delete(g.Energized, c)
	}
	for c := range g.seen {
		delete(g.seen, c)
	}
}

func (g Grid) CountEnergized() int {
	return len(g.Energized)
}

func (d day16) Part1(reader io.Reader) int {
	grid := parseGrid(reader)
	grid.TurnOn(CoordsVec{0, 0, 0, 1})

	return grid.CountEnergized()
}

func (d day16) Part2(reader io.Reader) int {
	grid := parseGrid(reader)

	max := 0
	// South from top
	for x := 0; x < len(grid.Rows[0]); x++ {
		grid.TurnOn(CoordsVec{0, x, 1, 0})
		max = maths.MaxInt(max, grid.CountEnergized())
		grid.TurnOff()
	}

	// North from bottom
	for x := 0; x < len(grid.Rows[0]); x++ {
		grid.TurnOn(CoordsVec{len(grid.Rows) - 1, x, -1, 0})
		max = maths.MaxInt(max, grid.CountEnergized())
		grid.TurnOff()
	}

	// East from left
	for y := 0; y < len(grid.Rows); y++ {
		grid.TurnOn(CoordsVec{y, 0, 0, 1})
		max = maths.MaxInt(max, grid.CountEnergized())
		grid.TurnOff()
	}

	// West from right
	for y := 0; y < len(grid.Rows); y++ {
		grid.TurnOn(CoordsVec{y, len(grid.Rows[y]) - 1, 0, -1})
		max = maths.MaxInt(max, grid.CountEnergized())
		grid.TurnOff()
	}

	return max
}
