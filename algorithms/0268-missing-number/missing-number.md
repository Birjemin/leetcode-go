## 问题
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:
```
Input: [3,0,1]
Output: 2
```

Example 2:
```
Input: [9,6,4,2,3,5,7,0,1]
Output: 8
```

Note:

Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

## 分析
给定一个包含 `0, 1, 2, ..., n` 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

补值就好了, 136题和137题类似解法

## 最高分
这就很刘辟了
```golang
func missingNumber(nums []int) int {
	ret := 0
	for i,v := range nums {
		ret ^= v ^ (i + 1)
	}
	return ret
}
```

## 实现
```golang
func missingNumber(nums []int) int {
	length := len(nums)
	sum := length * (length + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
```

## 改进
异或（相同为0，这个很有趣~~）
```golang
func missingNumber(nums []int) int {
	var ret int
	for i, v := range nums {
		ret ^= (i+1)^v
	}
	return ret
}
```

## 改进
```golang
func missingNumber(nums []int) int {
	var ret int
	for i, v := range nums {
		ret += i + 1 - v
	}
	return ret
}
```

## 反思

## 参考