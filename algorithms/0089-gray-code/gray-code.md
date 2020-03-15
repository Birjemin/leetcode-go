## 问题
The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

Example 1:
```
Input: 2
Output: [0,1,3,2]
Explanation:
00 - 0
01 - 1
11 - 3
10 - 2

For a given n, a gray code sequence may not be uniquely defined.
For example, [0,2,3,1] is also a valid gray code sequence.

00 - 0
10 - 2
11 - 3
01 - 1
```

Example 2:
```
Input: 0
Output: [0]
Explanation: We define the gray code sequence to begin with 0.
             A gray code sequence of n has size = 2n, which for n = 0 the size is 20 = 1.
             Therefore, for n = 0 the gray code sequence is [0].
```

## 分析
格雷码，镜像排列 [wiki](https://zh.wikipedia.org/zh-hans/%E6%A0%BC%E9%9B%B7%E7%A0%81)

## 最高分
类似的原理
```golang
func grayCode(n int) []int {
    if n <= 0 {
        return []int{0}
    }

    if n == 1 {
        return []int{0, 1}
    }

    s1 := grayCode(n - 1)

    s2 := make([]int, len(s1) * 2)
    copy(s2, s1)
    for i := 0; i < len(s1); i++ {
        s2[len(s1) + i] = 1 << uint(n - 1) + s1[len(s1) - 1 - i]
    }
    return s2
}
```

## 实现
镜像排列，实现这个算法即可，dp思想
```golang
func grayCode(n int) []int {
    if n < 0 {
        return []int{}
    }
    ret := []int{0}
    dp(&ret, uint(n), int(math.Pow(2, float64(n-1))))
    return ret
}

func dp(tmp *[]int, n uint, count int) {
    if n == 1 {
        *tmp = append(*tmp, 1)
    } else if n > 1 {
        // cal the front of n
        dp(tmp, n-1, count/2)
        // add array
        for i := count - 1; i >= 0; i-- {
            *tmp = append(*tmp, (*tmp)[i]+1<<(n-1))
        }
    }
}
```

## 改进
```golang

```

## 反思

## 参考