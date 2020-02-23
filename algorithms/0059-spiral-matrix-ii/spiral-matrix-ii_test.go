package spiral_matrix_ii

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
}

type result struct {
    one [][]int
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
                one: 3,
            },
            a: result{
                one: [][]int{
                    {1, 2, 3},
                    {8, 9, 4},
                    {7, 6, 5},
                },
            },
        },
        {
            p: param{
                one: 4,
            },
            a: result{
                one: [][]int{
                    {1, 2, 3, 4},
                    {12, 13, 14, 5},
                    {11, 16, 15, 6},
                    {10, 9, 8, 7},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, generateMatrix(p.one), "输入:%v", q)
    }
}
