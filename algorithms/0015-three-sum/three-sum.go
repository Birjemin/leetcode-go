package three_sum

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    length := len(nums)
    var ret [][]int
    if length < 3 {
        return ret
    }
    var left, right, sum int
    // sort nums
    sort.Ints(nums)
    for i := 0; i < length-2; i++ {
        // impossible value, stop search
        if nums[i] > 0 {
            break
        }
        // jump repeat i value
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left = i + 1
        right = length - 1
        // while left < right
        for left < right {
            // jump repeat left value
            if left != i+1 && nums[left] == nums[left-1] {
                left++
                continue
            }
            // question: why not jumping repeat right value?
            // got it!
            sum = nums[i] + nums[right] + nums[left]
            switch {
            case sum < 0:
                left++
            case sum > 0:
                right--
            default:
                ret = append(ret, []int{nums[i], nums[left], nums[right]})
                left++
                right--
            }
        }
    }
    return ret
}
