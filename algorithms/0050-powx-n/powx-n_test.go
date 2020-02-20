package powx_n

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one float64
    two int
}

type result struct {
    one float64
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
                one: 2.00000,
                two: 10,
            },
            a: result{
                one: 1024.00000,
            },
        },
        {
            p: param{
                one: 2.10000,
                two: 3,
            },
            a: result{
                one: 9.261000000000001,
            },
        },
        {
            p: param{
                one: 2.00000,
                two: -2,
            },
            a: result{
                one: 0.25000,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, myPow(p.one, p.two), "输入:%v", q)
    }
}
