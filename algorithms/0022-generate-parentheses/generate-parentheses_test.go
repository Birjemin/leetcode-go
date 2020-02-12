package generate_parentheses

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
}

type result struct {
    one []string
}

type question struct {
    p param
    a result
}

func Test(t *testing.T) {
    ast := assert.New(t)
    qs := []question{
        {
            p: param{
                one: 3,
            },
            a: result{
                one: []string{
                    "((()))",
                    "(()())",
                    "(())()",
                    "()(())",
                    "()()()",
                },
            },
        },
        {
            p: param{
                one: 3,
            },
            a: result{
                one: []string{
                    "((()))",
                    "(()())",
                    "(())()",
                    "()(())",
                    "()()()",
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, generateParenthesis(p.one), "输入:%v", q)
    }
}
