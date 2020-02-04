package zigzag_conversion

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
    two int
}

type result struct {
    one string
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
                one: "PAYPALISHIRING",
                two: 3,
            },
            a: result{
                one: "PAHNAPLSIIGYIR",
            },
        },
        {
            p: param{
                one: "PAYPALISHIRING",
                two: 4,
            },
            a: result{
                one: "PINALSIGYAHRPI",
            },
        },
        {
            p: param{
                one: "LEETCODEISHIRING",
                two: 3,
            },
            a: result{
                one: "LCIRETOESIIGEDHN",
            },
        },
        {
            p: param{
                one: "LEETCODEISHIRING",
                two: 4,
            },
            a: result{
                one: "LDREOEIIECIHNTSG",
            },
        },
        {
            p: param{
                one: "A",
                two: 1,
            },
            a: result{
                one: "A",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, convert(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, convert1(p.one, p.two), "输入:%v", q)
    }
}
