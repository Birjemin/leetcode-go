## 问题
Given an integer n, return the number of trailing zeroes in n!.

Example 1:
```
Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.
```

Example 2:
```
Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.
Note: Your solution should be in logarithmic time complexity.
```

## 分析
给定一个整数 n，返回 n! 结果尾数中零的数量。
- dp算法复杂度为O(n)，不符合要求
- 本质上就是计算有多少个因子5

## 最高分
```golang
func trailingZeroes(n int) int {
    res := 0
    for n >= 5 {
        n /= 5
        res += n
    }
    return res
}
```

## 实现
因子求和
```golang
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
```

## 改进
```golang

```

## 反思

## 参考