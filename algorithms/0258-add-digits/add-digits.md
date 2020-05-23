## 问题
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:
```
Input: 38
Output: 2 
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2. 
             Since 2 has only one digit, return it.
```

Follow up:
Could you do it without any loop/recursion in O(1) runtime?

## 分析
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

O(1) 时间复杂度内解决。

## 最高分
```golang
func addDigits(num int) int {
    if num < 10 {
        return num
    }
    
    if r := num % 9; r == 0 {
        return 9
    } else {
        return r
    }
}
```

## 实现
找规律的题目
```golang
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if ret := num % 9; ret == 0 {
		return 9
	} else {
		return ret
	}
}
```

## 改进
刘辟
```golang
func addDigits(num int) int {
	return (num-1)%9 + 1
}
```

## 反思

## 参考