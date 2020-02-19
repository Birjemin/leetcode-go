package permutations_ii

import (
    "sort"
)

func permuteUnique(nums []int) [][]int {
    length := len(nums)
    if length == 1 {
        return [][]int{
            nums,
        }
    }
    var ret [][]int
    sort.Ints(nums)
    dfs(&ret, nums, []int{})
    return ret
}

func dfs(ret *[][]int, nums, temp []int) {
    if len(nums) == 0 {
        b := make([]int, len(temp))
        copy(b, temp)
        *ret = append(*ret, b)
    }
    for i, v := range nums {

        if i == 0 {
            dfs(ret, nums[i+1:], append(temp, v))
        } else {
            // if the value is equal to the front value, just continue the circle
            if i != 0 && nums[i-1] == v {
                continue
            }
            b := make([]int, len(nums))
            copy(b, nums)
            dfs(ret, append(b[:i], b[i+1:]...), append(temp, v))
        }
    }
}
