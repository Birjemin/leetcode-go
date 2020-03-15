package pascals_triangle

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
                one: 0,
            },
            a: result{
                one: [][]int{},
            },
        },
        {
            p: param{
                one: 3,
            },
            a: result{
                one: [][]int{
                    {1},
                    {1, 1},
                    {1, 2, 1},
                },
            },
        },
        {
            p: param{
                one: 5,
            },
            a: result{
                one: [][]int{
                    {1},
                    {1, 1},
                    {1, 2, 1},
                    {1, 3, 3, 1},
                    {1, 4, 6, 4, 1},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, generate(p.one), "输入:%v", q)
    }
}
