package unique_binary_search_trees

func numTrees(n int) int {
    dp := make([]int, n)
    cal(&dp, n)
    return dp[n-1]
}

func cal(d *[]int, n int) {
    if n == 1 {
        (*d)[n-1] = 1
    } else {
        cal(d, n-1)
        (*d)[n-1] = (4*n - 2) * (*d)[n-2] / (n + 1)
    }
}
