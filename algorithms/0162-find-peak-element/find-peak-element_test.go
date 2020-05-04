package find_peak_element

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
                one: []int{1, 2, 3, 1},
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{1, 2, 1, 3, 5, 6, 4},
            },
            a: result{
                one: 5,
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
                one: []int{1},
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        // ast.Equal(a.one, findPeakElement(p.one), "输入:%v", q)
        ast.Equal(a.one, findPeakElement1(p.one), "输入:%v", q)
    }
}
