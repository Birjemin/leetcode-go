package roman_to_integer

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
                one: "III",
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: "IV",
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: "IX",
            },
            a: result{
                one: 9,
            },
        },
        {
            p: param{
                one: "LVIII",
            },
            a: result{
                one: 58,
            },
        },
        {
            p: param{
                one: "MCMXCIV",
            },
            a: result{
                one: 1994,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, romanToInt1(p.one), "输入:%v", q)
        ast.Equal(a.one, romanToInt2(p.one), "输入:%v", q)
    }
}
