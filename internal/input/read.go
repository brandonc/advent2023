package input

import (
	"bufio"
	"errors"
	"io"
	"strconv"

	"github.com/brandonc/advent2023/internal/ui"
)

type IntScanner struct {
	*bufio.Scanner
}

func NewIntScanner(reader io.Reader) IntScanner {
	return IntScanner{
		bufio.NewScanner(reader),
	}
}

func (i IntScanner) Int() int {
	item, err := strconv.Atoi(i.Text())
	if err != nil {
		ui.Die(errors.New("expected int"))
	}
	return item
}
