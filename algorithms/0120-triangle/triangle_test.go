package triangle

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
                    {2},
                    {3, 4},
                    {6, 5, 7},
                    {4, 1, 8, 3},
                },
            },
            a: result{
                one: 11,
            },
        },
        {
            p: param{
                one: [][]int{
                    {-7},
                    {-2, 1},
                    {-5, -5, 9},
                    {-4, -5, 4, 4},
                    {-6, -6, 2, -1, -5},
                    {3, 7, 8, -3, 7, -9},
                    {-9, -1, -9, 6, 9, 0, 7},
                    {-7, 0, -6, -8, 7, 1, -4, 9},
                    {-3, 2, -6, -9, -7, -6, -9, 4, 0},
                    {-8, -6, -3, -9, -2, -6, 7, -5, 0, 7},
                    {-9, -1, -2, 4, -2, 4, 4, -1, 2, -5, 5},
                    {1, 1, -6, 1, -2, -4, 4, -2, 6, -6, 0, 6},
                    {-3, -3, -6, -2, -6, -2, 7, -9, -5, -7, -5, 5, 1},
                },
            },
            a: result{
                one: -63,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, minimumTotal(p.one), "输入:%v", q)
    }
}
