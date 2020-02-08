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
输入一个数列nums，一个target，然后在nums中找到两个num，它们的和是target，返回这两个数的索引，三种方式：
- 直接两个循环查找
- 记录target - num[i]，逐一判断是否匹配
- 判断是否匹配，不匹配则记录要匹配的对象target-nums[i]

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
```
* 时间复杂度：O(n^2)
* 空间复杂度：O(1)

## 改进
题干提到每个元素唯一那么可以用map来映射，将值作为map键，将索引作为map的值
```golang
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int, len(nums))
	// i为索引，v为值
	for i, v := range nums {
		if j, ok := lookup[target-v]; ok {
			return []int{j, i}
		}
		// 填充v
		lookup[v] = i
	}

	return nil
}
```
* 时间复杂度：O(n)
* 空间复杂度：O(n)

## 反思
利用map来降低维度，空间换区时间~~

## 参考
1. [https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0001.two-sum/two-sum.go](https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0001.two-sum/two-sum.go)
2. [https://github.com/wind-liang/leetcode/blob/master/leetCode-1-Two-Sum.md](https://github.com/wind-liang/leetcode/blob/master/leetCode-1-Two-Sum.md)
3. [https://leetcode.com/problems/two-sum/discuss/261590/100-Golang](https://leetcode.com/problems/two-sum/discuss/261590/100-Golang)
