package merge_intervals

import (
    "sort"
)

func merge(intervals [][]int) [][]int {
    length := len(intervals)
    if length == 0 {
        return [][]int{}
    } else if length == 1 {
        return intervals
    }

    // sort the arrays
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // init the ret
    ret, j := [][]int{{intervals[0][0], intervals[0][1]}}, 0
    for i := 1; i < length; i++ {
        if ret[j][1] >= intervals[i][0] {
            ret[j][1] = max(ret[j][1], intervals[i][1])
        } else {
            ret = append(ret, intervals[i])
            j++
        }
    }
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
