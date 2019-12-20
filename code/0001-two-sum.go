package main

import "fmt"

/**
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

// 直接两个循环
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if target - nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// 降维度
func twoSum2(nums []int, target int) []int {
	lookup := make(map[int]int, len(nums))
	// 键名为nums的值，键值为索引
	for i, v := range nums {
		lookup[v] = i
	}

	// 循环i
	for i, v := range nums {
		// j为值 ok为布尔值，是否存在
		if j, ok := lookup[target-v]; ok {
			return []int{i, j}
		}
	}

	return []int{}
}

// 合并循环 来着参考2
func twoSum3(nums []int, target int) []int {
	lookup := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := lookup[target-v]; ok {
			return []int{j, i}
		}
		lookup[v] = i
	}

	return nil
}

// 完美解答（因为默认为0） 来着参考3
func twoSum4(nums []int, target int) []int {
	lookup := make(map[int]int, len(nums))
	for i, v := range nums {
		j, ok := lookup[-v]
		if ok {
			return []int{j, i}
		}
		lookup[v - target] = i
	}
	return []int{}
}

func main() {
	var nums = []int{2, 7, 11, 15}
	res1 := twoSum1(nums, 22)
	fmt.Println(res1)
	res2 := twoSum2(nums, 22)
	fmt.Println(res2)
	res3 := twoSum3(nums, 22)
	fmt.Println(res3)
	res4 := twoSum4(nums, 22)
	fmt.Println(res4)
}
