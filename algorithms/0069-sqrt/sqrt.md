## 问题
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:
```
Input: 4
Output: 2
```

Example 2:
```
Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since 
             the decimal part is truncated, 2 is returned.
```

## 分析
牛顿迭代法

## 最高分
```golang

```

## 实现
牛顿迭代法
```golang
func mySqrt(x int) int {
    r := x
    for r*r > x {
        r = (r + x / r)/2
    }
    return r
}
```

## 改进
二分法
```golang
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
```

## 反思
高中的平方根求法

## 参考
