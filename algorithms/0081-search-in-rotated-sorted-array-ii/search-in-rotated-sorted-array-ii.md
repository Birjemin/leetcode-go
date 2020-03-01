## 问题
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:
```
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
```

Example 2:
```
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
```

Follow up:
```
This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?
```

## 分析
33题的延伸
先找到翻转的分割点，然后使用二分法递归查找
和33题一样的解答

## 最高分
```golang
func search(nums []int, target int) bool {
    length := len(nums)
    if length == 0 {
        return false
    }
    k := 1
    for k < len(nums) && nums[k-1] <= nums[k] {
        k++
    }

    i, j := 0, length-1
    for i <= j {
        m := (i + j) / 2
        med := (m + k) % length
        switch {
        case nums[med] < target:
            i = m + 1
        case target < nums[med]:
            j = m - 1
        default:
            return true
        }
    }

    return false
}
```


## 实现
```golang
func search(nums []int, target int) bool {
    var idx int
    length := len(nums)
    // if array is null
    if length == 0 {
        return false
    }
    // search the reverse position
    for ; idx < length-1; idx++ {
        if nums[idx] > nums[idx+1] {
            break
        }
    }
    // handle
    if target >= nums[0] {
        // search the front part of array
        return dichotomy(&nums, 0, idx, target)
    } else {
        // search the next part of array
        return dichotomy(&nums, idx+1, length-1, target)
    }
}

func dichotomy(nums *[]int, start int, end int, target int) bool {
    // unbelievable!!! end < start
    if end < start {
        return false
    }
    div := (end - start) / 2
    if div == 0 {
        if target == (*nums)[end] || target == (*nums)[start] {
            return true
        } else {
            return false
        }
    }
    // search the front part of array
    if (*nums)[start+div] > target {
        return dichotomy(nums, start, end-div, target)
    } else {
        // search the next part of array
        return dichotomy(nums, start+div, end, target)
    }
}
```

## 改进
```golang

```

## 反思

## 参考