package day08

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2023/internal/input"
	"github.com/brandonc/advent2023/internal/maths"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day08 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day08{}
}

type Node struct {
	Label string
	Left  *Node
	Right *Node
}

func (n Node) Navigate(direction byte) *Node {
	if direction == 'L' {
		return n.Left
	}
	return n.Right
}

func (n *Node) init(nodes map[string]*Node, leftLabel, rightLabel string) {
	left, hasLeft := nodes[leftLabel]
	if !hasLeft {
		left = &Node{
			Label: leftLabel,
		}
		nodes[leftLabel] = left
	}

	right, hasLeft := nodes[rightLabel]
	if !hasLeft {
		right = &Node{
			Label: rightLabel,
		}
		nodes[rightLabel] = right
	}

	n.Left = left
	n.Right = right
	nodes[n.Label] = n
}

func parseInput(reader io.Reader) (string, map[string]*Node) {
	scanner := bufio.NewScanner(reader)

	nodes := make(map[string]*Node)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan() // Blank link

	for scanner.Scan() {
		line := scanner.Text()

		thisLabel := line[0:3]
		thisNode, hasSelf := nodes[thisLabel]

		if !hasSelf {
			thisNode = &Node{
				Label: thisLabel,
			}
		}

		thisNode.init(nodes, line[7:10], line[12:15])
	}

	return instructions, nodes
}

func (d day08) Part1(reader io.Reader) int {
	instructions, nodes := parseInput(reader)
	node, ok := nodes["AAA"]
	if !ok {
		return 0
	}

	buf := input.NewRingBuffer(instructions)

	result := 0
	for node.Label != "ZZZ" {
		result += 1
		node = node.Navigate(buf.Next())
	}
	return result
}

func (d day08) Part2(reader io.Reader) int {
	instructions, nodes := parseInput(reader)

	result := 1
	for label := range nodes {
		if label[2] != 'A' {
			continue
		}

		buf := input.NewRingBuffer(instructions)

		steps := 0
		current := nodes[label]
		for current.Label[2] != 'Z' {
			steps += 1
			current = current.Navigate(buf.Next())
		}

		result = maths.LCM(result, steps)
	}

	return result
}
