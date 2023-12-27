package day24

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day24 struct {
	testAreaMin int
	testAreaMax int
}

func Factory() solution.Solver {
	return day24{
		testAreaMin: 200000000000000,
		testAreaMax: 400000000000000,
	}
}

type Line struct {
	X, Y, Z    int
	VX, VY, VZ int
}

func (l Line) ToLeft() bool {
	return l.VX < 0
}

func (l Line) TwoPoints() (int, int, int, int) {
	return l.X, l.Y, l.X + l.VX, l.Y + l.VY
}

func (l Line) Intersects(other Line) (bool, float64, float64) {
	x1, y1, x2, y2 := l.TwoPoints()
	x3, y3, x4, y4 := other.TwoPoints()

	denom := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	if denom == 0 {
		return false, 0, 0
	}
	ua := float64((x4-x3)*(y1-y3)-(y4-y3)*(x1-x3)) / float64(denom)

	return true,
		float64(x1) + ua*float64(x2-x1),
		float64(y1) + ua*float64(y2-y1)
}

type Lines []Line

func parseLines(reader io.Reader) Lines {
	scanner := bufio.NewScanner(reader)
	result := make(Lines, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "@")

		coordsParser := func(s string) (int, int, int) {
			nums := strings.Split(s, ",")
			x, _ := strconv.Atoi(strings.TrimSpace(nums[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(nums[1]))
			z, _ := strconv.Atoi(strings.TrimSpace(nums[2]))

			return x, y, z
		}

		x, y, z := coordsParser(fields[0])
		vx, vy, vz := coordsParser(fields[1])

		result = append(result, Line{
			x, y, z, vx, vy, vz,
		})
	}
	return result
}

func (d day24) Part1(reader io.Reader) int {
	lines := parseLines(reader)

	intersecting := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		for o := i + 1; o < len(lines); o++ {
			other := lines[o]
			if intersects, x, y := line.Intersects(other); intersects {
				dx := x - float64(line.X)
				dy := y - float64(line.Y)
				pastA := (dx > 0) != (line.VX > 0) || (dy > 0) != (line.VY > 0)

				dx = x - float64(other.X)
				dy = y - float64(other.Y)
				pastB := (dx > 0) != (other.VX > 0) || (dy > 0) != (other.VY > 0)

				if pastA || pastB {
					continue
				}

				if x >= float64(d.testAreaMin) && y >= float64(d.testAreaMin) && x <= float64(d.testAreaMax) && y <= float64(d.testAreaMax) {
					intersecting += 1
				}
			}
		}
	}
	return intersecting
}

func (d day24) Part2(reader io.Reader) int {
	return 0
}
