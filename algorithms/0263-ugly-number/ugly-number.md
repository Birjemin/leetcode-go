## 问题
Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example 1:
```
Input: 6
Output: true
Explanation: 6 = 2 × 3
```

Example 2:
```
Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2
```

Example 3:
```
Input: 14
Output: false 
Explanation: 14 is not ugly since it includes another prime factor 7.
Note:
```

- 1 is typically treated as an ugly number.
- Input is within the 32-bit signed integer range: `[−2^31,  2^31 − 1]`.

## 分析
编写一个程序判断给定的数是否为丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

n = 2^x + 3^y +5^z(x、y、z有解)，只要不停求解，如果不能整除，则说明有其他质因子

## 最高分
```golang
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 || num == 2 || num == 3 || num == 5 {
		return true
	}
	if num%2 == 0 {
		return isUgly(num / 2)
	}
	if num%3 == 0 {
		return isUgly(num / 3)
	}
	if num%5 == 0 {
		return isUgly(num / 5)
	}
	return false
}
```

## 实现
超时(超时的原因是num == 0时超时。。。。犯了低级错误)
```golang
func isUgly(num int) bool {
	if num < 0 {
		return false
	}

	for {
		if num == 1 {
			return true
		}
		a1 := num % 2
		a2 := num % 3
		a3 := num % 5
		if a1 != 0 && a2 != 0 && a3 != 0 {
			return false
		}
		if a1 ==0 {
			num /= 2
		}
		if a2 == 0 {
			num /= 3
		}
		if a3 == 0 {
			num /= 5
		}
	}
}
```

## 改进
递归
```golang
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	}
	if num%3 == 0 {
		return isUgly(num / 3)
	}
	if num%5 == 0 {
		return isUgly(num / 5)
	}
	return num == 1
}

```

## 改进
循环
```golang
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1
}
```
## 反思

## 参考