package problem1To99

import "sort"

func A45jump(nums []int) int {
	var arr []int
	for i := 0; i < len(nums); i++ {
		arr = append(arr, 10000)
	}
	arr[0] = 0
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		for j := 1; j <= val && i+j < len(nums); j++ {
			arr[i+j] = min(arr[i]+1, arr[i+j])
		}
	}
	return arr[len(nums)-1]
}

var res [][]int

func A40combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combinationSum(candidates, target, 0, []int{})
	return res
}

func combinationSum(candidates []int, target int, index int, arr []int) {
	if target == 0 {
		a := make([]int, len(arr))
		copy(a, arr)
		res = append(res, a)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		arr = append(arr, candidates[i])
		combinationSum(candidates, target-candidates[i], i+1, arr)
		arr = arr[:len(arr)-1]
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}

}
