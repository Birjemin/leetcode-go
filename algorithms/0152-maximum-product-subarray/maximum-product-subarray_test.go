package maximum_product_subarray

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
                one: []int{2, 3, -2, 4},
            },
            a: result{
                one: 6,
            },
        },
        {
            p: param{
                one: []int{-2, 0, -1},
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: []int{2, -3, -2, 0, 3},
            },
            a: result{
                one: 12,
            },
        },
        {
            p: param{
                one: []int{-4, -3},
            },
            a: result{
                one: 12,
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
        {
            p: param{
                one: []int{-1},
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{3, -1, 4},
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: []int{2, -5, -2, -4, 3},
            },
            a: result{
                one: 24,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, maxProduct(p.one), "输入:%v", q)
        ast.Equal(a.one, maxProduct1(p.one), "输入:%v", q)
        ast.Equal(a.one, maxProduct2(p.one), "输入:%v", q)
    }
}
