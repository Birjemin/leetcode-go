package combination_sum_ii

import (
    "sort"
)

func combinationSum2(candidates []int, target int) [][]int {
    var ret [][]int
    sort.Ints(candidates)
    dfs(&ret, candidates, []int{}, target)
    return ret
}

func dfs(ret *[][]int, candidates []int, solution []int, target int) bool {
    if target == 0 {
        // *ret = append(*ret, solution)
        b := make([]int, len(solution))
        copy(b, solution)
        *ret = append(*ret, b)
        return true
    }
    for i, v := range candidates {
        // adding this expression could reduce the memory storage
        if v > target {
            return false
        }
        // next value equal to the current value, just jump up this value
        if i != 0 && candidates[i-1] == v {
            continue
        }
        // if find the value, then break this circle
        if dfs(ret, candidates[i+1:], append(solution, v), target-v) {
            break
        }
    }
    return false
}
