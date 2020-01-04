package plus_one

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
}

type result struct {
    one []int
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
                one: []int{1, 2, 3},
            },
            a: result{
                one: []int{1, 2, 4},
            },
        },
        {
            p: param{
                one: []int{4, 3, 2, 1},
            },
            a: result{
                one: []int{4, 3, 2, 2},
            },
        },
        {
            p: param{
                one: []int{0},
            },
            a: result{
                one: []int{1},
            },
        },
        {
            p: param{
                one: []int{9},
            },
            a: result{
                one: []int{1, 0},
            },
        },
        {
            p: param{
                one: []int{8, 9},
            },
            a: result{
                one: []int{9, 0},
            },
        },
        {
            p: param{
                one: []int{},
            },
            a: result{
                one: []int{1},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, plusOne(p.one), "输入:%v", q)
    }
}
