package combination_sum_ii

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
                one: []int{10, 1, 2, 7, 6, 1, 5},
                two: 8,
            },
            a: result{
                one: [][]int{
                    {1, 1, 6},
                    {1, 2, 5},
                    {1, 7},
                    {2, 6},
                },
            },
        },
        {
            p: param{
                one: []int{2, 5, 2, 1, 2},
                two: 5,
            },
            a: result{
                one: [][]int{
                    {1, 2, 2},
                    {5},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, combinationSum2(p.one, p.two), "输入:%v", q)
    }
}
