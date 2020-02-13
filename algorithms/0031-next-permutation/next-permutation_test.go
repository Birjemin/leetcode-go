package next_permutation

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
}

type result struct {
    one []int
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
                one: []int{1, 3, 2},
            },
        },
        {
            p: param{
                one: []int{3, 2, 1},
            },
            a: result{
                one: []int{1, 2, 3},
            },
        },
        {
            p: param{
                one: []int{1, 1, 5},
            },
            a: result{
                one: []int{1, 5, 1},
            },
        },
        {
            p: param{
                one: []int{1, 2, 7, 4, 3, 1},
            },
            a: result{
                one: []int{1, 3, 1, 2, 4, 7},
            },
        },
        {
            p: param{
                one: []int{1, 2},
            },
            a: result{
                one: []int{2, 1},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        nextPermutation(p.one)
        ast.Equal(a.one, p.one, "输入:%v", q)
    }
}
