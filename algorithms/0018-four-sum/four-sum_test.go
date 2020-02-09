package four_sum

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
                one: []int{1, 0, -1, 0, -2, 2},
                two: 0,
            },
            a: result{
                one: [][]int{
                    {-2, -1, 1, 2},
                    {-2, 0, 0, 2},
                    {-1, 0, 0, 1},
                },
            },
        },
        {
            p: param{
                one: []int{0, 0, 0, 0},
                two: 0,
            },
            a: result{
                one: [][]int{
                    {0, 0, 0, 0},
                },
            },
        },
        {
            p: param{
                one: []int{-1, 0, 1, 2, -1, -4},
                two: -1,
            },
            a: result{
                one: [][]int{
                    {-4, 0, 1, 2},
                    {-1, -1, 0, 1},
                },
            },
        },
        {
            p: param{
                one: []int{-2, -1, 0, 0, 1, 2, 3, 4, 5},
                two: 0,
            },
            a: result{
                one: [][]int{
                    {-2, -1, 0, 3},
                    {-2, -1, 1, 2},
                    {-2, 0, 0, 2},
                    {-1, 0, 0, 1},
                },
            },
        },
        {
            p: param{
                one: []int{1, -2, -5, -4, -3, 3, 3, 5},
                two: -11,
            },
            a: result{
                one: [][]int{
                    {-5, -4, -3, 1},
                },
            },
        },
        {
            p: param{
                one: []int{-1, -5, -5, -3, 2, 5, 0, 4},
                two: -7,
            },
            a: result{
                one: [][]int{
                    {-5, -5, -1, 4},
                    {-5, -3, -1, 2},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        // ast.Equal(a.one, fourSum(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, fourSum1(p.one, p.two), "输入:%v", q)
    }
}
