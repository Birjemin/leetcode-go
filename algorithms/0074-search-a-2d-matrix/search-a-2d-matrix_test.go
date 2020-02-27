package search_a_2d_matrix

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one [][]int
    two int
}

type result struct {
    one bool
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
                one: [][]int{{}},
                two: 1,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: [][]int{{1, 3}},
                two: 1,
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 3, 5, 7},
                    {10, 11, 16, 20},
                    {23, 30, 34, 50},
                },
                two: 3,
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 3, 5, 7},
                    {10, 11, 16, 20},
                    {23, 30, 34, 50},
                },
                two: 13,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 3, 5, 7, 8},
                    {10, 11, 16, 20, 21},
                    {23, 30, 34, 50, 100},
                },
                two: 16,
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, searchMatrix(p.one, p.two), "输入:%v", q)
    }
}
