## 问题
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

```
Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
```

## 分析
栈的使用

## 最高分
```golang
func isValid(s string) bool {
    parenth_dict := make(map[rune]rune)
    
    parenth_dict['('] = ')'
    parenth_dict['{'] = '}'
    parenth_dict['['] = ']'
    
    aux := make([]rune, 0)
    
    for i:= 0; i < len(s); i++ {
       if s[i] == '{' || s[i] == '(' || s[i] == '[' {
           aux = append(aux, rune(s[i]))
       } else if len(aux) == 0 || parenth_dict[aux[len(aux) - 1]] != rune(s[i]) {
           return false;
       } else {
           aux = aux[:len(aux) - 1]
       }
    }
    return len(aux) == 0
}

func isValid(s string) bool {
    if s == "" {
        return true
    }
    
    tr := map[rune]rune{
        '}': '{',
        ')': '(',
        ']': '[',
    }
    
    stack := make([]rune, 0, 1)
    
    for _, ch := range s {
        switch ch {
        case 40,123,91:
            stack = append(stack, ch)
        case 41,125,93:    
            if len(stack) == 0 || stack[len(stack)-1] != tr[ch] {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        }
    }
    
    return len(stack) == 0
}
```


## 实现
直接的想法实现
```golang
func isValid(s string) bool {
    if s == "" {
        return true
    }
    length := len(s)
    if length % 2 != 0 {
        return false
    }

    table := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }

    stack := make([]rune, length)
    var top int
    for _, ch := range s {
        if ch == '(' || ch == '[' || ch == '{' {
            stack[top] = table[ch]
            top++
        } else {
            top--
            // top < 0说明字符串不是左括号右括号的顺序
            if top < 0 || stack[top] != ch {
                return false
            }
        }
    }
    return top == 0
}
```

## 改进
去掉top参数，和字符串判断，直接使用ascII码
```golang
func isValid(s string) bool {
    if s == "" {
        return true
    }
    length := len(s)
    if length%2 != 0 {
        return false
    }

    table := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }
    // 预分配
    stack := make([]rune, 0, length/2+1)
    for _, ch := range s {
        switch ch {
        case 40, 123, 91:
            stack = append(stack, table[ch])
        default:
            {
                l := len(stack)
                if l < 1 || stack[l-1] != ch {
                    return false
                }
                stack = stack[:l-1]
            }
        }
    }
    return len(stack) == 0
}
```

## 反思
利用ASCII码判断和节省变量，不断简化逻辑

## 参考