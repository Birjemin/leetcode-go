package two_sum_ii_input_array_is_sorted

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
    two int
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
                one: []int{2, 7, 11, 15},
                two: 9,
            },
            a: result{
                one: []int{1, 2},
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, twoSum(p.one, p.two), "输入:%v", p)
    }
}
