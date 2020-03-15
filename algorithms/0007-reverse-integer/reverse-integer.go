package reverse_integer

import "math"

func reverse1(x int) int {
    var res int
    for x != 0 {
        res = res * 10 + x % 10
        x = x / 10
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    return res
}

func reverse2(x int) int {
    var res int
    var tag bool
    if x < 0 {
        tag = true
        x = -x
    }
    for x != 0 {
        res = res * 10 + x % 10
        x = x / 10
        if res > math.MaxInt32 || res < math.MinInt32 {
            return 0
        }
    }
    if tag {
        return -res
    }
    return res
}