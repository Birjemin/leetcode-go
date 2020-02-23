package merge_intervals

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
                    {1, 3},
                    {2, 6},
                    {8, 10},
                    {15, 18},
                },
            },
            a: result{
                one: [][]int{
                    {1, 6},
                    {8, 10},
                    {15, 18},
                },
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 4},
                    {4, 5},
                },
            },
            a: result{
                one: [][]int{
                    {1, 5},
                },
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 4},
                    {0, 4},
                },
            },
            a: result{
                one: [][]int{
                    {0, 4},
                },
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 4},
                    {0, 0},
                },
            },
            a: result{
                one: [][]int{
                    {0, 0},
                    {1, 4},
                },
            },
        },
        {
            p: param{
                one: [][]int{
                    {2,3},
                    {4,5},
                    {6,7},
                    {8,9},
                    {1,10},
                },
            },
            a: result{
                one: [][]int{
                    {1, 10},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, merge(p.one), "输入:%v", q)
    }
}
