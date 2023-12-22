package day22

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day22 struct{}

func Factory() solution.Solver {
	return day22{}
}

type Brick struct {
	// Bricks are a 3-dimensional line of cubes. Two sets of coordinates
	// represent the two ends of the line. For example, 2,2,2/2,2,3 would be a
	// two cubes oriented vertically in the z dimension.
	startX, startY, startZ int
	endX, endY, endZ       int
	name                   byte
}

type Game struct {
	Bricks []Brick
}

func (b Brick) String() string {
	return fmt.Sprintf("%c: %d,%d,%d~%d,%d,%d", b.name, b.startX, b.startY, b.startZ, b.endX, b.endY, b.endZ)
}

func (b Brick) Collides(other Brick) bool {
	// Bricks collide if range startX..endX intersects with other.startX..other.endX
	// and range startY..endY intersects with other.startY..other.endY
	// while this brick is as least as far down as other.
	zOverlap := b.startZ <= other.endZ
	xOverlap := b.endX >= other.startX && b.startX <= other.endX
	yOverlap := b.endY >= other.startY && b.startY <= other.endY

	return zOverlap && xOverlap && yOverlap
}

func (b *Brick) Move(dx, dy, dz int) {
	b.startX += dx
	b.startY += dy
	b.startZ += dz
	b.endX += dx
	b.endY += dy
	b.endZ += dz
}

func (g *Game) simulateBrick(index int, brick *Brick) int {
	// Move brick down until it rests on the ground or on another brick.
	initialZ := brick.startZ
next:
	for {
		if brick.startZ == 1 {
			ui.Debugf("Brick %c comes to rest at height 1", brick.name)
			break
		}

		brick.Move(0, 0, -1)

		for or := 0; or < index; or++ {
			if brick.Collides(g.Bricks[or]) {
				// Too far
				brick.Move(0, 0, 1)
				ui.Debugf("Brick %c comes to rest at height %d, on top of %c", brick.name, brick.startZ, g.Bricks[or].name)
				break next
			}
		}
	}

	ui.Debugf("Brick %c fell %d units", brick.name, initialZ-brick.startZ)
	return initialZ - brick.startZ
}

func (g *Game) Simulate() int {
	// Sort bricks by z-order. Sarting with the lowest brick, move it until
	// it rests on the ground or on another brick in the x or y dimension.

	fell := 0

	for br := 0; br < len(g.Bricks); br++ {
		brick := &g.Bricks[br]
		ui.Debugf("Simulating brick %c, which is at %d", brick.name, brick.startZ)

		if g.simulateBrick(br, brick) > 0 {
			fell += 1
		}
	}

	return fell
}

func parseCoords(coords string) (int, int, int) {
	result := make([]int, 3)
	for i, s := range strings.Split(coords, ",") {
		result[i], _ = strconv.Atoi(s)
	}

	return result[0], result[1], result[2]
}

func parseBricks(reader io.Reader) *Game {
	game := Game{
		Bricks: make([]Brick, 0),
	}

	scanner := bufio.NewScanner(reader)
	var nextChar byte = 'A'
	for scanner.Scan() {
		line := scanner.Text()

		coordsSet := strings.Split(line, "~")
		startX, startY, startZ := parseCoords(coordsSet[0])
		endX, endY, endZ := parseCoords(coordsSet[1])

		game.Bricks = append(game.Bricks, Brick{
			startX, startY, startZ, endX, endY, endZ,
			nextChar,
		})

		nextChar += 1
	}

	slices.SortFunc[[]Brick](game.Bricks, func(a, b Brick) int {
		if a.startZ == b.startZ {
			return 0
		}
		if a.startZ < b.startZ {
			return -1
		}
		return 1
	})

	return &game
}

func (g *Game) SimulateWithoutBrick(b int) int {
	gameWithoutBrick := Game{
		Bricks: make([]Brick, len(g.Bricks)-1),
	}
	copy(gameWithoutBrick.Bricks, g.Bricks[:b])
	copy(gameWithoutBrick.Bricks[b:], g.Bricks[b+1:])

	return gameWithoutBrick.Simulate()
}

func (d day22) Part1(reader io.Reader) int {
	game := parseBricks(reader)
	game.Simulate()

	safeToRemove := 0
	for b := 0; b < len(game.Bricks); b++ {
		if game.SimulateWithoutBrick(b) == 0 {
			safeToRemove += 1
		}
	}
	return safeToRemove
}

func (d day22) Part2(reader io.Reader) int {
	game := parseBricks(reader)
	game.Simulate()

	sum := 0
	for b := 0; b < len(game.Bricks); b++ {
		sum += game.SimulateWithoutBrick(b)
	}
	return sum
}
