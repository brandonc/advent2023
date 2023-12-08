package ui

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/brandonc/advent2023/internal/maths"
	"github.com/mitchellh/colorstring"
)

func Die(err error) {
	if err != nil {
		colorstring.Printf("[red]An unexpected error occurred:\n%s[reset]\n", err)
		os.Exit(1)
	}
}

func Assert(expr bool, description string) {
	if !expr {
		Die(errors.New(description))
		os.Exit(1)
	}
}

func Debug(message string) {
	if os.Getenv("LOG_LEVEL") != "debug" {
		return
	}
	colorstring.Printf("[dark_gray][DEBUG] %s\n", message)
}

func Debugf(message string, a ...any) {
	Debug(fmt.Sprintf(message, a...))
}

func rightAlign(v string, other ...string) string {
	maxOther := 0
	for _, o := range other {
		if len(o) > maxOther {
			maxOther = len(o)
		}
	}
	if len(v) > maxOther {
		return v
	} else {
		return fmt.Sprintf("%s%s", strings.Repeat(" ", maxOther-len(v)), v)
	}
}

func humanizeDuration(d time.Duration) string {
	units := []struct {
		unit   string
		amount int
	}{
		{"s", int(d.Seconds())},
		{"ms", int(d.Milliseconds())},
		{"μs", int(d.Microseconds())},
		{"ns", int(d.Nanoseconds())},
	}

	for _, u := range units {
		if u.amount == 0 {
			continue
		}

		return fmt.Sprintf("%d%s", u.amount, u.unit)
	}

	return "0 ns"
}

func Answer(part1 func() int, part2 func() int) {
	startTimePart1 := time.Now()
	answerPart1 := fmt.Sprintf("%d", part1())
	timePart1 := time.Since(startTimePart1)

	startTimePart2 := time.Now()
	answerPart2 := fmt.Sprintf("%d", part2())
	timePart2 := time.Since(startTimePart2)

	dashes := strings.Repeat("-", maths.MaxInt(len(answerPart1), len(answerPart2))+2+len("Part X / "))

	colorstring.Printf("[yellow]+%s+\n", dashes)
	colorstring.Printf("[yellow]| [cyan]Part 1 / [white]%s [yellow]| [dark_gray]%s\n", rightAlign(answerPart1, answerPart2), humanizeDuration(timePart1))
	colorstring.Printf("[yellow]| [cyan]Part 2 / [white]%s [yellow]| [dark_gray]%s\n", rightAlign(answerPart2, answerPart1), humanizeDuration(timePart2))
	colorstring.Printf("[yellow]+%s+\n", dashes)

	// +-------------------------+
	// | Part 1 / 54561213452435 |
	// | Part 2 /          54076 |
	// +-----------=-------------+

	colorstring.Printf("[dark_gray]%s\n", humanizeDuration(timePart1+timePart2))
}
