package container_with_most_water

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
                one: []int{1,8,6,2,5,4,8,3,7},
            },
            a: result{
                one: 49,
            },
        },
        {
            p: param{
                one: []int{1,1},
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{1, 2},
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{1, 2, 4, 3},
            },
            a: result{
                one: 4,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, maxArea(p.one), "输入:%v", q)
    }
}