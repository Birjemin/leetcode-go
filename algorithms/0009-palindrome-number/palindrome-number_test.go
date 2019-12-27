package palindrome_number

import (
"github.com/stretchr/testify/assert"
"testing"
)

type param struct {
    one int
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
                one: 121,
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: 12123,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: -121,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: -12123,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: 10,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: 9,
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, isPalindrome1(p.one), "输入:%v", q)
        ast.Equal(a.one, isPalindrome2(p.one), "输入:%v", q)
    }
}