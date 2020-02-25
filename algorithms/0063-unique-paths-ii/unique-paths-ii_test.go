package unique_paths_ii

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
                    {0, 0, 0},
                    {0, 1, 0},
                    {0, 0, 0},
                },
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: [][]int{
                    {0},
                },
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1},
                },
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1, 0},
                },
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: [][]int{
                    {1},
                    {0},
                },
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: [][]int{
                    {0},
                    {1},
                },
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: [][]int{
                    {0, 0},
                    {0, 1},
                },
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, uniquePathsWithObstacles(p.one), "输入:%v", q)
    }
}
