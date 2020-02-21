package spiral_matrix

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one [][]int
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
                one: [][]int{
                    {1},
                },
            },
            a: result{
                one: []int{1},
            },
        },
        {
            p: param{
                one: [][]int{
                    {7, 9, 6},
                },
            },
            a: result{
                one: []int{7, 9, 6},
            },
        },
        {
            p: param{
                one: [][]int{
                    {2, 5, 8},
                    {4, 0, -1},
                },
            },
            a: result{
                one: []int{2, 5, 8, -1, 0, 4},
            },
        },
        {
            p: param{
                one: [][]int{
                    {3},
                    {2},
                },
            },
            a: result{
                one: []int{3, 2},
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2, 3},
                    {4, 5, 6},
                    {7, 8, 9},
                    {10, 11, 12},
                },
            },
            a: result{
                one: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2, 3},
                    {4, 5, 6},
                    {7, 8, 9},
                },
            },
            a: result{
                one: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2, 3, 4},
                    {5, 6, 7, 8},
                    {9, 10, 11, 12},
                },
            },
            a: result{
                one: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2, 3, 4},
                    {5, 6, 7, 8},
                    {9, 10, 11, 12},
                    {13, 14, 15, 16},
                },
            },
            a: result{
                one: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2, 3, 4, 5},
                    {6, 7, 8, 9, 10},
                    {11, 12, 13, 14, 15},
                    {16, 17, 18, 19, 20},
                    {21, 22, 23, 24, 25},
                },
            },
            a: result{
                one: []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, spiralOrder(p.one), "输入:%v", q)
    }
}
