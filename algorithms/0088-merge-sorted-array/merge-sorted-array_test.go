package merge_sorted_array

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one   []int
    two   int
    three []int
    four  int
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
                one:   []int{1, 2, 3, 0, 0, 0},
                two:   3,
                three: []int{2, 5, 6},
                four:  3,
            },
            a: result{
                one: []int{1, 2, 2, 3, 5, 6},
            },
        },
        {
            p: param{
                one:   []int{1, 3, 5, 7, 0, 0},
                two:   4,
                three: []int{2, 4},
                four:  2,
            },
            a: result{
                one: []int{1, 2, 3, 4, 5, 7},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        merge(p.one, p.two, p.three, p.four)
        ast.Equal(a.one, p.one, "输入:%v", q)
    }
}
