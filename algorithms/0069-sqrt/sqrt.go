package sqrt

func mySqrt1(x int) int {
    r := x
    for r*r > x {
        r = (r + x/r) / 2
    }
    return r
}

func mySqrt(x int) int {
    if x == 1 {
        return 1
    }
    var low, mid, sqr int
    high := x
    for {
        mid = (low + high) / 2
        if mid == low {
            return low
        }
        sqr = mid * mid
        if sqr > x {
            high = mid
        } else if sqr != x {
            low = mid
        } else {
            return mid
        }
    }
}
