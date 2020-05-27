## 问题
Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

Example 1:
```
Input: [5,7]
Output: 4
```

Example 2:
```
Input: [0,1]
Output: 0
```

## 分析
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

最简单的方式遍历、效率低，说明是有规律的题目

## 最高分
```golang

```

## 实现
最简单的方式，效率最低
```golang
func rangeBitwiseAnd(m int, n int) int {
	ret := m
	for i := m + 1; i <= n; i++ {
		ret &= i
	}
	return ret
}
```

## 改进
找规律的题目，高位对其，若不相等，则后面全为0（从左向右查找）

[找到最高位](https://codeforwin.org/2016/01/c-program-to-get-highest-order-set-bit-of-number.html)
```golang

func rangeBitwiseAnd1(m int, n int) int {
	if m < 0 || n < 0 {
		return 0
	}

	var count int

	for i := 31; i > 0; i-- {
		// if high bit is 1, find it!
		if m>>uint(i)&1 == 1 {
			count = i
		}

		if n>>uint(i)&1 == 1 {
			// because n >= m, count <= i
			// if count < i, return 0
			if count < i {
				return 0
			}
			break
		}
	}

	var ret int
	for i := count; i >= 0; i-- {

		// get high bit
		a1 := n>>uint(i)
		if a1 != m>>uint(i) {
			break
		}
		tmp := a1 << uint(i)
		ret |= tmp
		// subtract the high bit
		n -= tmp
		m -= tmp
	}
	return ret
}
```

## 改进
合并循环（从左向右查找）
```golang
func rangeBitwiseAnd(m int, n int) int {
	if m < 0 || n < 0 {
		return 0
	}

	var ret int
	for i := 31; i >= 0; i-- {
		x := m >> uint(i)
		if x != n >> uint(i) {
			break
		}
		tmp := x << uint(i)
		ret |= tmp
		// subtract the high bit
		n -= tmp
		m -= tmp
	}
	return ret
}
```

## 改进
从低位开始查找，如果n==m，则说明之后的位都相等，直接跳出即可（从右向左查找）
```golang
func rangeBitwiseAnd(m int, n int) int {
	var ret uint
	for n != m {
		n >>= 1
		m >>= 1
		ret++
	}
	return n << ret
}
```

## 反思

## 参考