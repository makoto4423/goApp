package problem200To299

func A219containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i <= k && i < len(nums); i++ {
		if m[nums[i]] != 0 {
			return true
		}
		m[nums[i]] = 1
	}
	for i := k + 1; i < len(nums); i++ {
		m[nums[i-k-1]] = 0
		if m[nums[i]] != 0 {
			return true
		}
		m[nums[i]] = 1
	}
	return false
	//     set := map[int]struct{}{}
	//    for i, num := range nums {
	//        if i > k {
	//            delete(set, nums[i-k-1])
	//        }
	//        if _, ok := set[num]; ok {
	//            return true
	//        }
	//        set[num] = struct{}{}
	//    }
	//    return false
}
