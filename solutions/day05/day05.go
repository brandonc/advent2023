package day05

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/brandonc/advent2023/internal/input"
	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

var categories = []string{
	"soil", "fertilizer", "water", "light", "temperature", "humidity", "location",
}

type day05 struct {
	almanac map[string][]Mapping
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day05{
		almanac: make(map[string][]Mapping),
	}
}

type Interval struct {
}

type Mapping struct {
	SourceStart int
	Length      int
	DestStart   int
}

func (m Mapping) Contains(source int) bool {
	return source >= m.SourceStart && source <= m.SourceStart+m.Length
}

func (m Mapping) Destination(source int) int {
	return m.DestStart + (source - m.SourceStart)
}

func scanInts(s string) []int {
	scanner := input.NewIntScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	result := make([]int, 0)
	for scanner.Scan() {
		result = append(result, scanner.Int())
	}
	return result
}

func (d day05) findLocationForSeed(seed int) int {
	source := seed
	for _, cat := range categories {
		for _, mapping := range d.almanac[cat] {
			if mapping.Contains(source) {
				dest := mapping.Destination(source)
				source = dest
				break
			}
		}
	}

	return source
}

func (d day05) part1(seeds []int) int {
	min := math.MaxInt
	for _, seed := range seeds {
		if loc := d.findLocationForSeed(seed); loc < min {
			min = loc
		}
	}

	return min
}

func (d day05) part2(seeds []int) int {
	min := math.MaxInt

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		ui.Debugf("Beginning range %d ~ %d (%d seeds)", start, start+length, start+length-start)

		for seed := start; seed < start+length; seed++ {
			if loc := d.findLocationForSeed(seed); loc < min {
				min = loc
			}
		}
	}

	return min
}

func (d day05) parseAlmanac(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasSuffix(line, "map:") {
			ui.Die(fmt.Errorf("expected map labels line but got %q", line))
		}
		mapLabels := strings.Split(scanner.Text()[0:len(line)-len(" map:")], "-to-")
		ui.Assert(len(mapLabels) == 2, fmt.Sprintf("Unexpected map labels input %q", line))

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}

			ranges := scanInts(line)
			ui.Assert(len(ranges) == 3, fmt.Sprintf("Unexpected ranges size %d: %q", len(ranges), line))

			_, ok := d.almanac[mapLabels[1]]
			if !ok {
				d.almanac[mapLabels[1]] = make([]Mapping, 0)
			}

			d.almanac[mapLabels[1]] = append(d.almanac[mapLabels[1]], Mapping{
				SourceStart: ranges[1],
				Length:      ranges[2],
				DestStart:   ranges[0],
			})
		}
	}
}

func mergeSeedRanges(seeds []int) []int {

}

func (d day05) Solve(reader io.Reader) (any, any) {
	scanner := bufio.NewScanner(reader)
	var seeds []int
	if scanner.Scan() && strings.HasPrefix(scanner.Text(), "seeds: ") {
		seeds = scanInts(scanner.Text()[len("seeds: "):])
		scanner.Scan() // Eat the next blank line
	} else {
		ui.Die(errors.New("first line did not contain seeds"))
	}

	d.parseAlmanac(scanner)

	return d.part1(seeds), d.part2(mergeSeedRanges(seeds))
}
