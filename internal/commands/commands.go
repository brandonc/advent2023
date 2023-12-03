// Code generated by go generate; DO NOT EDIT.
package commands

import (
	"github.com/brandonc/advent2023/solutions/day01"
	"github.com/brandonc/advent2023/solutions/day02"
	"github.com/brandonc/advent2023/solutions/day03"
	"github.com/brandonc/advent2023/solutions/solution"
)

var SolutionCommands = map[string]solution.SolutionFactory{
	"1": day01.Factory,
	"2": day02.Factory,
	"3": day03.Factory,
}
