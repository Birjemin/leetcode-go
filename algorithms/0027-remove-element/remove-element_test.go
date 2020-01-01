package remove_element

import (
"github.com/stretchr/testify/assert"
"testing"
)

type param struct {
    one []int
    two int
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
                one: []int{3, 2, 2, 3},
                two: 3,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{0,1,2,2,3,0,4,2},
                two: 2,
            },
            a: result{
                one: 5,
            },
        },
        {
            p: param{
                one: []int{1, 2},
                two: 1,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{0, 0, 0, 0, 0},
                two: 0,
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, removeElement1(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, removeElement(p.one, p.two), "输入:%v", q)
    }
}
