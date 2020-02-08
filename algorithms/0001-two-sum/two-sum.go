package two_sum

/**
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

// 直接两个循环
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if target - nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}


// 降维度：先登记自己的匹配对象，然后遍历查找
func twoSum2(nums []int, target int) []int {
	lookup := make(map[int]int, len(nums))
	for i, v := range nums {
		lookup[v] = i
	}
	for i, v := range nums {
		if j, ok := lookup[target-v]; ok {
			if i != j {
				return []int{i, j}
			}

		}
	}

	return nil
}

// 合并循环 来着参考2：先查找自己的匹配对象，没有找到则登记自己
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
	return nil
}