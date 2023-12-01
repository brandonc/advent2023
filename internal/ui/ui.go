package ui

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func rightAlign(v, other string) string {
	if len(v) > len(other) {
		return v
	} else {
		return fmt.Sprintf("%s%s", strings.Repeat(" ", len(other)-len(v)), v)
	}
}

func answerString(first, second string) {
	dashes := strings.Repeat("-", maths.Max(len(first), len(second))+2+len("Part X / "))

	colorstring.Printf("[yellow]+%s+\n", dashes)
	colorstring.Printf("[yellow]| [cyan]Part 1 / [white]%s [yellow]|\n", rightAlign(first, second))
	colorstring.Printf("[yellow]| [cyan]Part 2 / [white]%s [yellow]|\n", rightAlign(second, first))
	colorstring.Printf("[yellow]+%s+\n", dashes)

	// +-------------------------+
	// | Part 1 / 54561213452435 |
	// | Part 2 /          54076 |
	// +-----------=-------------+
}

func answerInt(first, second int) {
	a1 := strconv.FormatInt(int64(first), 10)
	a2 := strconv.FormatInt(int64(second), 10)

	answerString(a1, a2)
}

func Answer(first, second any) {
	switch first.(type) {
	case int, int64:
		answerInt(first.(int), second.(int))
	case string:
		answerString(first.(string), second.(string))
	}
}
