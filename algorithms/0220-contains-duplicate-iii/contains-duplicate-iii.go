package contains_duplicate_iii

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length && j <= i+k; j++ {
			tmp := nums[j] - nums[i]
			if tmp > 0 {
				if tmp <= t {
					return true
				}
			} else {
				if -tmp <= t {
					return true
				}
			}
		}
	}
	return false
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	ret := make(map[int]int, k+1)
	for i, v := range nums {
		for j := range ret {
			j -= nums[i]
			if j >= 0 {
				if j <= t {
					return true
				}
			} else {
				if -j <= t {
					return true
				}
			}
		}
		ret[v] = i
		if len(ret) > k {
			delete(ret, nums[i-k])
		}
	}
	return false
}
