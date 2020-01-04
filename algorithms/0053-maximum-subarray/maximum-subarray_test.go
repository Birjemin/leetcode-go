package maximum_subarray

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
                one: []int{-1},
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
            },
            a: result{
                one: 6,
            },
        },
        {
            p: param{
                one: []int{1, -2, -1, 3, -1, 4},
            },
            a: result{
                one: 6,
            },
        },
        {
            p: param{
                one: []int{-1, -2, -3},
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{-2, 1},
            },
            a: result{
                one: 1,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, maxSubArray(p.one), "输入:%v", q)
        ast.Equal(a.one, maxSubArray1(p.one), "输入:%v", q)
    }
}