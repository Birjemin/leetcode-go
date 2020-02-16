package combination_sum

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
    two int
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
                one: []int{2, 3, 6, 7},
                two: 7,
            },
            a: result{
                one: [][]int{
                    {2, 2, 3},
                    {7},
                },
            },
        },
        {
            p: param{
                one: []int{2, 3, 5},
                two: 8,
            },
            a: result{
                one: [][]int{
                    {2, 2, 2, 2},
                    {2, 3, 3},
                    {3, 5},
                },
            },
        },
        {
            p: param{
                one: []int{8, 7, 4, 3},
                two: 11,
            },
            a: result{
                one: [][]int{
                    {3, 4, 4},
                    {3, 8},
                    {4, 7},
                },
            },
        },
        {
            p: param{
                one: []int{2, 3, 4},
                two: 14,
            },
            a: result{
                one: [][]int{
                    {2, 2, 2, 2, 2, 2, 2},
                    {2, 2, 2, 2, 2, 4},
                    {2, 2, 2, 2, 3, 3},
                    {2, 2, 2, 4, 4},
                    {2, 2, 3, 3, 4},
                    {2, 3, 3, 3, 3},
                    {2, 4, 4, 4},
                    {3, 3, 4, 4},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, combinationSum(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, combinationSum1(p.one, p.two), "输入:%v", q)
    }
}
