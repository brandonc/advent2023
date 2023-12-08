package day03

import (
	"strings"
	"testing"
)

var sample = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var custom = `..3*...
....827`

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample))

	if expected := 4361; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample))

	if expected := 467835; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}

func TestCustomPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(custom))

	if expected := 2481; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
