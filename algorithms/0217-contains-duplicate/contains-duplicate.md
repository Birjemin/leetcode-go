## 问题
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:
```
Input: [1,2,3,1]
Output: true
```

Example 2:
```
Input: [1,2,3,4]
Output: false
```

Example 3:
```
Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
```

## 分析
给定一个整数数组，判断是否存在重复元素。

如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

最简单的方式就是map、或者先排序再查找即可

## 最高分
```golang
func containsDuplicate(nums []int) bool {
	appearance := make(map[int]bool)

	for _, num := range nums {
		if _, ok := appearance[num]; ok {
			return true
		}
		appearance[num] = true
	}
	return false
}
```

## 实现
map映射
```golang
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
```

## 改进
快速排序的改进
```golang
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
```

## 反思

## 参考