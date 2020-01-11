## 问题
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:
```
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

Example 2:
```
Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

## 分析
f(n) = f(n-1)+f(n-2)，本来想要使用递归的，但是超时执行了。。。

## 最高分
```golang
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	cache := make([]int, n)
	cache[0], cache[1] = 1, 2

	for i := 2; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n-1]
}
```

## 实现
```golang
func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```

## 改进
使用计数
```golang
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
```

## 反思
递归不是万能大法~~，不过可以使用累加来解决

## 参考
