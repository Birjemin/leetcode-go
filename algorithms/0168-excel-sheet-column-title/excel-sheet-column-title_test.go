package excel_sheet_column_title

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
                one: 1,
            },
            a: result{
                one: "A",
            },
        },
        {
            p: param{
                one: 28,
            },
            a: result{
                one: "AB",
            },
        },
        {
            p: param{
                one: 701,
            },
            a: result{
                one: "ZY",
            },
        },
        {
            p: param{
                one: 52,
            },
            a: result{
                one: "AZ",
            },
        },
        {
            p: param{
                one: 703,
            },
            a: result{
                one: "AAA",
            },
        },
        {
            p: param{
                one: 24568,
            },
            a: result{
                one: "AJHX",
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, convertToTitle(p.one), "输入:%v", p)
        ast.Equal(a.one, convertToTitle1(p.one), "输入:%v", p)
    }
}
