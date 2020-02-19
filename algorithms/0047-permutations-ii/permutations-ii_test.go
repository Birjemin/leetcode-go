package permutations_ii

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
                one: []int{1, 1, 2},
            },
            a: result{
                one: [][]int{
                    {1, 1, 2},
                    {1, 2, 1},
                    {2, 1, 1},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, permuteUnique(p.one), "输入:%v", q)
    }
}
