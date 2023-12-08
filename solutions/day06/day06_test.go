package day06

import (
	"strings"
	"testing"
)

var sample = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample))

	if expected := 288; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample))

	if expected := 71503; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
