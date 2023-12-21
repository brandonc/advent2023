package day21

import (
	"strings"
	"testing"
)

var (
	sample1 = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`
)

func TestPart1(t *testing.T) {
	answer := day21{testSteps: 6}.Part1(strings.NewReader(sample1))

	if expected := 16; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	tc := []struct{ steps, expected int }{
		{50, 1594},
		{100, 6536},
		{500, 167004},
	}

	for _, c := range tc {
		answer := day21{testSteps: c.steps}.Part2(strings.NewReader(sample1))

		if answer != c.expected {
			t.Fatalf("Expected %d steps to be %d, got %d", c.steps, c.expected, answer)
		}
	}
}
