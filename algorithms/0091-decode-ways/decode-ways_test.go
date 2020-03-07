package decode_ways

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
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
                one: "12",
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: "226",
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: "0",
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: "101",
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: "00",
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, numDecodings(p.one), "输入:%v", q)
        ast.Equal(a.one, numDecodings1(p.one), "输入:%v", q)
    }
}
