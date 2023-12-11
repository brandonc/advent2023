package day11

import (
	"bufio"
	"bytes"
	"io"
	"slices"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day11 struct{}

type Coord struct {
	Y, X int
}

type Image struct {
	raw       [][]byte
	emptyRows []int
	emptyCols []int
	Galaxies  []Coord
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day11{}
}

func (g Image) NumberOfPairs() int {
	result := 0
	for i := len(g.Galaxies) - 1; i >= 1; i-- {
		result += i
	}
	return result
}

func (g Image) Pairs() [][]Coord {
	result := make([][]Coord, g.NumberOfPairs())
	index := 0
	for a := 0; a < len(g.Galaxies); a++ {
		for b := a + 1; b < len(g.Galaxies); b++ {
			result[index] = []Coord{
				g.Galaxies[a],
				g.Galaxies[b],
			}
			index += 1
		}
	}
	return result
}

func (g Image) String() string {
	return string(bytes.Join(g.raw, []byte("\n")))
}

func parseImage(reader io.Reader) *Image {
	scanner := bufio.NewScanner(reader)
	result := &Image{
		raw:      make([][]byte, 0),
		Galaxies: make([]Coord, 0),
	}
	for scanner.Scan() {
		row := []byte(scanner.Text())
		result.raw = append(result.raw, row)

		if !bytes.Contains(row, []byte{'#'}) {
			result.emptyRows = append(result.emptyRows, len(result.raw)-1)
		}
	}

	for column := 0; column < len(result.raw[0]); column++ {
		containsGalaxy := false
		for row := 0; row < len(result.raw); row++ {
			if result.raw[row][column] == '#' {
				containsGalaxy = true
				result.Galaxies = append(result.Galaxies, Coord{row, column})
			}
		}
		if !containsGalaxy {
			result.emptyCols = append(result.emptyCols, column)
		}
	}

	return result
}

func measurePairDim(a, b int, emptyDim []int, emptyDistance int) int {
	result := 0
	d := 1
	if a > b {
		d = -1
	}
	for move := a; move != b; move += d {
		result += 1
		if _, found := slices.BinarySearch[[]int](emptyDim, move); found {
			result += emptyDistance
			if emptyDistance > 1 {
				result -= 1
			}
		}
	}
	return result
}

func (i Image) MeasurePairs(emptyDistance int) int {
	pairs := i.Pairs()
	sum := 0
	for pair := 0; pair < len(pairs); pair++ {
		a, b := pairs[pair][0], pairs[pair][1]

		sum += measurePairDim(a.X, b.X, i.emptyCols, emptyDistance) +
			measurePairDim(a.Y, b.Y, i.emptyRows, emptyDistance)
	}
	return sum
}

func (d day11) Part1(reader io.Reader) int {
	image := parseImage(reader)
	return image.MeasurePairs(1)
}

func (d day11) Part2(reader io.Reader) int {
	image := parseImage(reader)
	return image.MeasurePairs(1_000_000)
}
