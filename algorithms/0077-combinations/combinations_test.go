package combinations

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
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
                one: 4,
                two: 2,
            },
            a: result{
                one: [][]int{
                    {1, 2},
                    {1, 3},
                    {1, 4},
                    {2, 3},
                    {2, 4},
                    {3, 4},
                },
            },
        },
        {
            p: param{
                one: 5,
                two: 4,
            },
            a: result{
                one: [][]int{
                    {1, 2, 3, 4},
                    {1, 2, 3, 5},
                    {1, 2, 4, 5},
                    {1, 3, 4, 5},
                    {2, 3, 4, 5},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, combine(p.one, p.two), "输入:%v", q)
    }
}
