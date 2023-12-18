package day18

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/maths"
	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day18 struct{}

func Factory() solution.Solver {
	return day18{}
}

type dir byte

const (
	none  dir = ' '
	up    dir = 'U'
	down  dir = 'D'
	left  dir = 'L'
	right dir = 'R'
)

var deltas = map[dir][]int{
	up:    {-1, 0},
	down:  {1, 0},
	left:  {0, -1},
	right: {0, 1},
}

type Coords struct {
	Y, X int
}

type Instruction struct {
	Direction dir
	Steps     int
}

type InstructionSet []Instruction

type Lagoon struct {
	Vertices  []Coords
	Perimeter int
}

func parseInstructionsPart2(reader io.Reader) InstructionSet {
	result := make(InstructionSet, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		steps, err := strconv.ParseInt(fields[2][2:len(fields[2])-2], 16, strconv.IntSize)
		ui.Die(err)

		d := none
		switch fields[2][7] {
		case '0':
			d = right
		case '1':
			d = down
		case '2':
			d = left
		case '3':
			d = up
		default:
			ui.Die(errors.New("Invalid direction"))
		}

		result = append(result, Instruction{
			Direction: d,
			Steps:     int(steps),
		})
	}
	return result
}

func parseInstructionsPart1(reader io.Reader) InstructionSet {
	result := make(InstructionSet, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		steps, err := strconv.Atoi(fields[1])
		ui.Die(err)

		result = append(result, Instruction{
			Direction: dir(fields[0][0]),
			Steps:     steps,
		})
	}
	return result
}

// Work through the instructions, storing the perimeter and vertices
// of the resulting polygon.
func (i InstructionSet) Drill() Lagoon {
	current := Coords{0, 0}
	result := Lagoon{
		Vertices: make([]Coords, 0),
	}

	for _, ins := range i {
		result.Vertices = append(result.Vertices, current)
		result.Perimeter += ins.Steps
		d := deltas[ins.Direction]
		current = Coords{
			current.Y + d[0]*ins.Steps,
			current.X + d[1]*ins.Steps,
		}
	}

	return result
}

// Area is calculated using the shoelace formula:
// https://rosettacode.org/wiki/Shoelace_formula_for_polygonal_area
func (l Lagoon) Area() int {
	area := 0
	for c := 0; c < len(l.Vertices)-1; c++ {
		area += l.Vertices[c].X*l.Vertices[c+1].Y - l.Vertices[c+1].X*l.Vertices[c].Y
	}
	area += l.Vertices[len(l.Vertices)-1].X*l.Vertices[0].Y - l.Vertices[0].X*l.Vertices[len(l.Vertices)-1].Y
	area = maths.AbsInt(area)

	return l.Perimeter + (area-l.Perimeter)/2 + 1
}

func (d day18) Part1(reader io.Reader) int {
	instructions := parseInstructionsPart1(reader)
	return instructions.Drill().Area()
}

func (d day18) Part2(reader io.Reader) int {
	instructions := parseInstructionsPart2(reader)
	return instructions.Drill().Area()
}
