package day24

import (
	"strings"
	"testing"
)

var (
	sample1 = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`
)

func TestPart1(t *testing.T) {
	answer := day24{
		testAreaMin: 7,
		testAreaMax: 27,
	}.Part1(strings.NewReader(sample1))

	if expected := 2; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample1))

	if expected := 0; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
