package day15

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day15 struct{}

func Factory() solution.Solver {
	return day15{}
}

func hash(input string) int {
	hash := 0
	for i := 0; i < len(input); i++ {
		hash += int(input[i])
		hash *= 17
		hash %= 256
	}
	return hash
}

func (d day15) Part1(reader io.Reader) int {
	result := 0
	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		ui.Die(errors.New("No input"))
	}

	for _, s := range strings.Split(scanner.Text(), ",") {
		result += hash(s)
	}

	return result
}

type Pair struct {
	Label string
	Value int
}

type HashMap struct {
	Boxes [256][]Pair
}

func (h *HashMap) Set(label string, value int) {
	box := hash(label)
	all := h.Boxes[box]
	for i := 0; i < len(all); i++ {
		if h.Boxes[box][i].Label == label {
			h.Boxes[box][i].Value = value
			return
		}
	}
	h.Boxes[box] = append(h.Boxes[box], Pair{label, value})
}

func (h *HashMap) Unset(label string) {
	box := hash(label)
	slots := h.Boxes[box]
	for i := 0; i < len(slots); i++ {
		if slots[i].Label == label {
			h.Boxes[box] = append(slots[:i], slots[i+1:]...)
			return
		}
	}
}

func (h HashMap) FocusingPower() int {
	total := 0
	for box := 0; box < 256; box++ {
		for slot := 0; slot < len(h.Boxes[box]); slot++ {
			score := (box + 1) * (slot + 1) * h.Boxes[box][slot].Value
			ui.Debugf(
				"%s: %d (box %d) * %d (slot %d) * %d (focal length) = %d",
				h.Boxes[box][slot].Label, (box + 1), box, slot+1, slot, h.Boxes[box][slot].Value, score,
			)
			total += score
		}
	}
	return total
}

func (d day15) Part2(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		ui.Die(errors.New("No input"))
	}

	hm := HashMap{}
	for _, s := range strings.Split(scanner.Text(), ",") {
		if strings.HasSuffix(s, "-") {
			hm.Unset(s[:len(s)-1])
		} else {
			fields := strings.Split(s, "=")
			val, err := strconv.Atoi(fields[1])
			ui.Die(err)
			hm.Set(fields[0], val)
		}
	}
	return hm.FocusingPower()
}
