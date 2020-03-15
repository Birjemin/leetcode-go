## 问题
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
```
Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

Follow up:

Coud you solve it without converting the integer to a string?

## 分析
最直接的方式就是整除，不断求解

## 最高分
```golang
func isPalindrome(x int) bool {
    if math.Signbit(float64(x)) {
        return false
    }
    rev := 0
    remainder := 0
    argz := x
    for argz > 0 {
        rev *= 10
        remainder = argz % 10
        rev += remainder
        argz /= 10
    }
    if rev == x {
        return true
    }
    return false
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x <= 9 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    
    var y int
    for x > y {
        r := x % 10
        x = x / 10
        y = y * 10 + r
        
        if x == y || x / 10 == y {
            return true
        }
    }
    return false
}
```


## 实现
```golang
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x < 9 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    var y, temp int
    for temp = x; temp != 0; temp /= 10 {
        y = 10 * y + temp % 10
    }
    if y == x {
        return true
    } else {
        return false
    }
}
```

## 改进
回文数字，只需要拆分到一半判断就好了，不需要全部拆分
```golang
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x < 10 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    for y := 0; x > y; {
        y = 10 * y + x % 10
        x /= 10
        // 只需要计算一半
        if x == y || x / 10 == y {
            return true
        }
    }
    return false
}
```
## 反思
无

## 参考