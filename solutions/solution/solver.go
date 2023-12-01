package solution

import "io"

type Solver interface {
	Solve(input io.Reader) (any, any)
}

type SolutionFactory func() Solver
