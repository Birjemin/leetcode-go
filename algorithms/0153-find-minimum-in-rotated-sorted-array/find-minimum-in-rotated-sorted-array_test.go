package find_minimum_in_rotated_sorted_array

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
                one: []int{3, 4, 5, 1, 2},
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{4, 5, 6, 7, 0, 1, 2},
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: []int{1, 2},
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{1, 2, 3},
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{2, 3, 1},
            },
            a: result{
                one: 1,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, findMin(p.one), "输入:%v", q)
        ast.Equal(a.one, findMin1(p.one), "输入:%v", q)
    }
}
