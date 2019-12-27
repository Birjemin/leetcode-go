package valid_parentheses

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
}

type result struct {
    one bool
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
                one: "()",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "()[]{}",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "([)]",
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: "{[]}",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "(){[({[]})]}",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "((([[{{{",
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: "){",
            },
            a: result{
                one: false,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, isValid1(p.one), "输入:%v", q)
        ast.Equal(a.one, isValid2(p.one), "输入:%v", q)
    }
}
