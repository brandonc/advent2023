package day07

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2 := Factory().Solve(strings.NewReader(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`))

	if expected := 6440; a1 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a1)
	}

	if expected := 5905; a2 != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, a2)
	}
}
