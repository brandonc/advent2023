package day13

import (
	"bufio"
	"errors"
	"io"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day13 struct{}

type Grid struct {
	Terrain [][]byte
	// A row to skip when finding horizontal symmetry
	NotRow int
	// A column to skip when finding vertical symmetry
	NotColumn int
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day13{}
}

func parseTerrain(reader io.Reader) []Grid {
	result := make([]Grid, 0)

	scanner := bufio.NewScanner(reader)
	grid := Grid{
		Terrain: make([][]byte, 0),
	}
	for scanner.Scan() {
		if scanner.Text() == "" {
			result = append(result, grid)
			grid = Grid{
				Terrain: make([][]byte, 0),
			}
			continue
		}

		grid.Terrain = append(grid.Terrain, []byte(scanner.Text()))
	}
	result = append(result, grid)
	return result
}

// Recursively
func (g Grid) compareColumnReflection(a, b int) bool {
	if a < 0 || b >= len(g.Terrain[0]) {
		return true
	}

	for y := 0; y < len(g.Terrain); y++ {
		if g.Terrain[y][a] != g.Terrain[y][b] {
			return false
		}
	}

	return g.compareColumnReflection(a-1, b+1)
}

func (g Grid) compareRowReflection(a, b int) bool {
	if a < 0 || b >= len(g.Terrain) {
		return true
	}

	for x := 0; x < len(g.Terrain[0]); x++ {
		if g.Terrain[a][x] != g.Terrain[b][x] {
			return false
		}
	}

	return g.compareRowReflection(a-1, b+1)
}

func (g Grid) VerticalSymmetry() int {
	for x := 0; x < len(g.Terrain[0])-1; x++ {
		if x+1 == g.NotColumn {
			continue
		}
		if g.compareColumnReflection(x, x+1) {
			return x + 1
		}
	}
	return 0
}

func (g Grid) HorizontalSymmetry() int {
	for y := 0; y < len(g.Terrain)-1; y++ {
		if y+1 == g.NotRow {
			continue
		}
		if g.compareRowReflection(y, y+1) {
			return y + 1
		}
	}
	return 0
}

func (g Grid) Clone() Grid {
	result := Grid{
		Terrain: make([][]byte, len(g.Terrain)),
	}
	for y := 0; y < len(g.Terrain); y++ {
		result.Terrain[y] = make([]byte, len(g.Terrain[y]))
		copy(result.Terrain[y], g.Terrain[y])
	}
	return result
}

func (g Grid) Smudges() []Grid {
	// Wasteful to precompute all grid permutations but ultimately it's not
	// that much data and also helps to reuse code from part 1.
	result := make([]Grid, 0, len(g.Terrain[0])*len(g.Terrain))

	for y := 0; y < len(g.Terrain); y++ {
		for x := 0; x < len(g.Terrain[0]); x++ {
			permutation := g.Clone()
			if g.Terrain[y][x] == '#' {
				permutation.Terrain[y][x] = '.'
			} else {
				permutation.Terrain[y][x] = '#'
			}
			result = append(result, permutation)
		}
	}

	return result
}

func (d day13) Part1(reader io.Reader) int {
	grids := parseTerrain(reader)

	sum := 0
	for _, grid := range grids {
		if col := grid.VerticalSymmetry(); col > 0 {
			sum += col
		} else if row := grid.HorizontalSymmetry(); row > 0 {
			sum += 100 * row
		}
	}

	return sum
}

func (d day13) Part2(reader io.Reader) int {
	grids := parseTerrain(reader)

	sum := 0
	for _, grid := range grids {
		found := false

		for _, fixed := range grid.Smudges() {
			fixed.NotColumn = grid.VerticalSymmetry()
			if col := fixed.VerticalSymmetry(); col > 0 {
				sum += col
				found = true
				break
			}

			fixed.NotRow = grid.HorizontalSymmetry()
			if row := fixed.HorizontalSymmetry(); row > 0 {
				sum += 100 * row
				found = true
				break
			}
		}

		if !found {
			ui.Die(errors.New("No solution found"))
		}
	}

	return sum
}
