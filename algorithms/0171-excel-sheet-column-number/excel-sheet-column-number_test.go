package excel_sheet_column_number

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
                one: "A",
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: "AB",
            },
            a: result{
                one: 28,
            },
        },
        {
            p: param{
                one: "ZY",
            },
            a: result{
                one: 701,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, titleToNumber(p.one), "输入:%v", p)
        ast.Equal(a.one, titleToNumber1(p.one), "输入:%v", p)
    }
}

