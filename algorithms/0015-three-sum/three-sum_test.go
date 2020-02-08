package three_sum

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
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
                one: []int{-1, 0, 1, 2, -1, -4},
            },
            a: result{
                one: [][]int{
                    {-1, -1, 2},
                    {-1, 0, 1},
                },
            },
        },
        {
            p: param{
                one: []int{-1, 0, 1, 2, -1, -4},
            },
            a: result{
                one: [][]int{
                    {-1, -1, 2},
                    {-1, 0, 1},
                },
            },
        },
        {
            p: param{
                one: []int{0, 0, 0, 0},
            },
            a: result{
                one: [][]int{
                    {0, 0, 0},
                },
            },
        },
        {
            p: param{
                one: []int{-2, 0, 0, 2, 2},
            },
            a: result{
                one: [][]int{
                    {-2, 0, 2},
                },
            },
        },
        {
            p: param{
                one: []int{1, -1, -1, 0},
            },
            a: result{
                one: [][]int{
                    {-1, 0, 1},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, threeSum(p.one), "输入:%v", q)
    }
}
