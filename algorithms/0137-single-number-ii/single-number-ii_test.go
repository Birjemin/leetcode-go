package single_number_ii

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
                one: []int{2, 2, 3, 2},
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: []int{0, 1, 0, 1, 0, 1, 99},
            },
            a: result{
                one: 99,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        // ast.Equal(a.one, singleNumber(p.one), "输入:%v", q)
        ast.Equal(a.one, singleNumber1(p.one), "输入:%v", q)
    }
}
