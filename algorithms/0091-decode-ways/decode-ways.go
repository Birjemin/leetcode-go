package decode_ways

import (
    "strconv"
)

func numDecodings(s string) int {
    var count int
    cal(&count, s)
    return count
}

func cal(count *int, s string) {
    if s == "" {
        *count += 1
        return
    } else if s[0] == '0' {
        return
    } else if len(s) == 1 {
        *count += 1
        return
    }
    num, _ := strconv.Atoi(s[0:2])
    if num <= 26 {
        cal(count, s[2:])
    }
    cal(count, s[1:])
}

func numDecodings1(s string) int {
    if len(s) == 0 {
        return 0
    }
    length := len(s)
    dp := make([]int, length+1)
    dp[0] = 1
    if s[0] == '0' {
        dp[1] = 0
    } else {
        dp[1] = 1
    }
    for i := 2; i <= length; i++ {
        lastNum, _ := strconv.Atoi(s[i-1 : i])
        if lastNum >= 1 && lastNum <= 9 {
            dp[i] += dp[i-1]
        }
        lastNum, _ = strconv.Atoi(s[i-2 : i])
        if lastNum >= 10 && lastNum <= 26 {
            dp[i] += dp[i-2]
        }
    }
    return dp[len(s)]
}
