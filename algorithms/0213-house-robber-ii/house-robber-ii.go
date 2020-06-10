package house_robber_ii

import (
	"sync"
)

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	one, two := cal(nums, 1, length), cal(nums, 0, length-1)
	if one > two {
		return one
	}
	return two
}

func cal(nums []int, start, length int) int {
	currMax, prevMax := 0, 0
	for i := start; i < length; i++ {
		temp := currMax
		if prevMax+nums[i] > currMax {
			currMax = prevMax + nums[i]
		}
		prevMax = temp
	}
	return currMax
}

func rob1(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	var wg sync.WaitGroup
	one, two := 0, 0

	wg.Add(2)

	go func() {
		one = cal(nums, 1, length)
		wg.Done()
	}()

	go func() {
		two = cal(nums, 0, length-1)
		wg.Done()
	}()

	wg.Wait()

	if one > two {
		return one
	}
	return two
}
