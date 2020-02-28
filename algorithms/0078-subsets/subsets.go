package subsets

func subsets(nums []int) [][]int {
    var ret [][]int
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
        dfs(ret, append(tmp, nums[i]), nums, i+1, end)
    }
}
