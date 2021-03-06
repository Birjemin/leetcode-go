package climbing_stairs

func climbStairs1(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    one, two := 1, 2
    for i := 3; i <= n; i++ {
        one, two = two, one+two
    }
    return two
}
