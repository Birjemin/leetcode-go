package unique_paths_ii

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
