package string_to_integer_atoi

import (
    "math"
    "strings"
)

func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    if str == "" {
        return 0
    }
    var num int
    var flag bool
    for i, s := range str {
        // check calculation sign
        if i == 0 {
            if s == '-' {
                flag = true
                continue
            } else if s == '+' {
                continue
            }
        }
        // valid numbers
        if s >= '0' && s <= '9' {
            num = 10*num + int(s) - '0'
            // when num > maxInt32, break for expression
            if num > math.MaxInt32 {
                break
            }
        } else {
            break
        }
    }
    // check num
    if num >= 0 && num <= math.MaxInt32 {
        if flag {
            return -num
        } else {
            return num
        }
    } else {
        // out of range
        if flag {
            return math.MinInt32
        } else {
            return math.MaxInt32
        }
    }
}

func myAtoi1(str string) int {
    str = strings.TrimSpace(str)
    if str == "" {
        return 0
    }

    res, flag, len, idx := 0, false, len(str), 0

    // check sign
    if str[idx] == '+' {
        idx++
    } else if str[idx] == '-' {
        flag = true
        idx++
    }
    for ; idx < len && str[idx] >= '0' && str[idx] <= '9'; idx++ {
        res = 10*res + int(str[idx]) - '0'
        // when num > maxInt32, out of range
        if res > math.MaxInt32 {
            if flag {
                return math.MinInt32
            } else {
                return math.MaxInt32
            }
        }
    }
    if flag {
        res = -res
    }
    return res
}
