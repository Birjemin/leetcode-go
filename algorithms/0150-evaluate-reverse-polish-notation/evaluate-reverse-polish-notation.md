## 问题
Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:

Division between two integers should truncate toward zero.
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
Example 1:
```
Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
```

Example 2:
```
Input: ["4", "13", "5", "/", "+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
```

Example 3:
```
Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
Output: 22
Explanation: 
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

## 分析
根据逆波兰表示法，求表达式的值。（栈的使用）

## 最高分
```golang
func evalRPN(tokens []string) int {
    var s []int
    for i := 0; i < len(tokens); i++ {
        switch ch := tokens[i]; ch {
        case "+", "-", "*", "/":
            int1, int2 := s[len(s)-2], s[len(s)-1]
            s = s[:len(s)-2]
            
            var res int
            if ch == "+" {
                res = int1 + int2
            } else if ch == "-" {
                res = int1 - int2
            } else if ch == "*" {
                res = int1 * int2
            } else {
                res = int1 / int2
            }
            s = append(s, res)
        default:
            i, _ := strconv.Atoi(ch)
            s = append(s, i)
        }
    }
    
    return s[0] // should be only one element left as long as the input is valid
}
```

## 实现
栈的使用
```golang
func evalRPN(tokens []string) int {
    var tmp []int
    var length int
    for _, v := range tokens {
        // check the string is number
        if num, err := strconv.Atoi(v); err == nil {
            tmp = append(tmp, num)
            length++
        } else {
            operate(&tmp, &length, v)
            length--
        }
    }
    return tmp[0]
}

func operate(ret *[]int, length *int, operate string) {
    num := (*ret)[*length-1]
    *ret = (*ret)[:*length-1]
    switch operate {
    case "+":
        (*ret)[*length-2] += num
    case "-":
        (*ret)[*length-2] -= num
    case "*":
        (*ret)[*length-2] *= num
    case "/":
        (*ret)[*length-2] /= num
    }
}
```

## 改进
```golang

```

## 反思

## 参考