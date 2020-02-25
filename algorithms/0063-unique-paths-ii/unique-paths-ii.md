## 问题
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:
```
Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
```

## 分析
和62题类似，不过这一题中有一个障碍物的概念

## 最高分
```golang

```

## 实现
遇到为1的，将其设置为0
```golang
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n := len(obstacleGrid)
    if n == 0 {
        return 0
    }
    m := len(obstacleGrid[0])

    // allocate space of array
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }

    // handle unexpected situations
    if obstacleGrid[n-1][m-1] == 1 {
        return 0
    }

    // default value
    for i := 0; i < m; i++ {
        if obstacleGrid[0][i] == 1 {
            break
        }
        dp[0][i] = 1
    }
    for i := 0; i < n; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }

    // calculate each element of array
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if obstacleGrid[i-1][j] == 1 {
                dp[i-1][j] = 0
            }
            if obstacleGrid[i][j-1] == 1 {
                dp[i][j-1] = 0
            }
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[n-1][m-1]
}
```

## 改进
```golang

```

## 反思

## 参考