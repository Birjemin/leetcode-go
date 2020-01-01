package search_insert_position

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
                one: []int{1, 3, 5, 6},
                two: 5,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: []int{1, 3, 5, 6},
                two: 2,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: []int{1, 3, 5, 6},
                two: 7,
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: []int{1, 3, 5, 6},
                two: 0,
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, searchInsert(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, searchInsert1(p.one, p.two), "输入:%v", q)
    }
}
