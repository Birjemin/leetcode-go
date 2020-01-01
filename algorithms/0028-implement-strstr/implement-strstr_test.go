package implement_strstr

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
    two string
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
                one: "hello",
                two: "ll",
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: "aaaaa",
                two: "bba",
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: "aaaaa",
                two: "",
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: "a",
                two: "a",
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, strStr(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, strStr1(p.one, p.two), "输入:%v", q)
    }
}