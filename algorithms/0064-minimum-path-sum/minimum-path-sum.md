## 问题
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:
```
Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
```

## 分析
mxn的网格，找出一条最近的路，这题真的不会，没有思路。题解为逐个计算每一个值，然后得出结果。。。

## 最高分
```golang
func minPathSum(grid [][]int) int {
    var m int = len(grid)
    var n int = len(grid[0])
    for i := 1; i < m; i++{
        grid[i][0] += grid[i-1][0];
    }
    for j := 1; j < n; j++{
        grid[0][j] += grid[0][j-1];
    }
    for i:= 1; i < m; i++{
        for j:=1; j<n; j++{
            grid[i][j] += min(grid[i-1][j], grid[i][j-1]);
        }
    }
    return grid[m-1][n-1];
    
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 实现
计算每一个值，答案竟然没啥好方法，看来是想复杂了。。。
```golang
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j > 0 {
                // width
                grid[i][j] += grid[i][j-1]
            } else if j == 0 && i > 0 {
                // height
                grid[i][j] += grid[i-1][j]
            } else if j != 0 && i != 0 {
                grid[i][j] += min(grid[i-1][j], grid[i][j-1])
            }
        }
    }
    return grid[m-1][n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 改进
```golang

```

## 反思

## 参考