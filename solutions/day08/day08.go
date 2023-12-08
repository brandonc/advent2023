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

func (d day08) Part1(instructions string, nodes map[string]*Node) int {
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

func (d day08) Part2(instructions string, nodes map[string]*Node) int {
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

func (d day08) Solve(reader io.Reader) (any, any) {
	scanner := bufio.NewScanner(reader)

	nodes := make(map[string]*Node)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan() // Blank link

	for scanner.Scan() {
		line := scanner.Text()

		thisLabel := line[0:3]
		leftNodeLabel := line[7:10]
		rightNodeLabel := line[12:15]

		thisNode, hasSelf := nodes[thisLabel]

		if !hasSelf {
			thisNode = &Node{
				Label: thisLabel,
			}
		}

		leftNode, hasLeft := nodes[leftNodeLabel]
		if !hasLeft {
			leftNode = &Node{
				Label: leftNodeLabel,
			}
			nodes[leftNodeLabel] = leftNode
		}

		rightNode, hasRight := nodes[rightNodeLabel]
		if !hasRight {
			rightNode = &Node{
				Label: rightNodeLabel,
			}
			nodes[rightNodeLabel] = rightNode
		}

		thisNode.Left = leftNode
		thisNode.Right = rightNode

		nodes[thisNode.Label] = thisNode
	}

	return d.Part1(instructions, nodes), d.Part2(instructions, nodes)
}
