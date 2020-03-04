package gray_code

import (
    "math"
)

func grayCode(n int) []int {
    if n < 0 {
        return []int{}
    }
    ret := []int{0}
    cal(&ret, uint(n), int(math.Pow(2, float64(n-1))))
    return ret
}

func cal(tmp *[]int, n uint, count int) {
    if n == 1 {
        *tmp = append(*tmp, 1)
    } else if n > 1 {
        // cal the front of n
        cal(tmp, n-1, count/2)
        // add array
        for i := count - 1; i >= 0; i-- {
            *tmp = append(*tmp, (*tmp)[i]+1<<(n-1))
        }
    }
}
