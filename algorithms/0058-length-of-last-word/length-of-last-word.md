## 问题
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word (last word means the last appearing word if we loop from left to right) in the string.

If the last word does not exist, return 0.

Note: A word is defined as a maximal substring consisting of non-space characters only.

Example:
```
Input: "Hello World"
Output: 5
```

## 分析
其实就是从后往前计算，遇到空格，则返回即可

## 最高分
```golang
func lengthOfLastWord(s string) int {
    var count int
    // 去除首尾空格
    s = strings.TrimSpace(s)
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == 32 {
            break
        }
        count++
    }
    return count
}
```

## 实现
其实就是从后往前计算，遇到空格，则返回即可
```golang
func lengthOfLastWord(s string) int {
    var flag int
    for i := len(s) - 1; i >= 0; i-- {
        // 32为空格的ASCII码
        if s[i] == 32 {
            if flag > 0 {
                break
            }
        } else {
            flag++
        }
    }
    return flag
}
```

## 改进
去除首尾空格，然后再倒序计算
```golang
func lengthOfLastWord(s string) int {
    var count int
    // 去除首尾空格
    s = strings.TrimSpace(s)
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == 32 {
            break
        }
        count++
    }
    return count
}
```

## 反思
除去边界条件，再重新计算

## 参考
