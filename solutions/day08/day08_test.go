package day08

import (
	"strings"
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	a1, _ := Factory().Solve(strings.NewReader(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`))

	if expected := 2; a1 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a1)
	}
}

func TestPart1Sample2(t *testing.T) {
	a1, _ := Factory().Solve(strings.NewReader(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`))

	if expected := 6; a1 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a1)
	}
}

func TestPart2Sample1(t *testing.T) {
	_, a2 := Factory().Solve(strings.NewReader(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`))

	if expected := 6; a2 != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, a2)
	}
}
