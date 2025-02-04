package problem1To99

// A81search 一眼二分，但是没想到怎么处理 0100000的情况，看自己代码 // 如果没办法判断哪半段是有序的，移动左坐标，故整个数列都是同一个数时算法退化为O(n) 悟了
func A81search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	if nums[0] == target {
		return true
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right + left) / 2
		if nums[left] == target || nums[right] == target || nums[mid] == target {
			return true
		}
		if nums[mid] > nums[left] {
			if nums[left] < target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[right] > nums[mid] {
			if nums[mid] < target && nums[right] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}
	return false
}
