package climbing_stairs

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
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
                one: 2,
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: 3,
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: 1,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: 10,
            },
            a: result{
                one: 89,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, climbStairs(p.one), "输入:%v", q)
    }
}
