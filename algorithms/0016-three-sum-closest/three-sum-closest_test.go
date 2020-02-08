package three_sum_closest

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
    two int
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
                one: []int{-1, 2, 1, -4},
                two: 1,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{0, 0, 0, 0},
                two: 1,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: []int{-1, 2, 1, -4},
                two: 1,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{-1, 2, 1, -4},
                two: 1,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{1, 1, -1, -1, 3},
                two: -1,
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{1, 1, 1, 0},
                two: -100,
            },
            a: result{
                one: 2,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, threeSumClosest(p.one, p.two), "输入:%v", q)
    }
}
