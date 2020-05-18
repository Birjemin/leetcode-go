package contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	ret := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := ret[v]; ok {
			if i-j <= k {
				return true
			}
		}
		ret[v] = i
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	ret := make(map[int]int, k+1)
	for i, v := range nums {
		if _, ok := ret[v]; ok {
			return true
		}
		ret[v] = i
		if len(ret) > k {
			delete(ret, nums[i-k])
		}
	}
	return false
}
