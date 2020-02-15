package find_first_and_last_position_of_element_in_sorted_array

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
                one: []int{5, 7, 7, 8, 8, 10},
                two: 8,
            },
            a: result{
                one: []int{3, 4},
            },
        },
        {
            p: param{
                one: []int{5, 7, 7, 8, 8, 10},
                two: 6,
            },
            a: result{
                one: []int{-1, -1},
            },
        },
        {
            p: param{
                one: []int{1, 2, 5, 5, 5, 9},
                two: 5,
            },
            a: result{
                one: []int{2, 4},
            },
        },
        {
            p: param{
                one: []int{5, 5, 5, 5, 5},
                two: 5,
            },
            a: result{
                one: []int{0, 4},
            },
        },
        {
            p: param{
                one: []int{},
                two: 0,
            },
            a: result{
                one: []int{-1, -1},
            },
        },
        {
            p: param{
                one: []int{1},
                two: 0,
            },
            a: result{
                one: []int{-1, -1},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, searchRange(p.one, p.two), "输入:%v", q)
    }
}
