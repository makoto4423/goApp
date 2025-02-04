package problem500To599

import "container/list"

func A576findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	mod := 1000000007
	xGrid, yGrid := []int{-1, 1, 0, 0}, []int{0, 0, 1, -1}
	res := 0
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	arr[startRow][startColumn] = 1
	l := list.New()
	l.PushBack([2]int{startRow, startColumn})
	k := 0
	for k < maxMove {
		tmp, ele := list.New(), l.Front()
		for ele != nil {
			p := ele.Value.([2]int)
			for i := 0; i < 4; i++ {
				x := p[0] + xGrid[i]
				y := p[1] + yGrid[i]
				if x < 0 || x >= m || y < 0 || y >= n {
					res = (res + arr[p[0]][p[1]]) % mod
				} else {
					arr[x][y] = (arr[x][y] + arr[p[0]][p[1]]) % mod
					tmp.PushBack([2]int{x, y})
				}
			}
			arr[p[0]][p[1]] = 0
			ele = ele.Next()
		}
		k++
		l = tmp
	}
	return res
	//mod := 1000000007
	//xGrid, yGrid := []int{-1, 1, 0, 0}, []int{0, 0, 1, -1}
	//res := 0
	//arr := make([][]int, m)
	//for i := 0; i < m; i++ {
	//	arr[i] = make([]int, n)
	//}
	//arr[startRow][startColumn] = 1
	//l := [][]int{{startRow, startColumn}}
	//k := 0
	//for k < maxMove {
	//	// 是go的问题，还是lc的问题，逻辑是对的，但是预测内存容量这块，不同的测试用例， 容量要不同的值，否则就报 cannot allocate memory
	//	// 不设置又超时, lc 官方答案不是通过中间数组，而是原地替换
	//	tmp := make([][]int, 0, len(l)*4)
	//	for _, p := range l {
	//		for i := 0; i < 4; i++ {
	//			x := p[0] + xGrid[i]
	//			y := p[1] + yGrid[i]
	//			if x < 0 || x >= m || y < 0 || y >= n {
	//				res = (res + arr[p[0]][p[1]]) % mod
	//			} else {
	//				arr[x][y] = (arr[x][y] + arr[p[0]][p[1]]) % mod
	//				tmp = append(tmp, []int{x, y})
	//			}
	//		}
	//		arr[p[0]][p[1]] = 0
	//	}
	//	k++
	//	l = tmp
	//}
	//return res
}
