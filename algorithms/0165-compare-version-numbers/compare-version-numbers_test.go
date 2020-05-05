package compare_version_numbers

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
                one: "0.1",
                two: "1.1",
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: "1.0.1",
                two: "1",
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: "7.5.2.4",
                two: "7.5.3",
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: "1.01",
                two: "1.001",
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: "1.0",
                two: "1.0.0",
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, compareVersion(p.one, p.two), "输入:%v", q)
    }
}