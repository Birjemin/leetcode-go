package factorial_trailing_zeroes

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
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
                one: 3,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: 5,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: 30,
            },
            a: result{
                one: 7,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, trailingZeroes(p.one), "输入:%v", p)
        ast.Equal(a.one, trailingZeroes1(p.one), "输入:%v", p)
        ast.Equal(a.one, trailingZeroes2(p.one), "输入:%v", p)
    }
}

