## 问题
A message containing letters from A-Z is being encoded to numbers using the following mapping:
```
'A' -> 1
'B' -> 2
...
'Z' -> 26
```
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:
```
Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
```

Example 2:
```
Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
```

## 分析
给定一个只包含数字的非空字符串，请计算解码方法的总数。有种类似进制的转换
- 递归查找
- DP动态规划，得找到规律（利用各个阶段之间的关系，逐个求解，最终求得全局最优解）

## 最高分
DP算法
```golang
func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    }
    dp := make([]int, len(s)+1)
    dp[0] = 1
    if s[:1] == "0" {
        dp[1] = 0
    } else {
        dp[1] = 1
    }
    for i := 2; i <= len(s); i++ {
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
```

## 实现
最基本的方式实现，递归
```golang
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
```

## 改进
DP动态规划，首先得发现这个规律，然后使用空间获取时间。不过这个规律难找
```golang
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
```

## 反思

## 参考