## 问题
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:
```
Input: dividend = 10, divisor = 3
Output: 3
```

Example 2:
```
Input: dividend = 7, divisor = -3
Output: -2
```

Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

## 分析
自行编写除法方程，先转成正数，然后使用进位来处理

## 最高分
```golang

```

## 实现
左移1即乘以2
```golang
func divide(dividend int, divisor int) int {
    if divisor == 0 {
        return math.MaxInt32
    }
    var pos bool
    var count, num int
    dividend, divisor, pos = abs(dividend, divisor)
    if divisor == 1 {
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

    num = divisor << 1
    for {
        dividend -= num
        count += 2
        if dividend < 0 {
            if -dividend > divisor {
                count = 0
            } else {
                count -= 1
            }
            break
        } else if dividend < divisor {
            break
        } else if dividend < num {
            count += 1
            break
        }
    }
    if pos {
        return count
    }
    return -count
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
```

## 改进
使用根据移位递归来做，比如83/4 = 4 * 16 + 4 * 4 + 3
```golang
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

```

## 反思

## 参考
