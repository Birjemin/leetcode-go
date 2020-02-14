## 问题
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

Example 2:
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

## 分析
将一个升序数列在某一个分割点翻转，然后查找目标值
先找到分割点index，然后判断target在前部分查找还是后部分查找，反之返回-1
- 二分法
- 可以使用递归实现

## 最高分
// 直接一个循环中移动，也是使用了二分法，很巧妙的方式处理循环
```golang
func search(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n - 1

    for l <= r {
        if nums[l] == target {
            return l
        }
        if nums[r] == target {
            return r
        }
        mid := (l + r) / 2
        if nums[mid] == target {
            return mid
        }
        // move l and r
        if nums[l] > nums[r] {
            l += 1
            r -= 1
        } else {
            if nums[mid] > target {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}
```


## 实现
先找到翻转的分割点，然后使用二分法递归查找
```golang
func search(nums []int, target int) int {
    var idx int
    length := len(nums)
    // if array is null
    if length == 0 {
        return -1
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

func dichotomy(nums *[]int, start int, end int, target int) int {
    // unbelievable!!! end < start
    if end < start {
        return -1
    }
    div := (end - start) / 2
    if div == 0 {
        if target == (*nums)[end] {
            return end
        } else if target == (*nums)[start] {
            return start
        } else {
            return -1
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
