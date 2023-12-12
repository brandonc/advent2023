package day12

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day12 struct {
}

type Puzzle struct {
	Record     []byte
	Contiguous string
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day12{}
}

var cache = make(map[string]int, 4096)

// A memoized version of numWorkingArrangementsInternal
func numWorkingArrangements(line []byte, groups string) int {
	key := fmt.Sprintf("%s|%s", line, groups)
	if val, ok := cache[key]; ok {
		return val
	}
	actual := numWorkingArrangementsInternal(line, groups)
	cache[key] = actual
	return actual
}

func numWorkingArrangementsInternal(line []byte, groups string) int {
	// Base case: either the line and the groups are exhausted, in which case
	// the arrangement is counted, or there are some group(s) left, in which
	// case it's not a match.
	if len(line) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	// If there are no groups left, ensure all remaining records are not '#'
	if len(groups) == 0 {
		if bytes.Contains(line, []byte{'#'}) {
			return 0
		}
		return 1
	}

	switch line[0] {
	case '.':
		// Ignore '.'
		return numWorkingArrangements(line[1:], groups)

	case '?':
		// This is the '?' permutation case. Try both '.' and '#' for matches
		return numWorkingArrangements(append([]byte{'#'}, line[1:]...), groups) +
			numWorkingArrangements(append([]byte{'.'}, line[1:]...), groups)

	default: // '#': do next group eval
		groupsSplit := strings.SplitN(groups, ",", 2)
		group, _ := strconv.Atoi(groupsSplit[0])
		for i := 0; i < group; i++ {
			if line[i] == '.' {
				return 0
			}
		}

		// Adjacent '#' indicates too many springs for this group
		if line[group] == '#' {
			return 0
		}

		if len(groupsSplit) == 1 {
			// This just ensures that the final group is an empty string
			// if SplitN returned 1 substring
			groupsSplit = append(groupsSplit, "")
		}

		return numWorkingArrangements(line[group+1:], groupsSplit[1])
	}
}

func (p *Puzzle) Unfold() *Puzzle {
	return &Puzzle{
		Record:     bytes.Join([][]byte{p.Record, p.Record, p.Record, p.Record, p.Record}, []byte{'?'}),
		Contiguous: strings.Join([]string{p.Contiguous, p.Contiguous, p.Contiguous, p.Contiguous, p.Contiguous}, ","),
	}
}

func parse(reader io.Reader) []Puzzle {
	scanner := bufio.NewScanner(reader)

	result := make([]Puzzle, 0)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")

		puzzle := Puzzle{
			Record:     []byte(fields[0]),
			Contiguous: fields[1],
		}
		result = append(result, puzzle)
	}

	return result
}

func (d day12) Part1(reader io.Reader) int {
	sum := 0
	puzzles := parse(reader)
	for _, p := range puzzles {
		// Appends a . at the end of each line to simplify final group matching
		sum += numWorkingArrangements([]byte(string(p.Record)+"."), p.Contiguous)
	}
	return sum
}

func (d day12) Part2(reader io.Reader) int {
	sum := 0
	puzzles := parse(reader)
	for _, p := range puzzles {
		unfolded := p.Unfold()
		// Appends a . at the end of each line to simplify final group matching
		sum += numWorkingArrangements([]byte(string(unfolded.Record)+"."), unfolded.Contiguous)
	}
	return sum
}
