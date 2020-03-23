package rotate_array

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
                one: []int{1, 2, 3, 4, 5, 6, 7},
                two: 3,
            },
            a: result{
                one: []int{5, 6, 7, 1, 2, 3, 4},
            },
        },
        {
            p: param{
                one: []int{-1, -100, 3, 99},
                two: 2,
            },
            a: result{
                one: []int{3, 99, -1, -100},
            },
        },
        {
            p: param{
                one: []int{1, 2},
                two: 3,
            },
            a: result{
                one: []int{2, 1},
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        // rotate(p.one, p.two)
        rotate1(p.one, p.two)
        ast.Equal(a.one, p.one, "输入:%v", p)
    }
}
