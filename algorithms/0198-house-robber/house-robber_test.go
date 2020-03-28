package house_robber

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
                one: []int{1, 2, 3, 1},
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: []int{2, 7, 9, 3, 1},
            },
            a: result{
                one: 12,
            },
        },
        {
            p: param{
                one: []int{1, 2},
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{},
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: []int{0},
            },
            a: result{
                one: 0,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, rob(p.one), "输入:%v", p)
        ast.Equal(a.one, rob1(p.one), "输入:%v", p)
    }
}
