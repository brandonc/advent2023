package day14

import (
	"bufio"
	"io"
	"strings"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day14 struct{}

type Platform [][]byte

var Rock byte = 'O'
var Empty byte = '.'

func Factory() solution.Solver {
	return day14{}
}

func parsePlatform(reader io.Reader) Platform {
	scanner := bufio.NewScanner(reader)
	result := make(Platform, 0)

	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}

	return result
}

func (f Platform) Tilt(dy, dx int) {
	done := true
	for {
		done = true
		for y := 0; y < len(f); y++ {
			for x := 0; x < len(f[y]); x++ {
				if f[y][x] == Rock {
					// Range check
					if y+dy < 0 || y+dy >= len(f) || x+dx < 0 || x+dx >= len(f[y]) {
						continue
					}

					if f[y+dy][x+dx] == Empty {
						f[y+dy][x+dx] = Rock
						f[y][x] = Empty
						done = false
					}
				}
			}
		}

		if done {
			break
		}
	}
}

func (f Platform) WeightLoad() int {
	result := 0

	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if f[y][x] == Rock {
				result += len(f) - y
			}
		}
	}
	return result
}

func (f Platform) String() string {
	builder := strings.Builder{}
	for y := 0; y < len(f); y++ {
		builder.WriteString(string(f[y]))
		builder.WriteString("\n")
	}
	return builder.String()
}

func (d day14) Part1(reader io.Reader) int {
	platform := parsePlatform(reader)
	platform.Tilt(-1, 0)
	return platform.WeightLoad()
}

func (d day14) Part2(reader io.Reader) int {
	platform := parsePlatform(reader)

	seen := map[string]int{
		platform.String(): 0,
	}
	index := map[int]string{
		0: platform.String(),
	}

	// Cycle detection: determine if we've seen this platform configuration
	// before using the `seen` map and then use the period and start of
	// the cycle to deduce the step within the cycle that 1000000000
	// iterations would be.
	iteration := 0
	for {
		iteration += 1
		platform.Tilt(-1, 0) // North
		platform.Tilt(0, -1) // West
		platform.Tilt(1, 0)  // South
		platform.Tilt(0, 1)  // East

		str := platform.String()
		first, already := seen[str]
		if already {
			var period = iteration - first
			var cycleStep = first + ((1000000000 - first) % period)

			// Rehydrate the platform from the string key to determine weight load
			return parsePlatform(strings.NewReader(index[cycleStep])).WeightLoad()
		}
		seen[str] = iteration
		index[iteration] = str
	}
}
