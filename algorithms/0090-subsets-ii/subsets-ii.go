package subsets_ii

import (
    "sort"
)

func subsetsWithDup(nums []int) [][]int {
    var ret [][]int
    sort.Ints(nums)
    dfs(&ret, []int{}, nums, 0, len(nums))
    return ret
}

func dfs(ret *[][]int, tmp, nums []int, start, end int) {
    b := make([]int, len(tmp))
    copy(b, tmp)
    *ret = append(*ret, b)
    if start == end {
        return
    }
    for i := start; i < end; i++ {
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        dfs(ret, append(tmp, nums[i]), nums, i+1, end)
    }
}
