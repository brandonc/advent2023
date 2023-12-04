package ds

type IntSet map[int]struct{}

func NewIntSet(numbers []int) IntSet {
	result := make(IntSet)

	for _, n := range numbers {
		result[n] = struct{}{}
	}

	return result
}

func (i IntSet) Intersect(other IntSet) []int {
	shared := make([]int, 0)
	for n := range i {
		if _, has := other[n]; has {
			shared = append(shared, n)
		}
	}
	return shared
}
