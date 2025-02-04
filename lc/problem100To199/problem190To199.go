package problem100To199

func A198rob(nums []int) int {
	n := len(nums)
	// make n 长度len(arr)，2 容量cap(arr)，容量大于长度
	//arr := make([][]int, n, 2)
	arr := make([][2]int, n)
	arr[0][1] = nums[0]
	for i := 1; i < n; i++ {
		arr[i][0] = arr[i-1][1]
		arr[i][1] = max(arr[i-1][1], arr[i-1][0]+nums[i])
	}
	return max(arr[n-1][0], arr[n-1][1])
}
