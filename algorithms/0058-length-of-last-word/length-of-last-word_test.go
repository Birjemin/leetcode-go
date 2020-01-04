package length_of_last_word

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
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
                one: "Hello World",
            },
            a: result{
                one: 5,
            },
        },
        {
            p: param{
                one: "   ddd",
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: "   ddd   ",
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: "",
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, lengthOfLastWord(p.one), "输入:%v", q)
        ast.Equal(a.one, lengthOfLastWord1(p.one), "输入:%v", q)
    }
}