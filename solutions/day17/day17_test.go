package day17

import (
	"strings"
	"testing"
)

var (
	sample1 = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

	sample2 = `111111111111
999999999991
999999999991
999999999991
999999999991`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 102; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample1))

	if expected := 94; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}

func TestPart2b(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample2))

	if expected := 71; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
