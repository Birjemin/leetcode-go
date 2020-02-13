package divide_two_integers

import (
    "math"
)

func divide(dividend int, divisor int) int {
    var pos bool
    var count int
    dividend, divisor, pos = abs(dividend, divisor)
    if divisor == 1 {
        return special(dividend, pos)
    }
    count = div(dividend, divisor)
    if pos {
        return count
    }
    return -count
}

func special(dividend int, pos bool) int {
    if pos {
        if dividend > math.MaxInt32 {
            return math.MaxInt32
        }
        return dividend
    } else {
        if -dividend < math.MinInt32 {
            return math.MinInt32
        }
        return -dividend
    }
}

func div(dividend int, divisor int) int {
    if dividend < divisor {
        return 0
    }
    num, count := divisor, 1
    for dividend > num<<1 {
        num, count = num<<1, count<<1
    }
    count += div(dividend-num, divisor)
    return count
}

func abs(num1, num2 int) (int, int, bool) {
    pos := true
    if num1 < 0 {
        num1 = -num1
        pos = false
    }

    if num2 < 0 {
        num2 = -num2
        if pos {
            pos = false
        } else {
            pos = true
        }
    }
    return num1, num2, pos
}
