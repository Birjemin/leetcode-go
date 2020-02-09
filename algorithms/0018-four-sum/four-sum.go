package four_sum

import (
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    length := len(nums)
    var ret [][]int
    if length < 3 {
        return ret
    }
    var left, right, sum int
    // sort nums
    sort.Ints(nums)
    for j := 0; j < length-3; j++ {
        // jump repeat j value
        if j > 0 && nums[j] == nums[j-1] {
            continue
        }
        for i := j + 1; i < length-2; i++ {
            // impossible value
            if target > 0 && nums[j]+nums[i] > target {
                break
            }
            // jump repeat i value
            if i > j+1 && nums[i] == nums[i-1] {
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
                // got it!
                sum = nums[j] + nums[i] + nums[right] + nums[left]
                switch {
                case sum < target:
                    left++
                case sum > target:
                    right--
                default:
                    ret = append(ret, []int{nums[j], nums[i], nums[left], nums[right]})
                    left++
                    right--
                }
            }
        }
    }
    return ret
}

func fourSum1(nums []int, target int) [][]int {
    length := len(nums)
    var ret [][]int
    if length < 3 {
        return ret
    }
    var left, right, sum int
    // sort nums
    sort.Ints(nums)
    for j := 0; j < length-3; j++ {
        // jump repeat j value
        if j > 0 && nums[j] == nums[j-1] {
            continue
        }
        // boundary check
        if nums[j]+nums[length-3]+nums[length-2]+nums[length-1] < target {
            continue
        }
        for i := j + 1; i < length-2; i++ {
            // impossible value
            if target > 0 && nums[j]+nums[i] > target {
                break
            }
            // jump repeat i value
            if i > j+1 && nums[i] == nums[i-1] {
                continue
            }
            // boundary check
            if nums[j]+nums[i]+nums[length-2]+nums[length-1] < target {
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
                // got it!
                sum = nums[j] + nums[i] + nums[right] + nums[left]
                switch {
                case sum < target:
                    left++
                case sum > target:
                    right--
                default:
                    ret = append(ret, []int{nums[j], nums[i], nums[left], nums[right]})
                    left++
                    right--
                }
            }
        }
    }
    return ret
}
