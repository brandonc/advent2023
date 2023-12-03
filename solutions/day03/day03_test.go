package day03

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2 := Factory().Solve(strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`))

	if expected := 4361; a1 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a1)
	}

	if expected := 467835; a2 != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, a2)
	}
}

func TestCustomInput(t *testing.T) {
	_, a2 := Factory().Solve(strings.NewReader(`..3*...
....827`))

	if expected := 2481; a2 != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, a2)
	}
}
