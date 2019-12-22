## 问题

Given a 32-bit signed integer, reverse digits of an integer.

```
Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
```

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

## 分析
简单

## 最高分
无
```golang

```

## 实现
```golang
func reverse(x int) int {
	var res int
	for x != 0 {
		res = res * 10 + x % 10
		x = x / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

func reverse(x int) int {
	var res int
	for x != 0 {
		res = res * 10 + x % 10
		x = x / 10
        if res > math.MaxInt32 || res < math.MinInt32 {
		    return 0
	    }
	}
	return res
}
```

## 改进
减少内存的使用
```golang
func reverse(x int) int {
	var res int
	var tag bool
	if x < 0 {
		tag = true
		x = -x
	}
	for x != 0 {
		res = res * 10 + x % 10
		x = x / 10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}
	if tag {
		return -res
	}
	return res
}
```

## 反思
利用正负特性，在不降低性能的情况下，减少内存使用

## 参考
