package day14

import (
	"strings"
	"testing"
)

var (
	sample1 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 136; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample1))

	if expected := 64; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
