package day17

import (
	"bufio"
	"container/heap"
	"io"
	"strconv"

	"github.com/brandonc/advent2023/solutions/solution"
)

type day17 struct{}

func Factory() solution.Solver {
	return day17{}
}

type Map [][]int

type dir rune

const (
	none  dir = ' '
	north dir = 'N'
	south dir = 'S'
	east  dir = 'E'
	west  dir = 'W'
)

var deltas = map[dir][]int{
	north: {-1, 0},
	south: {1, 0},
	east:  {0, 1},
	west:  {0, -1},
}

var oppositeDirs = map[dir]dir{
	north: south,
	south: north,
	east:  west,
	west:  east,
}

type Coords struct {
	y, x int
	dir  string
}

// An Item is something we manage in a priority queue.
type Crucible struct {
	coords Coords
	cost   int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// Implements a min heap
type PriorityQueue []*Crucible

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Would be better to implement an a* heuristic score
	return pq[i].cost <= pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Crucible)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (m Map) FindBestPath(minTravel, maxTravel int) int {
	open := PriorityQueue{}
	closed := make(map[Coords]struct{})
	heap.Push(&open, &Crucible{coords: Coords{0, 0, "S"}, cost: 0})
	heap.Push(&open, &Crucible{coords: Coords{0, 0, "E"}, cost: 0})

	for open.Len() > 0 {
		current := heap.Pop(&open).(*Crucible)
		if _, ok := closed[current.coords]; ok {
			continue
		}
		closed[current.coords] = struct{}{}

		if current.coords.x == len(m[0])-1 && current.coords.y == len(m)-1 && len(current.coords.dir) >= minTravel {
			return current.cost
		}

		for d, change := range deltas {
			lastDirection := dir(current.coords.dir[len(current.coords.dir)-1])
			if d == oppositeDirs[dir(lastDirection)] {
				// Cannot turn around
				continue
			}

			if d != lastDirection && len(current.coords.dir) < minTravel {
				// Cannot change direction until we've traveled at least minTravel
				continue
			}

			// Cannot move more than maxTravel in a single direction
			if d == lastDirection && len(current.coords.dir) == maxTravel {
				continue
			}

			newDir := string(d)
			if newDir == string(lastDirection) {
				newDir = current.coords.dir + newDir
			}

			child := Coords{
				y:   current.coords.y + change[0],
				x:   current.coords.x + change[1],
				dir: newDir,
			}

			if child.y < 0 || child.y >= len(m) || child.x < 0 || child.x >= len(m[0]) {
				continue
			}

			if _, ok := closed[child]; ok {
				continue
			}

			heap.Push(&open, &Crucible{
				coords: child,
				cost:   current.cost + m[child.y][child.x],
			})
		}
	}
	panic("Did not find exit")
}

func parseMap(reader io.Reader) Map {
	result := make(Map, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			num, _ := strconv.Atoi(string(line[i]))
			row[i] = num
		}
		result = append(result, row)
	}
	return result
}

func (d day17) Part1(reader io.Reader) int {
	m := parseMap(reader)

	answer := m.FindBestPath(1, 3)
	return answer
}

func (d day17) Part2(reader io.Reader) int {
	m := parseMap(reader)

	answer := m.FindBestPath(4, 10)
	return answer
}
