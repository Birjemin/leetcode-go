package permutations

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
                one: []int{1, 2, 3},
            },
            a: result{
                one: [][]int{
                    {1, 2, 3},
                    {1, 3, 2},
                    {2, 1, 3},
                    {2, 3, 1},
                    {3, 1, 2},
                    {3, 2, 1},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, permute(p.one), "输入:%v", q)
    }
}
