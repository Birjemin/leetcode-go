## 问题
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:
```
Input: 2.00000, 10
Output: 1024.00000
```

Example 2:
```
Input: 2.10000, 3
Output: 9.26100
```

Example 3:
```
Input: 2.00000, -2
Output: 0.25000
Explanation: 2^-2 = 1/(2^2) = 1/4 = 0.25
```

Note:

- -100.0 < x < 100.0
- n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]

## 分析
实现幂次方方程，和29题类似的情况

## 最高分
```golang
func myPow(x float64, n int) float64 {
    if n == 0 {return 1}
    if n < 0 {
        //take care Integer.MIN_VALUE
        return 1 / (x * (myPow(x, -(n + 1))))
    }
    res := 1.0
    for n > 1 {
        if n % 2 == 1 {
            res *= x
        }
        x = x * x
        n /= 2
    }
    res *= x
    return res
}
```

## 实现
和29题类似 2^10 = 2^8+2^4
```golang
func myPow(x float64, n int) float64 {
    if n < 0 {
        x, n = 1/x, -n
    }
    return pow(x, n)
}

func pow(x float64, n int) float64 {
    if n <= 0 {
        return 1
    }
    temp, i := x, 1
    for 2*i <= n {
        x, i = x*x, 2*i
    }
    return x * pow(temp, n-i)
}
```

## 改进
```golang

```

## 反思

## 参考