## 问题
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```
Example 2:
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

## 分析
一个有序数组，给定一个target，找到第一次出现的位置和最后一次出现的位置
- 二分法变种题目

## 最高分
```golang
func searchRange(nums []int, target int) []int {
    nlen := len(nums)   
    if nlen == 0 {
        return []int{-1, -1}
    }
    lo, hi := 0, nlen - 1
    var mid int
    
    // search left
    for lo < hi && lo >= 0 && hi < nlen {
        if mid = lo + (hi - lo) / 2; nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    if lo < 0 || hi >= nlen || nums[lo] != target {
        return []int{-1, -1}
    }
    left := hi
    
    //search right
    lo, hi = left, nlen - 1
    for lo < hi && lo >= 0 && hi < nlen {
        if mid = hi - (hi - lo) / 2; nums[mid] <= target {
            lo = mid
        } else {
            hi = mid - 1
        }  
    }
    right := lo
    
    return []int{left, right}
}
```

## 实现
使用二分法，分别查找最大值和最小值（使用了两次二分法）
```
func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    length := len(nums)
    // if nums is invalid
    if length == 0 {
        return res
    }
    // find first position of nums
    searchLeft(&nums, 0, length-1, target, &res[0])
    // could not find the first position of nums
    if res[0] == -1 {
        return res
    }
    // find last position of nums
    searchRight(&nums, res[0], length-1, target, &res[1])
    return res
}

func searchLeft(nums *[]int, start int, end int, target int, res *int) {
    // find it!
    if target == (*nums)[start] {
        *res = start
        return
    }
    div := (end - start) / 2
    // that's mean the index of end is target value or could not found the index
    if div <= 0 {
        if target == (*nums)[end] {
            *res = end
        }
        return
    }
    if target <= (*nums)[div+start] {
        searchLeft(nums, start, end-div, target, res)
    } else {
        searchLeft(nums, start+div, end, target, res)
    }
}

func searchRight(nums *[]int, start int, end int, target int, res *int) {
    if target == (*nums)[end] {
        *res = end
        return
    }
    div := (end - start) / 2
    if div <= 0 {
        if target == (*nums)[start] {
            *res = start
        }
        return
    }
    if target < (*nums)[div+start] {
        searchRight(nums, start, end-div, target, res)
    } else {
        searchRight(nums, start+div, end, target, res)
    }
}
```

## 改进

## 反思

## 参考
