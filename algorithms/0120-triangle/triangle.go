package triangle

import (
    "math"
)

func minimumTotal(triangle [][]int) int {
    var line, val = 0, math.MaxInt64
    for i, t := range triangle {
        for j := range t {
            if i == 0 {
                continue
            }
            var l, m = val, val
            // top left
            if j != 0 {
                l = triangle[i-1][j-1]
            }
            // top middle
            if j < i {
                m = triangle[i-1][j]
            }
            triangle[i][j] += getMin(l, m)
        }
        line += 1
    }

    for _, v := range triangle[line-1] {
        if v < val {
            val = v
        }
    }

    return val
}

func getMin(a, b int) int {
    if a > b {
        return b
    }
    return a
}
