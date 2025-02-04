package problem100To199

func A119getRow(rowIndex int) []int {
	res := []int{1}
	for i := 1; i <= rowIndex; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		for j := 0; j < i-1; j++ {
			tmp = append(tmp, res[j]+res[j+1])
		}
		tmp = append(tmp, 1)
		res = tmp
	}
	return res
}
