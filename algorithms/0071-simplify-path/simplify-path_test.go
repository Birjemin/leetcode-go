package simplify_path

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
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
                one: "/home/",
            },
            a: result{
                one: "/home",
            },
        },
        {
            p: param{
                one: "/../",
            },
            a: result{
                one: "/",
            },
        },
        {
            p: param{
                one: "/home//foo/",
            },
            a: result{
                one: "/home/foo",
            },
        },
        {
            p: param{
                one: "/a/./b/../../c/",
            },
            a: result{
                one: "/c",
            },
        },
        {
            p: param{
                one: "/a/../../b/../c//.//",
            },
            a: result{
                one: "/c",
            },
        },
        {
            p: param{
                one: "/a//b////c/d//././/..",
            },
            a: result{
                one: "/a/b/c",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, simplifyPath(p.one), "输入:%v", q)
    }
}
