package remove_duplicates_from_sorted_array_ii

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
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
                one: []int{1, 1, 1, 2, 2, 3},
            },
            a: result{
                one: 5,
            },
        },
        {
            p: param{
                one: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
            },
            a: result{
                one: 7,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, removeDuplicates(p.one), "输入:%v", q)
    }
}
