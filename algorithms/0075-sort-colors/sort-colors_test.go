package sort_colors

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
                one: []int{2, 0, 2, 1, 1, 0},
            },
            a: result{
                one: []int{0, 0, 1, 1, 2, 2},
            },
        },
        {
            p: param{
                one: []int{2, 0, 2, 1, 1},
            },
            a: result{
                one: []int{0, 1, 1, 2, 2},
            },
        },
        {
            p: param{
                one: []int{2, 0, 1},
            },
            a: result{
                one: []int{0, 1, 2},
            },
        },
        {
            p: param{
                one: []int{1, 2, 0},
            },
            a: result{
                one: []int{0, 1, 2},
            },
        },
        {
            p: param{
                one: []int{1, 2},
            },
            a: result{
                one: []int{1, 2},
            },
        },
        {
            p: param{
                one: []int{2, 1, 2},
            },
            a: result{
                one: []int{1, 2, 2},
            },
        },
        {
            p: param{
                one: []int{1, 2, 0, 0},
            },
            a: result{
                one: []int{0, 0, 1, 2},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        sortColors(p.one)
        ast.Equal(a.one, p.one, "输入:%v", q)
    }
}
