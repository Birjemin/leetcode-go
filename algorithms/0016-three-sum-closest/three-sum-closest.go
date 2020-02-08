package three_sum_closest

import (
    "math"
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    length := len(nums)
    var left, right, sub int
    ret := math.MaxInt32
    // sort nums
    sort.Ints(nums)
    for i := 0; i < length-2; i++ {
        left = i + 1
        right = length - 1
        // while left < right
        for left < right {
            // jump repeat left value
            if left != i+1 && nums[left] == nums[left-1] {
                left++
                continue
            }
            sub = nums[i] + nums[right] + nums[left] - target
            // got it!
            if abs(sub) < abs(ret-target) {
                ret = sub + target
            }
            switch {
            case sub < 0:
                left++
            case sub > 0:
                right--
            default:
                return ret
            }
        }
    }
    return ret
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
