package problem500To599

func A598maxCount(m int, n int, ops [][]int) int {
	a, b := m, n
	for _, op := range ops {
		a = min(a, op[0])
		b = min(b, op[1])
	}
	return a * b
}
