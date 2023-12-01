// package day01 is a substring problem with the objective of
// finding digits within lines of text.
//
// For part 1, the first and last digits from each line are
// concatenated, converted to an integer, and added to a sum.
//
// Part 2 adds the possibility of spelled numbers "zero" through
// "nine" which are concatenated value-wise and added to a sum
// as before.
package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/brandonc/advent2023/internal/ui"
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

func combineDigits(a, b byte) int {
	number := fmt.Sprintf("%c%c", a, b)
	amount, err := strconv.Atoi(number)
	ui.Assert(err == nil, fmt.Sprintf("Cannot convert %q to int", number))
	return amount
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

		sum += combineDigits(first, last)
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

		sum += combineDigits(first, last)
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
