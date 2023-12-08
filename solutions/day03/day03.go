package day03

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day03 struct {
	grid [][]byte
}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day03{
		grid: make([][]byte, 0),
	}
}

// symbolAdjacent determines if a non-digit, non-period character appears
// immediately adjacent to the characters on specifed row at the specified
// start/end positions.
func (d *day03) symbolAdjacent(row, start, end int) bool {
	for y := row - 1; y <= row+1; y++ {
		for x := start - 1; x < end+1; x++ {
			if y < 0 || y >= len(d.grid) || x < 0 || x >= len(d.grid[0]) {
				// Out of bounds
				continue
			}

			if !isDigit(d.grid[y][x]) && d.grid[y][x] != '.' {
				ui.Debugf("Found symbol %q at %d, %d", d.grid[y][x], y, x)
				return true
			}
		}
	}
	return false
}

func (d day03) part1() int {
	var answer = 0

	// Loop through each byte in the grid. If a number is found, parse it and determine if any
	// symbols are adjacent to it. If so, add the parsed number to the answer.
	for y := 0; y < len(d.grid); y++ {
		var startNum = -1
		for x := 0; x < len(d.grid[0]); x++ {
			c := d.grid[y][x]
			if isDigit(c) {
				if startNum == -1 {
					startNum = x
				}
			} else if startNum > -1 {
				// Indicates that the end of the number was found before the end of the line
				if d.symbolAdjacent(y, startNum, x) {
					answer += d.parseNumberAt(y, startNum)
				}
				startNum = -1
			}
		}
		if startNum > -1 {
			// Indicates that the end of the line was found before the end of the number
			x := len(d.grid[0])
			if d.symbolAdjacent(y, startNum, x) {
				answer += d.parseNumberAt(y, startNum)
			}
		}
	}
	return answer
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (d day03) parseNumberAt(row, col int) int {
	ui.Assert(isDigit(d.grid[row][col]), fmt.Sprintf("There was no digit at %d, %d", row, col))

	begin, end := col-1, col+1
	for begin >= 0 && isDigit(d.grid[row][begin]) {
		begin--
	}

	for end < len(d.grid[0]) && isDigit(d.grid[row][end]) {
		end++
	}

	num, err := strconv.Atoi(string(d.grid[row][begin+1 : end]))
	ui.Assert(err == nil, fmt.Sprintf("Couldn't parse number %q", d.grid[row][begin+1:end]))

	return num
}

// This hairy function parses numbers that appear anywhere around a grid location.
// The multitude of conditions reflect the fact that some digits above or below
// the row are connected to the same number so they should not be parsed twice.
func (d day03) adjacentNumbers(row, col int) []int {
	var nums = make([]int, 0)
	var columns = len(d.grid[0])
	var rows = len(d.grid)
	var hasNW, hasSW = false, false

	// Look NW
	if row >= 1 {
		if col >= 1 && isDigit(d.grid[row-1][col-1]) {
			nums = append(nums, d.parseNumberAt(row-1, col-1))
			hasNW = true
		}

		// Look N
		if isDigit(d.grid[row-1][col]) {
			if !hasNW {
				nums = append(nums, d.parseNumberAt(row-1, col))
			}
		} else if d.grid[row-1][col] == '.' && col < columns-1 && isDigit(d.grid[row-1][col+1]) {
			// Nothing N so look NE
			nums = append(nums, d.parseNumberAt(row-1, col+1))
		}
	}

	// Look E
	if col >= 1 && isDigit(d.grid[row][col-1]) {
		nums = append(nums, d.parseNumberAt(row, col-1))
	}

	// Look W
	if col < columns-1 && isDigit(d.grid[row][col+1]) {
		nums = append(nums, d.parseNumberAt(row, col+1))
	}

	// Look SW
	if row < rows {
		if col >= 1 && isDigit(d.grid[row+1][col-1]) {
			nums = append(nums, d.parseNumberAt(row+1, col-1))
			hasSW = true
		}

		// Look S
		if isDigit(d.grid[row+1][col]) {
			if !hasSW {
				nums = append(nums, d.parseNumberAt(row+1, col))
			}
		} else if d.grid[row+1][col] == '.' && col < columns-1 && isDigit(d.grid[row+1][col+1]) {
			// Nothing S so look SE
			nums = append(nums, d.parseNumberAt(row+1, col+1))
		}
	}

	return nums
}

func (d day03) part2() int {
	var sum = 0
	for y := 0; y < len(d.grid); y++ {
		for x := 0; x < len(d.grid[0]); x++ {
			c := d.grid[y][x]
			if c == '*' {
				ui.Debugf("Found gear '*' at %d, %d", y, x)
				nums := d.adjacentNumbers(y, x)
				ui.Debugf("The adjacent numbers are %+v", nums)
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}

	return sum
}

func (d *day03) parseGrid(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		d.grid = append(d.grid, []byte(line))
	}
}

func (d day03) Part1(reader io.Reader) int {
	d.parseGrid(reader)

	return d.part1()
}

func (d day03) Part2(reader io.Reader) int {
	d.parseGrid(reader)

	return d.part2()
}
