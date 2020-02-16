package combination_sum

import (
    "sort"
)

func combinationSum(candidates []int, target int) [][]int {
    var ret [][]int
    sort.Ints(candidates)
    dsf(&ret, candidates, []int{}, target)
    return ret
}

func dsf(ret *[][]int, candidates []int, solution []int, target int) {
    if target == 0 {
        // *ret = append(*ret, solution)
        b := make([]int, len(solution))
        copy(b, solution)
        *ret = append(*ret, b)
        return
    }
    for i, v := range candidates {
        // adding this expression could reduce the memory storage
        if v > target {
            return
        }
        dsf(ret, candidates[i:], append(solution, v), target-v)
    }
}

func combinationSum1(candidates []int, target int) [][]int {
    var ret [][]int
    sort.Ints(candidates)
    dsf1(&ret, candidates, []int{}, target)
    return ret
}

func dsf1(ret *[][]int, candidates []int, solution []int, target int) {
    for i, v := range candidates {
        if v > target {
            return
        }
        // add v
        solution = append(solution, v)
        if target == v {
            // *ret = append(*ret, solution)
            b := make([]int, len(solution))
            copy(b, solution)
            *ret = append(*ret, b)
            return
        }
        dsf1(ret, candidates[i:], solution, target-v)
        // should subtract the last element of solution
        solution =solution[:len(solution)-1]
    }
}