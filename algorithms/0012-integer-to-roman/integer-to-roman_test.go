package integer_to_roman

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
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
                one: 3,
            },
            a: result{
                one: "III",
            },
        },
        {
            p: param{
                one: 4,
            },
            a: result{
                one: "IV",
            },
        },
        {
            p: param{
                one: 9,
            },
            a: result{
                one: "IX",
            },
        },
        {
            p: param{
                one: 58,
            },
            a: result{
                one: "LVIII",
            },
        },
        {
            p: param{
                one: 1994,
            },
            a: result{
                one: "MCMXCIV",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, intToRoman(p.one), "输入:%v", q)
    }
}