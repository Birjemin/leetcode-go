## 问题
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:
```
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
```

Example 2:
```
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
```

## 分析
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。
找规律的题目

## 最高分
```golang
func maxProduct(nums []int) int {
    tmp := nums[0]
    max := nums[0]
    min := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            max, min = min, max
        }
        max *= nums[i]
        if max < nums[i] {
            max = nums[i]
        }
        min *= nums[i]
        if min > nums[i] {
            min = nums[i]
        }
        if max > tmp {
            tmp = max
        }
    }
    return tmp
}
```

## 实现
性能一般，可以返现两次循环，是否可以是否dp算法
```golang
func maxProduct(nums []int) int {
    sum, max, length := 1, nums[0], len(nums)

    for i := range nums {
        sum = 1
        for j := i; j < length; j++ {
            if nums[j] > max {
                max = nums[j]
            }
            sum *= nums[j]
            if sum > max {
                max = sum
            }
            if sum == 0 {
                sum = 1
            }
        }
    }
    return max
}
```

## 改进
dp算法（看了答案，有O(n)复杂度的解法，所以还得精炼）
```golang
func maxProduct(nums []int) int {
    max, length := math.MinInt64, len(nums)
    dp := make([]int, length)
    for i := 0; i < length; i++ {
        for j := 0; j <= i; j++ {
            if i == j {
                dp[j] = nums[length-1-i]
            } else {
                dp[j] *= nums[length-1-i]
            }
            if max < dp[j] {
                max = dp[j]
            }
        }
    }
    return max
}
```

## 改进
每次循环只需要计算最大值和最小值即可~~，然后和当前val对比，三者取出最大值和最小值
```golang
func maxProduct2(nums []int) int {
    var ret, min, max int
    for i, v := range nums {
        if i == 0 {
            ret, min, max = v, v, v
            continue
        }

        // get min and max
        min, max = v*min, v*max
        if min > max {
            min, max = max, min
        }

        // compare min max v values
        if v > max {
            max = v
        } else if v < min {
            min = v
        }

        if ret < max {
            ret = max
        }
    }
    return ret
}
```

## 反思

## 参考