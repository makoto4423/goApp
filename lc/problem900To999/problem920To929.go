package problem900To999

func A922sortArrayByParityII(nums []int) []int {
	n := len(nums)
	a, b := make([]int, 0, n/2), make([]int, 0, n/2)
	for _, v := range nums {
		if v&1 == 0 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	res := make([]int, n)
	for i := 0; i < n/2; i++ {
		res[2*i], res[2*i+1] = a[i], b[i]
	}
	return res
}
