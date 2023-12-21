package day20

import (
	"strings"
	"testing"
)

var (
	sample1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

	sample2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 32000000; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample2))

	if expected := 11687500; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}
