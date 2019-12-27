package valid_parentheses

func isValid1(s string) bool {
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

func isValid2(s string) bool {
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
