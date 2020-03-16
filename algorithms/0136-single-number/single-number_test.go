package single_number

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
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
                one: []int{2, 2, 1},
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{4, 1, 2, 1, 2},
            },
            a: result{
                one: 4,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, singleNumber(p.one), "输入:%v", q)
        ast.Equal(a.one, singleNumber1(p.one), "输入:%v", q)
    }
}
