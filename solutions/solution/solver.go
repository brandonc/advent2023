package solution

import "io"

type Solver interface {
	Part1(input io.Reader) int
	Part2(input io.Reader) int
}

type SolutionFactory func() Solver
