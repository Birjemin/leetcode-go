package unique_paths

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
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[n-1][m-1]
}

func uniquePaths1(m int, n int) int {
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
