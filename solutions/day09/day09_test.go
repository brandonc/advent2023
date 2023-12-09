package day09

import (
	"strings"
	"testing"
)

var (
	sample1 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 114; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample1))

	if expected := 2; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
