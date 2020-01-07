package sqrt

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
                one: 4,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: 8,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: 0,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: 1,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: 2,
            },
            a: result{
                one: 1,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, mySqrt(p.one), "输入:%v", q)
        ast.Equal(a.one, mySqrt1(p.one), "输入:%v", q)
    }
}