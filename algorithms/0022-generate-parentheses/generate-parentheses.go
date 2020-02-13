package generate_parentheses

func generateParenthesis(n int) []string {
    var ret []string
    backtrack(&ret, "", 0, 0, n)
    return ret
}

func backtrack(ans *[]string, cur string, open, close, max int) {
    if len(cur) == max*2 {
        *ans = append(*ans, cur)
        return
    }
    if open < max {
        backtrack(ans, cur+"(", open+1, close, max)
    }
    if close < open {
        backtrack(ans, cur+")", open, close+1, max)
    }
}
