## 问题
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:
```
Input: "A man, a plan, a canal: Panama"
Output: true
```

Example 2:
```
Input: "race a car"
Output: false
```

## 分析
判断是否是回文字符串，忽略符号和大小写

## 最高分
```golang
func isPalindrome(s string) bool {
    l := len(s)
    if l == 0 {
        return true
    }
    head := 0
    tail := l - 1
    for head < tail {
        s1 := convert(rune(s[head]))
        s2 := convert(rune(s[tail]))
        for s1 == ' ' {
            head++
            if head >= tail {
                return true
            }
            s1 = convert(rune(s[head]))
        }
        for s2 == ' ' {
            tail--
            if head >= tail {
                return true
            }
            s2 = convert(rune(s[tail]))
        }
        if s1 != s2 {
            return false
        }
        head++
        tail--
    }
    return true
}

func convert(s rune) rune {
    if s >= 'a' && s <= 'z' || s >= '0' && s <= '9' {
        return s
    }
    if s >= 'A' && s <= 'Z' {
        return s + 'a' - 'A'
    }
    return ' '
}
```

## 实现
最基本的实现
```golang
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    for left, right := 0, len(s)-1; left < right; {
        if isNotString(s[left]) {
            left++
            continue
        }
        if isNotString(s[right]) {
            right--
            continue
        }
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func isNotString(s byte) bool {
    if s >= 'a' && s <= 'z' || s >= '0' && s <= '9' {
        return false
    }
    return true
}

```

## 改进
不使用strings库方法
```golang
func isPalindrome(s string) bool {
    for left, right := 0, len(s)-1; left < right; {
        // to lower
        l, r := toLower(s[left]), toLower(s[right])

        if isNotString(l) {
            left++
            continue
        }
        if isNotString(r) {
            right--
            continue
        }

        if l != r {
            return false
        }
        left++
        right--
    }

    return true
}

func isNotString(s byte) bool {
    if s >= 'a' && s <= 'z' || s >= '0' && s <= '9' {
        return false
    }
    return true
}

func toLower(s byte) byte {
    if s >= 'A' && s <= 'Z' {
        return s + 32
    }
    return s
}

```

## 反思

## 参考