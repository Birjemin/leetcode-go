package factorial_trailing_zeroes

func trailingZeroes(n int) int {
    var k, ret int
    t := 1
    for  {
        t *= 5
        k = n / t
        if k < 1 {
            break
        }
        ret += k
    }
    return ret
}

func trailingZeroes1(n int) int {
    var ret int
    for  n > 1 {
        n /= 5
        ret += n
    }
    return ret
}

func trailingZeroes2(n int) int {
    if n < 1 {
        return 0
    }
    l := n/5
    return l + trailingZeroes2(l)
}
