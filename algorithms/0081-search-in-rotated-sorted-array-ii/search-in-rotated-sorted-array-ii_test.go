package search_in_rotated_sorted_array_ii

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
    two int
}

type result struct {
    one bool
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
                one: []int{2, 5, 6, 0, 0, 1, 2},
                two: 0,
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: []int{2, 5, 6, 0, 0, 1, 2},
                two: 3,
            },
            a: result{
                one: false,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, search(p.one, p.two), "输入:%v", q)
    }
}
