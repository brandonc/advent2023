package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/brandonc/advent2023/internal/commands"
	"github.com/brandonc/advent2023/internal/ui"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	var input io.Reader
	var err error

	if len(os.Args) > 2 {
		input, err = os.Open(os.Args[2])
		ui.Die(err)
	} else {
		input = os.Stdin
	}

	solutionFactory, ok := commands.SolutionCommands[os.Args[1]]
	if !ok {
		printUsage()
	}

	var buf bytes.Buffer
	tee := io.TeeReader(input, &buf)
	solution := solutionFactory()

	ui.Answer(
		func() int { return solution.Part1(tee) },
		func() int { return solution.Part2(&buf) },
	)
}

func printUsage() {
	fmt.Println("Usage:", os.Args[0], "<day> [input]")
	os.Exit(127)
}
