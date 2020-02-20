package rotate_image

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one [][]int
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
                one: [][]int{
                    {1, 2},
                    {3, 4},
                },
            },
            a: result{
                one: [][]int{
                    {3, 1},
                    {4, 2},
                },
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
                one: [][]int{
                    {7, 4, 1},
                    {8, 5, 2},
                    {9, 6, 3},
                },
            },
        },
        {
            p: param{
                one: [][]int{
                    {5, 1, 9, 11},
                    {2, 4, 8, 10},
                    {13, 3, 6, 7},
                    {15, 14, 12, 16},
                },
            },
            a: result{
                one: [][]int{
                    {15, 13, 2, 5},
                    {14, 3, 4, 1},
                    {12, 6, 8, 9},
                    {16, 7, 10, 11},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        // rotate(p.one)
        // ast.Equal(a.one, p.one, "输入:%v", q)
        rotate1(p.one)
        ast.Equal(a.one, p.one, "输入:%v", q)
    }
}
