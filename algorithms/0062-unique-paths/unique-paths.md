## 问题
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Above is a 7 x 3 grid. How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:
```
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
```

Example 2:
```
Input: m = 7, n = 3
Output: 28
```

## 分析
f(m,n) = f(m-1,n)+f(m,n-1)

下面的方式超时(重复计算导致的，可以将计算结果保存下来，然后直接加)：
```
func uniquePaths(m int, n int) int {
    if m == 1 || n==1 {
        return 1
    }
    return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

```
## 最高分
```golang
func uniquePaths(m int, n int) int {
    cur, prev := make([]int, m), make([]int, m)
    for i := 0; i < m; i++ {
        cur[i], prev[i] = 1, 1
    }
    
    for j := 1; j < n; j++ {
        for i := 1; i < m; i++ {
            cur[i] = cur[i-1] + prev[i]
        }
        cur, prev = prev, cur
    }
    return prev[m-1]
}
```

## 实现
```golang
func uniquePaths(m int, n int) int {
    // allocate space of array
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }

    // default value
    for i := 0; i < m; i++ {
        dp[0][i] = 1
    }
    for i := 0; i < n; i++ {
        dp[i][0] = 1
    }

    // calculate each element of array
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[i][j] = dp[i - 1][j] + dp[i][j-1]
        }
    }
    return dp[n-1][m-1]
}
```

## 改进
合并循环
```golang
func uniquePaths(m int, n int) int {
    // allocate space of array
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            // default value
            if i == 0 || j == 0 {
                dp[i][j] = 1
            } else {
                // calculate each element of array
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }
    return dp[n-1][m-1]
}
```

## 反思

## 参考