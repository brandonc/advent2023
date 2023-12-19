package day19

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day19 struct{}

func Factory() solution.Solver {
	return day19{}
}

type Condition struct {
	Attribute byte
	Operation byte
	Value     int
	Target    string
}

type Workflow struct {
	Conditions []Condition
	Else       string
}

type Workflows map[string]Workflow

type Part struct {
	Attributes map[byte]int
}

func parseWorkflows(scanner *bufio.Scanner) Workflows {
	result := make(Workflows)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var name string
		var symbol string
		workflow := Workflow{}
		current := Condition{}
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '{':
				// End of workflow name
				name = symbol
				symbol = ""
			case '<', '>':
				current.Operation = line[i]
				current.Attribute = symbol[0]
				symbol = ""
			case ',':
				// Next condition
				current.Target = symbol
				workflow.Conditions = append(workflow.Conditions, current)
				current = Condition{}
				symbol = ""
			case ':':
				// End of value
				val, err := strconv.Atoi(symbol)
				ui.Die(err)
				current.Value = val
				symbol = ""
			case '}':
				// End of workflow
				workflow.Else = symbol
			default:
				symbol += string(line[i])
			}
			result[name] = workflow
		}
	}

	return result
}

func parseParts(scanner *bufio.Scanner) []Part {
	result := make([]Part, 0)

	for scanner.Scan() {
		current := Part{
			Attributes: make(map[byte]int),
		}
		line := scanner.Text()
		// specimen: {x=787,m=2655,a=1222,s=2876}
		pairs := strings.Split(strings.Trim(line, "{}"), ",")

		for _, pair := range pairs {
			attribute := pair[0]
			value, err := strconv.Atoi(pair[2:])

			current.Attributes[attribute] = value
			ui.Die(err)
		}

		result = append(result, current)
	}

	return result
}

func (w Workflows) Execute(name string, part Part) bool {
	if name == "A" {
		return true
	} else if name == "R" {
		return false
	}

	workflow := w[name]

	for _, condition := range workflow.Conditions {
		partValue := part.Attributes[condition.Attribute]
		if (condition.Operation == '<' && partValue < condition.Value) ||
			(condition.Operation == '>' && partValue > condition.Value) {
			return w.Execute(condition.Target, part)
		}
	}
	return w.Execute(workflow.Else, part)
}

func (w Workflows) ExecuteDistinct(name string, bounds map[byte][2]int) int {
	if name == "A" {
		// Return the product of all the remaining ranges.
		products := 1
		for _, bound := range bounds {
			products *= bound[1] - bound[0] + 1 // (inclusive)
		}
		return products
	}
	if name == "R" {
		return 0
	}

	workflow := w[name]

	sum := 0
	for _, condition := range workflow.Conditions {
		// Copy the bounds map for the true side of the condition. The argument
		// bounds will be used for the false side.
		trueBounds := make(map[byte][2]int)
		for k, v := range bounds {
			trueBounds[k] = v
		}

		if condition.Operation == '<' {
			// True case: [lower, condition.Value - 1]
			trueBounds[condition.Attribute] = [2]int{bounds[condition.Attribute][0], condition.Value - 1}
			sum += w.ExecuteDistinct(condition.Target, trueBounds)

			// False case: [condition.Value, upper]
			bounds[condition.Attribute] = [2]int{condition.Value, bounds[condition.Attribute][1]}
		} else {
			// True case: [condition.Value + 1, upper]
			trueBounds[condition.Attribute] = [2]int{condition.Value + 1, bounds[condition.Attribute][1]}
			sum += w.ExecuteDistinct(condition.Target, trueBounds)

			// False case: [lower, condition.Value]
			bounds[condition.Attribute] = [2]int{bounds[condition.Attribute][0], condition.Value}
		}
	}

	// Finally, execute the else case with the remaining 'false' bounds.
	sum += w.ExecuteDistinct(workflow.Else, bounds)
	return sum
}

func (d day19) Part1(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	workflows := parseWorkflows(scanner)
	parts := parseParts(scanner)

	sum := 0
	for _, p := range parts {
		if workflows.Execute("in", p) {
			sum += p.Attributes['x'] + p.Attributes['m'] + p.Attributes['a'] + p.Attributes['s']
		}
	}

	return sum
}

func (d day19) Part2(reader io.Reader) int {
	// Figure out how many distinct combinations of x, m, a, s attributes will end
	// up being accepted by the workflows. I'll start with a full possible
	// range for each attributes, and narrow each side of the range based on
	// the condition encountered, recursively for each true condition plus
	// the fallback condition.
	//
	// If 'A' is ultimately reached, the partial answer will be the product of all
	// the differences of each attribute's remaining range.
	// If 'R' is reached, the partial answer will be 0.
	//
	// The final answer will be the sum of all partial answers.

	scanner := bufio.NewScanner(reader)
	workflows := parseWorkflows(scanner)

	return workflows.ExecuteDistinct("in", map[byte][2]int{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	})
}
