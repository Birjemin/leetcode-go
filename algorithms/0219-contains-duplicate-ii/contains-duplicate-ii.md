## 问题
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:
```
Input: nums = [1,2,3,1], k = 3
Output: true
```

Example 2:
```
Input: nums = [1,0,1,1], k = 1
Output: true
```

Example 3:
```
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
```

## 分析
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 `nums [i] = nums [j]`，并且 i 和 j 的差的 绝对值 至多为 k。

和217基础上改进即可

## 最高分
```golang
func containsNearbyDuplicate(nums []int, k int) bool {
    // We only need a map of capacity k
    index := make(map[int]int, k)
    for i, n := range nums {
        // For each number, we only need to consider the previous k numbers
        // So it's safe to remove i - k - 1
        if i > k {
            obsoleteIdx := i - k - 1
            if index[nums[obsoleteIdx]] == obsoleteIdx {
                delete(index, nums[obsoleteIdx])
            }
        }
        if prev, ok := index[n]; ok {
            if i - prev <= k {
                return true
            } 
        }
        index[n] = i
    }
    return false
}
```

## 实现
map解决
```golang
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
```

## 改进
滑动窗口+map解决，实际上就是上面的方法改进，降低空间使用
```golang
func containsNearbyDuplicate(nums []int, k int) bool {
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
```

## 反思

## 参考