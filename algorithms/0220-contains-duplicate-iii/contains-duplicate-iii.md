## 问题
Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

Example 1:
```
Input: nums = [1,2,3,1], k = 3, t = 0
Output: true
```

Example 2:
```
Input: nums = [1,0,1,1], k = 1, t = 2
Output: true
```

Example 3:
```
Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
```

## 分析
给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。

217题找出重复元素；219找出重复元素并且索引的差的绝对值最大为k；本题是的条件是值的差的绝对值最大为t，索引的差的绝对值最大为k

## 最高分
```golang
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n && count < k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
			count++
		}
	}
	return false
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```

## 实现
最基本的实现
```golang
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
```

## 改进
```golang
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
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
```

## 反思

## 参考