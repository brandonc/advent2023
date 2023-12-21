package maths

func SumSlice(slice []int) int {
	r := 0
	for _, n := range slice {
		r += n
	}
	return r
}

func AbsInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func MaxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a int, b ...int) int {
	result := a
	for _, n := range b {
		result = result * n / GCD(result, n)
	}
	return result
}
