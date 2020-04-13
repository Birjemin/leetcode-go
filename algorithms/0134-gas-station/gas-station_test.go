package gas_station

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
    two []int
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
                one: []int{1, 2, 3, 4, 5},
                two: []int{3, 4, 5, 1, 2},
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: []int{2, 3, 4},
                two: []int{3, 4, 3},
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: []int{5, 1, 2, 3, 4},
                two: []int{4, 4, 1, 5, 1},
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: []int{5, 8, 2, 8},
                two: []int{6, 5, 6, 6},
            },
            a: result{
                one: 3,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, canCompleteCircuit1(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, canCompleteCircuit(p.one, p.two), "输入:%v", q)
    }
}
