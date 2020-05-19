## 问题
Given an integer, write a function to determine if it is a power of two.

Example 1:
```
Input: 1
Output: true 
Explanation: 20 = 1
```

Example 2:
```
Input: 16
Output: true
Explanation: 24 = 16
```

Example 3:
```
Input: 218
Output: false
```

## 分析
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

题意解答即可

## 最高分
```golang
func isPowerOfTwo(num int) bool {
	return (num > 0 && ((num & (num - 1)) == 0))
}
```

## 实现
位移运算
```golang
func isPowerOfTwo(n int) bool {
	a := 1
	for a < n {
		a <<= 1
	}
	if n == a {
		return true
	}
	return false
}
```

## 改进
位移运算
```golang
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n % 2 == 1 {
			return false
		}
		n >>= 1
	}
	return true
}
```

## 反思

## 参考