package majority_element

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
                one: []int{3, 2, 3},
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: []int{2, 2, 1, 1, 1, 2, 2},
            },
            a: result{
                one: 2,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, majorityElement(p.one), "输入:%v", p)
        // ast.Equal(a.one, majorityElement1(p.one), "输入:%v", p)
        ast.Equal(a.one, majorityElement2(p.one), "输入:%v", p)
    }
}
