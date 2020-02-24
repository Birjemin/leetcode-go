package unique_paths

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
    two int
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
                one: 3,
                two: 2,
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: 7,
                two: 3,
            },
            a: result{
                one: 28,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, uniquePaths(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, uniquePaths1(p.one, p.two), "输入:%v", q)
    }
}
