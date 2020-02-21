package jump_game

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
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
                one: []int{2, 3, 1, 1, 4},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: []int{3, 2, 1, 0, 4},
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: []int{2, 0},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: []int{2, 0, 0},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: []int{0},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: []int{2, 5, 0, 0},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: []int{2, 2, 0, 2, 0},
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, canJump(p.one), "输入:%v", q)
    }
}
