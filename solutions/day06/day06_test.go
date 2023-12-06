package day06

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, a2 := Factory().Solve(strings.NewReader(`Time:      7  15   30
	Distance:  9  40  200`))

	if expected := 288; a1 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a1)
	}

	if expected := 71503; a2 != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, a2)
	}
}
