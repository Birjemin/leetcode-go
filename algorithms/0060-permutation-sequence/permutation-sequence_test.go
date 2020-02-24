package permutation_sequence

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
    two int
}

type result struct {
    one string
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
                one: 3,
                two: 3,
            },
            a: result{
                one: "213",
            },
        },
        {
            p: param{
                one: 4,
                two: 9,
            },
            a: result{
                one: "2314",
            },
        },
        {
            p: param{
                one: 4,
                two: 6,
            },
            a: result{
                one: "1432",
            },
        },
        {
            p: param{
                one: 2,
                two: 2,
            },
            a: result{
                one: "21",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, getPermutation(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, getPermutation1(p.one, p.two), "输入:%v", q)
    }
}
