package maximum_product_subarray

import (
    "math"
)

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

func maxProduct1(nums []int) int {
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
