package day25

import (
	"bufio"
	"io"
	"slices"
	"strings"

	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day25 struct{}

func Factory() solution.Solver {
	return day25{}
}

type Edge struct {
	A string
	B string
}

type set map[string]struct{}

type Graph struct {
	Edges map[string]set
}

func parseGraph(reader io.Reader) Graph {
	scanner := bufio.NewScanner(reader)
	result := Graph{
		Edges: make(map[string]set),
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ": ")

		from := fields[0]
		for _, to := range strings.Split(fields[1], " ") {
			childrenFrom := result.Edges[from]
			childrenTo := result.Edges[to]

			if childrenFrom == nil {
				childrenFrom = make(set)
			}
			if childrenTo == nil {
				childrenTo = make(set)
			}

			childrenFrom[to] = struct{}{}
			childrenTo[from] = struct{}{}

			result.Edges[from] = childrenFrom
			result.Edges[to] = childrenTo
		}
	}
	return result
}

type routes map[string][]string
type bestRoutes map[string]routes
type frontier struct {
	node string
	path []string
}

type RankedEdge struct {
	Edge
	Rank int
}

func (b bestRoutes) Rank() []RankedEdge {
	edges := make(map[Edge]int)
	result := make([]RankedEdge, 0)
	for _, routes := range b {
		for _, path := range routes {
			for i := 0; i < len(path)-1; i++ {
				a := path[i]
				b := path[i+1]
				if a > b {
					a, b = b, a
				}
				edges[Edge{a, b}]++
			}
		}
	}

	for edge, rank := range edges {
		result = append(result, RankedEdge{edge, rank})
	}

	slices.SortFunc[[]RankedEdge](result, func(i, j RankedEdge) int {
		if i.Rank == j.Rank {
			return 0
		}
		if i.Rank < j.Rank {
			return 1
		}
		return -1
	})

	return result
}

func (g Graph) findRoutesToOtherNodes(start string) map[string][]string {
	// Find shortest route between all points

	// Map of node -> list of nodes to visit
	queue := []frontier{{start, []string{}}}
	result := make(routes)

	for len(queue) > 0 {
		next, path := queue[0].node, queue[0].path
		queue = queue[1:]

		// If we've already visited this node, skip it
		if _, ok := result[next]; ok {
			continue
		}

		path = append(path, next)
		if len(path) > 1 {
			result[next] = path
		}

		for child := range g.Edges[next] {
			queue = append(queue, frontier{child, path})
		}
	}

	return result
}

func (g Graph) CountReacheable(start string, ignore map[Edge]struct{}) int {
	// Map of node -> list of nodes to visit
	queue := []frontier{{start, []string{}}}
	reached := make(routes)

	for len(queue) > 0 {
		next := queue[0].node
		queue = queue[1:]

		// If we've already visited this node, skip it
		if _, ok := reached[next]; ok {
			continue
		}

		reached[next] = []string{}

		for child := range g.Edges[next] {
			a := next
			b := child

			if a > b {
				a, b = b, a
			}

			if _, ok := ignore[Edge{a, b}]; ok {
				continue
			}

			queue = append(queue, frontier{child, []string{}})
		}
	}

	return len(reached)
}

func (d day25) Part1(reader io.Reader) int {
	graph := parseGraph(reader)

	// Find shortest route between all points
	allRoutes := make(bestRoutes)
	for start := range graph.Edges {
		allRoutes[start] = graph.findRoutesToOtherNodes(start)
		ui.Debugf("Best routes from %q:", start)
		for end, path := range allRoutes[start] {
			ui.Debugf("  %s: %v", end, path)
		}
	}

	// Sort edges according to the # of times they appear in the list of shortest routes
	ignore := make(map[Edge]struct{})
	for _, edge := range allRoutes.Rank()[:3] {
		ignore[edge.Edge] = struct{}{}
	}

	e := allRoutes.Rank()[0]

	// Remove the three most used edges and walk both groups
	return graph.CountReacheable(e.A, ignore) * graph.CountReacheable(e.B, ignore)
}

func (d day25) Part2(reader io.Reader) int {
	return 0
}
