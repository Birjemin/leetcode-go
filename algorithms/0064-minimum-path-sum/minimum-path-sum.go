package minimum_path_sum

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
