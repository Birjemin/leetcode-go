package search_in_rotated_sorted_array

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
                one: []int{4, 5, 6, 7, 0, 1, 2},
                two: 0,
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: []int{4, 5, 6, 7, 0, 1, 2},
                two: 3,
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{1},
                two: 1,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: []int{1, 3},
                two: 1,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: []int{1, 3},
                two: 0,
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{1, 3, 5},
                two: 3,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{1, 3, 5},
                two: 5,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{4, 5, 6, 7, 0, 1, 2},
                two: 2,
            },
            a: result{
                one: 6,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, search(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, search1(p.one, p.two), "输入:%v", q)
    }
}
