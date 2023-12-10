package day10

import (
	"bufio"
	"io"
	"strings"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day10 struct{}

type Direction int

type Coord struct {
	Y, X int
}

type Loop struct {
	raw           [][]byte
	Start         Coord
	Visited       map[Coord]struct{}
	StartHasNorth bool
}

const (
	North Direction = 1
	East  Direction = 2
	South Direction = 3
	West  Direction = 4
)

var Directions = []Direction{North, East, South, West}

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	panic("Not a direction")
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day10{}
}

func parse(reader io.Reader) Loop {
	scanner := bufio.NewScanner(reader)
	raw := make([][]byte, 0)
	y, startY, startX := 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if s := strings.IndexByte(line, 'S'); s >= 0 {
			startX, startY = s, y
		}
		raw = append(raw, []byte(line))
		y += 1
	}

	return Loop{
		raw:     raw,
		Start:   Coord{startY, startX},
		Visited: make(map[Coord]struct{}),
	}
}

// Returns new coord and the direction from which you moved
func moveTo(y, x int, d Direction) (int, int, Direction) {
	switch d {
	case North:
		return y - 1, x, South
	case East:
		return y, x + 1, West
	case South:
		return y + 1, x, North
	case West:
		return y, x - 1, East
	}
	panic("Not a direction")
}

func (p Loop) connects(d Direction, y, x int) bool {
	switch d {
	case North:
		return p.raw[y][x] == '|' || p.raw[y][x] == 'J' || p.raw[y][x] == 'L'
	case East:
		return p.raw[y][x] == '-' || p.raw[y][x] == 'L' || p.raw[y][x] == 'F'
	case South:
		return p.raw[y][x] == '|' || p.raw[y][x] == '7' || p.raw[y][x] == 'F'
	case West:
		return p.raw[y][x] == '-' || p.raw[y][x] == 'J' || p.raw[y][x] == '7'
	}
	panic("Not a direction")
}

func (p *Loop) traverse() {
	y, x := p.Start.Y, p.Start.X
	var from Direction = 0

	var north, east, south, west byte = 0, 0, 0, 0

	// Initialize potential starting directions
	if y > 0 {
		north = p.raw[y-1][x]
	}
	if x < len(p.raw[0])-1 {
		east = p.raw[y][x+1]
	}
	if y < len(p.raw)-1 {
		south = p.raw[y+1][x]
	}
	if x > 0 {
		west = p.raw[y][x-1]
	}

	// Important to check north first for part 2
	if north == 'F' || north == '7' || north == '|' {
		p.StartHasNorth = true
		y, x, from = moveTo(y, x, North)
	} else if east == '7' || east == 'J' || east == '-' {
		y, x, from = moveTo(y, x, East)
	} else if south == 'J' || south == 'L' || south == '|' {
		y, x, from = moveTo(y, x, South)
	} else if west == 'F' || east == 'L' || east == '-' {
		y, x, from = moveTo(y, x, West)
	}

	for {
		p.Visited[Coord{y, x}] = struct{}{}

		// Returned to beginning
		if p.raw[y][x] == 'S' {
			break
		}

		for _, dir := range Directions {
			if from == dir {
				continue
			}
			if p.connects(dir, y, x) {
				y, x, from = moveTo(y, x, dir)
				break
			}
		}
	}
}

func (d day10) Part1(reader io.Reader) int {
	loop := parse(reader)
	loop.traverse()
	return len(loop.Visited) / 2
}

func (d day10) Part2(reader io.Reader) int {
	loop := parse(reader)
	loop.traverse()

	// For each place, count the number of north-ended pipes that were also
	// visited. All inside non-visited coordinates will have seen an odd
	// number of north-ended pipes.
	norths := 0
	inside := 0
	for y := 0; y < len(loop.raw); y++ {
		for x := 0; x < len(loop.raw[y]); x++ {
			piece := loop.raw[y][x]

			_, visited := loop.Visited[Coord{y, x}]
			if visited {
				if piece == 'L' || piece == 'J' || piece == '|' || (loop.StartHasNorth && piece == 'S') {
					norths += 1
				}
				continue
			}

			if norths%2 != 0 {
				inside += 1
			}
		}
	}
	return inside
}
