## 问题

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 分析
两层循环最简单

## 最高分
```golang
func twoSum(nums []int, target int) []int {
    lookup := make(map[int]int)
    for i, v := range nums {
        j, ok := lookup[-v]
        lookup[v - target] = i
        if ok {
            return []int{j, i}
        }
    }
    return []int{}
}
```


## 实现
```golang
func twoSum(nums []int, target int) []int {
	var res []int
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if target - nums[i] == nums[j] {
				res = append(res, i, j)
			}
		}
	}
	return res
}
```

## 改进
题干提到每个元素唯一那么可以用map来映射
```golang
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := lookup[target-v]; ok {
			return []int{j, i}
		}
		lookup[v] = i
	}

	return nil
}
```

## 反思
利用map来降低维度，空间换区时间~~

## 参考
1. [https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0001.two-sum/two-sum.go](https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0001.two-sum/two-sum.go)
2. [https://github.com/wind-liang/leetcode/blob/master/leetCode-1-Two-Sum.md](https://github.com/wind-liang/leetcode/blob/master/leetCode-1-Two-Sum.md)
3. [https://leetcode.com/problems/two-sum/discuss/261590/100-Golang](https://leetcode.com/problems/two-sum/discuss/261590/100-Golang)
