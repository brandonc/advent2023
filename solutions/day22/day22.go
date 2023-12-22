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
	ID                     int
}

type Game []Brick

func (b Brick) String() string {
	return fmt.Sprintf("%d: %d,%d,%d~%d,%d,%d", b.ID, b.startX, b.startY, b.startZ, b.endX, b.endY, b.endZ)
}

func (b Brick) Collides(other Brick) bool {
	// Bricks collide if range startX..endX intersects with other.startX..other.endX
	// and range startY..endY intersects with other.startY..other.endY
	// while this brick is as least as far down as other.
	zOverlap := b.startZ <= other.endZ
	if !zOverlap {
		return false
	}

	xOverlap := b.endX >= other.startX && b.startX <= other.endX
	yOverlap := b.endY >= other.startY && b.startY <= other.endY

	return xOverlap && yOverlap
}

func (b *Brick) MoveZ(dz int) {
	b.startZ += dz
	b.endZ += dz
}

func (g Game) simulateBrick(index int) int {
	// Move brick down until it rests on the ground or on another brick.
	brick := &g[index]
	initialZ := brick.startZ
next:
	for {
		if brick.startZ == 1 {
			ui.Debugf("Brick %d comes to rest at height 1", brick.ID)
			break
		}

		brick.MoveZ(-1)

		for or := 0; or < index; or++ {
			// This is the slow part: checking every lower brick at each step down.
			// There must be a faster way to do this.
			if brick.Collides(g[or]) {
				// Too far, so move back up one
				brick.MoveZ(1)
				ui.Debugf("Brick %d comes to rest at height %d, on top of %d", brick.ID, brick.startZ, g[or].ID)
				break next
			}
		}
	}

	ui.Debugf("Brick %d fell %d units", brick.ID, initialZ-brick.startZ)
	return initialZ - brick.startZ
}

func (g Game) Simulate() int {
	// Bricks are already sorted by z-order. Sarting with the lowest brick,
	// move it until it rests on the ground or on another brick in
	// the x or y dimension.
	fell := 0

	for br := 0; br < len(g); br++ {
		if g.simulateBrick(br) > 0 {
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

func parseBricks(reader io.Reader) Game {
	game := make(Game, 0)

	scanner := bufio.NewScanner(reader)
	nextID := 1
	for scanner.Scan() {
		line := scanner.Text()

		coordsSet := strings.Split(line, "~")
		startX, startY, startZ := parseCoords(coordsSet[0])
		endX, endY, endZ := parseCoords(coordsSet[1])

		game = append(game, Brick{
			startX, startY, startZ, endX, endY, endZ,
			nextID,
		})

		nextID += 1
	}

	slices.SortFunc[[]Brick](game, func(a, b Brick) int {
		if a.startZ == b.startZ {
			return 0
		}
		if a.startZ < b.startZ {
			return -1
		}
		return 1
	})

	return game
}

func (g Game) SimulateWithoutBrick(b int) int {
	gameWithoutBrick := make(Game, len(g)-1)
	copy(gameWithoutBrick, g[:b])
	copy(gameWithoutBrick[b:], g[b+1:])

	return gameWithoutBrick.Simulate()
}

func (d day22) Part1(reader io.Reader) int {
	game := parseBricks(reader)
	game.Simulate()

	safeToRemove := 0
	for b := 0; b < len(game); b++ {
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
	for b := 0; b < len(game); b++ {
		sum += game.SimulateWithoutBrick(b)
	}
	return sum
}
