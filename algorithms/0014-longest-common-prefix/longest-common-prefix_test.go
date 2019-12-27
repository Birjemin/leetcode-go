package longest_common_prefix

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []string
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
                one: []string{"flower","flow","flight"},
            },
            a: result{
                one: "fl",
            },
        },
        {
            p: param{
                one: []string{"dog","racecar","car"},
            },
            a: result{
                one: "",
            },
        },
        {
            p: param{
                one: []string{},
            },
            a: result{
                one: "",
            },
        },
        {
            p: param{
                one: []string{"abc"},
            },
            a: result{
                one: "abc",
            },
        },
        {
            p: param{
                one: []string{"abc", "abcd", "abcde"},
            },
            a: result{
                one: "abc",
            },
        },
        {
            p: param{
                one: []string{"", "abc", "abcde"},
            },
            a: result{
                one: "",
            },
        },
        {
            p: param{
                one: []string{"aa", "a"},
            },
            a: result{
                one: "a",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, longestCommonPrefix1(p.one), "输入:%v", q)
        ast.Equal(a.one, longestCommonPrefix2(p.one), "输入:%v", q)
    }
}
