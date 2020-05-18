package contains_duplicate

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	ret := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := ret[v]; ok {
			return true
		}
		ret[v] = true
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
