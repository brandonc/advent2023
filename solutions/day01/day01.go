package day01

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day01 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day01{}
}

func substrAt(index int, needle, haystack string) bool {
	if index+len(needle) > len(haystack) {
		// Not possible to fit needle at index within haystack
		return false
	}

	for i := index; i < index+len(needle); i++ {
		if haystack[i] != needle[i-index] {
			return false
		}
	}

	return true
}

var numbersSpelled = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
	"zero":  '0',
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		var first, last byte

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = line[i]
				break
			}
		}

		// The test data for part 2 doesn't work for part 1
		if first == 0 || last == 0 {
			continue
		}

		sum += (int(first-'0') * 10) + int(last-'0')
	}

	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		var first, last byte

		// Look for first value from beginning
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = line[i]
				break
			}

			for spelled, symbol := range numbersSpelled {
				if substrAt(i, spelled, line) {
					first = symbol
					break
				}
			}

			if first != 0 {
				break
			}
		}

		// Look for last value from end
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = line[i]
				break
			}

			for spelled, symbol := range numbersSpelled {
				if substrAt(i, spelled, line) {
					last = symbol
					break
				}
			}

			if last != 0 {
				break
			}
		}

		sum += (int(first-'0') * 10) + int(last-'0')
	}

	return sum
}

func (d day01) Solve(reader io.Reader) (any, any) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return part1(lines), part2(lines)
}
