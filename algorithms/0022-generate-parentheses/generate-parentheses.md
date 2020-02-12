## 问题
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:
```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

## 分析
n代表生成括号的对数，输出所有可能的排列组合
- 动态规划
- 深度优先遍历
- 广度优先遍历
需要考究一下

## 最高分
这个解法真的是刘辟！！！
```golang
func generateParenthesis(n int) []string {
    var ret []string
    backtrack(&ret, "", 0, 0, n)
    return ret
}

func backtrack(ans *[]string, cur string, open, close, max int) {
    if len(cur) == 2 * max {
        *ans = append(*ans, cur)
        return
    }
    if open < max {
        backtrack(ans, cur+"(", open+1, close, max)
    }
    if close < open {
        backtrack(ans, cur +")", open, close+1, max)
    }
}
```

## 实现
```golang
func generateParenthesis(n int) []string {
    var ret []string
    backtrack(&ret, "", 0, 0, n)
    return ret
}

func backtrack(ans *[]string, cur string, open, close, max int) {
    if len(cur) == 2 * max {
        *ans = append(*ans, cur)
        return
    }
    if open < max {
        backtrack(ans, cur+"(", open+1, close, max)
    }
    if close < open {
        backtrack(ans, cur +")", open, close+1, max)
    }
}
```

## 改进
```golang

```

## 反思

## 参考
