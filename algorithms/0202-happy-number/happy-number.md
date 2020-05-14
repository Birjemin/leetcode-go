## 问题
Write an algorithm to determine if a number n is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Return True if n is a happy number, and False if not.

Example: 
```
Input: 19
Output: true
Explanation: 
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
```

## 分析
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。

也可使用141题目的解法

## 最高分
```golang
func isHappy(n int) bool {
    fast := nextNum(n);
    for n != 1 && fast != 1 {
        if n == fast {
            return false;
        }
        n = nextNum(n);
        fast = nextNum(nextNum(fast))
    }
    return true;
}

func nextNum(n int) int {
    ans := 0
    for n != 0 {
        num := n % 10
        ans += num * num;
        n = n / 10
    }
    return ans;
}
```

## 实现
找规律的题目，多举几个例子推演一下就好了
```golang
func isHappy(n int) bool {
	var sum, remainder int
	for {
		for {
			remainder = n % 10
			sum += remainder * remainder
			n /= 10

			if n == 0 {
				n, sum = sum, 0
				break
			}
		}

		if n == 1 {
			return true
		} else if n == 4 {
			return false
		}
	}
}
```

## 改进
可以使用
```golang

```

## 反思

## 参考