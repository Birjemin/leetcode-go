package valid_palindrome

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
                one: "A man, a plan, a canal: Panama",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "race a car",
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: "0p",
            },
            a: result{
                one: false,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, isPalindrome(p.one), "输入:%v", q)
    }
}
