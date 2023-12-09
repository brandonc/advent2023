package day09

import (
	"bufio"
	"io"
	"strings"

	"github.com/brandonc/advent2023/internal/input"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day09 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day09{}
}

func diffs(values []int) []int {
	result := make([]int, len(values)-1)
	for a, b := 0, 1; b < len(values); a, b = a+1, b+1 {
		result[a] = values[b] - values[a]
	}
	return result
}

func all(history []int, value int) bool {
	for _, n := range history {
		if n != value {
			return false
		}
	}
	return true
}

func predictNext(history []int) int {
	if all(history, 0) {
		return 0
	}

	increaseBy := predictNext(diffs(history))
	return history[len(history)-1] + increaseBy
}

func predictPrevious(history []int) int {
	if all(history, 0) {
		return 0
	}

	decreaseBy := predictPrevious(diffs(history))
	return history[0] - decreaseBy
}

func parseInts(line string) []int {
	intScanner := input.NewIntScanner(strings.NewReader(line))
	intScanner.Split(bufio.ScanWords)

	result := make([]int, 0, 32)
	for intScanner.Scan() {
		result = append(result, intScanner.Int())
	}
	return result
}

func (d day09) Part1(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		sum += predictNext(parseInts(scanner.Text()))
	}
	return sum
}

func (d day09) Part2(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		sum += predictPrevious(parseInts(scanner.Text()))
	}
	return sum
}
