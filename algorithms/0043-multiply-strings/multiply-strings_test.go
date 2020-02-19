package multiply_strings

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
    two string
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
                one: "2",
                two: "3",
            },
            a: result{
                one: "6",
            },
        },
        {
            p: param{
                one: "123",
                two: "456",
            },
            a: result{
                one: "56088",
            },
        },
        {
            p: param{
                one: "9",
                two: "99",
            },
            a: result{
                one: "891",
            },
        },
        {
            p: param{
                one: "498828660196",
                two: "840477629533",
            },
            a: result{
                one: "419254329864656431168468",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, multiply(p.one, p.two), "输入:%v", q)
    }
}
