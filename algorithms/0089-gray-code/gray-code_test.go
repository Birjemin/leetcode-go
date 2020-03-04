package gray_code

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
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
                one: 1,
            },
            a: result{
                one: []int{0, 1},
            },
        },
        {
            p: param{
                one: 2,
            },
            a: result{
                one: []int{0, 1, 3, 2},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, grayCode(p.one), "输入:%v", q)
    }
}
