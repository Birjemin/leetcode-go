## 问题
A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:
```
Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
```

Example 2:
```
Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5 
Explanation: Your function can return either index number 1 where the peak element is 2, 
             or index number 5 where the peak element is 6.
```

Note:

- Your solution should be in logarithmic complexity.

## 分析
查找数组中的峰值。和153题好像呀~
最简单的方法就是三个数字并排查找

## 最高分
```golang
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] > nums[i+1] {
			return i
		}else if i+1 == len(nums){
			return i
		}
	}
	return 0
}
```

## 实现
```golang
func findPeakElement(nums []int) int {
    var idx int
    for i, v := range nums {
        if nums[idx] > v {
            break
        }
        idx = i
    }
    return idx
}
```

## 改进
使用二分法，来降低时间复杂度，可以知道如果`m[mid] > m[mid]`，说明左边区域有峰值，反之右边有峰值（这个理解满分）
```golang
func findPeakElement(nums []int) int {
    low, high := 0, len(nums)-1

    for low < high {
        mid := (low + high) / 2
        if nums[mid] < nums[mid+1] {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return high
}
```

## 反思

## 参考