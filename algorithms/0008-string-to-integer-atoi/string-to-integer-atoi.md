## 问题
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

* Only the space character ' ' is considered as whitespace character.
* Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:
```
Input: "42"
Output: 42
```

Example 2:
```
Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
```

Example 3:
```
Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
```

Example 4:
```
Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical 
             digit or a +/- sign. Therefore no valid conversion could be performed.
```

Example 5:
```
Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
```
## 分析
实现string转int函数atoi，难度一般，但是有很多边界需要考虑

## 最高分
代码更加简洁，可以借鉴
```golang
func myAtoi(str string) int {
    res, sign, len, idx := 0, 1, len(str), 0

    // Skip leading spaces
    for idx < len && (str[idx] == ' ' || str[idx] == '\t') {
        idx++
    }

    // +/- Sign
    if idx < len {
        if str[idx] == '+' {
            sign = 1
            idx++
        } else if str[idx] == '-' {
            sign = -1
            idx++
        }
    }

    // Numbers
    for idx < len && str[idx] >= '0' && str[idx] <= '9'{
        res = res * 10 + int(str[idx]) - '0'
        if sign * res > math.MaxInt32 {
            return math.MaxInt32
        } else if sign * res < math.MinInt32 {
            return math.MinInt32
        }
        idx++
    }

    return res * sign
}
```


## 实现
```golang
func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    if str == "" {
        return 0
    }
    var num int
    var flag bool
    for i, s := range str {
        // check calculation signs
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
```

## 改进
```golang
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
```

## 反思

## 参考
