package evaluate_reverse_polish_notation

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []string
}

type result struct {
    one int
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
                one: []string{"2", "1", "+", "3", "*"},
            },
            a: result{
                one: 9,
            },
        },
        {
            p: param{
                one: []string{"4", "13", "5", "/", "+"},
            },
            a: result{
                one: 6,
            },
        },
        {
            p: param{
                one: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
            },
            a: result{
                one: 22,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, evalRPN(p.one), "输入:%v", q)
    }
}
