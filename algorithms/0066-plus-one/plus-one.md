## 问题
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:
```
Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

Example 2:
```
Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

## 分析
只要注意处理进位问题就好了

## 最高分
好像也没啥简单的
```golang
func plusOne(digits []int) []int {
	temp := digits[len(digits)-1] + 1
	if temp < 10 {
		digits[len(digits)-1] = temp
		return digits
	}
	carry := 1
	digits[len(digits)-1] = 0
	for i := len(digits) - 2; i >= 0; i-- {
		temp = digits[i] + carry
		if temp < 10 {
			digits[i] = temp
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append(digits[:0], append([]int{1}, digits[0:]...)...)
}
```


## 实现
最直接的方式
```golang
func plusOne(digits []int) []int {
    // 处理进位
    for i := len(digits) - 1; i >= 0; i-- {
        // 当前位不会产生进位
        if digits[i] < 9 {
            // 如果i = len(digits) - 1则代表+1，反之代表进位1
            digits[i]++
            return digits
        }
        // 如果digits[i] == 9 则产生进位1，当前位为0
        digits[i] = 0
    }
    // 处理首部进位
    return append([]int{1}, digits...)
}
```

## 改进
```golang

```

## 反思

## 参考
