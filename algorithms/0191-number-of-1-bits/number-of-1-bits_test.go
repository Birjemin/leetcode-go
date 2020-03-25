package number_of_1_bits

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one uint32
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
                one: 11,
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: 128,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: 4294967293,
            },
            a: result{
                one: 31,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, hammingWeight(p.one), "输入:%v", p)
        ast.Equal(a.one, hammingWeight1(p.one), "输入:%v", p)
    }
}
