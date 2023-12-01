package day01

import (
	"strings"
	"testing"
)

func TestSampleInput(t *testing.T) {
	a1, _ := Factory().Solve(strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`))

	_, a2 := Factory().Solve(strings.NewReader(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`))

	if expected := 142; a1 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a1)
	}

	if expected := 281; a2 != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, a2)
	}
}
