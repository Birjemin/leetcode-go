package reverse_bits

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one uint32
}

type result struct {
    one uint32
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
                one: 43261596,
            },
            a: result{
                one: 964176192,
            },
        },
        {
            p: param{
                one: 4294967293,
            },
            a: result{
                one: 3221225471,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, reverseBits(p.one), "输入:%v", p)
        ast.Equal(a.one, reverseBits1(p.one), "输入:%v", p)
        ast.Equal(a.one, reverseBits2(p.one), "输入:%v", p)
    }
}

