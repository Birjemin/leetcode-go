package minimum_path_sum

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one [][]int
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
                one: [][]int{
                    {1, 3, 1},
                    {1, 5, 1},
                    {4, 2, 1},
                },
            },
            a: result{
                one: 7,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2, 5},
                    {3, 2, 1},
                },
            },
            a: result{
                one: 6,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 2},
                    {5, 6},
                    {1, 1},
                },
            },
            a: result{
                one: 8,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 4, 8, 6, 2, 2, 1, 7},
                    {4, 7, 3, 1, 4, 5, 5, 1},
                    {8, 8, 2, 1, 1, 8, 0, 1},
                    {8, 9, 2, 9, 8, 0, 8, 9},
                    {5, 7, 5, 7, 1, 8, 5, 5},
                    {7, 0, 9, 4, 5, 6, 5, 6},
                    {4, 9, 9, 7, 9, 1, 9, 0},
                },
            },
            a: result{
                one: 47,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, minPathSum(p.one), "输入:%v", q)
    }
}
