package day13

import (
	"strings"
	"testing"
)

var (
	sample1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	sample2 = `###.##.###...####
.#.####.#..##..#.
##..##..########.
##......########.
.##.##.##..##..##
..##..##...##...#
##......###..###.`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 405; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample1))

	if expected := 400; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}

func TestPart2B(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample2))

	if expected := 12; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
