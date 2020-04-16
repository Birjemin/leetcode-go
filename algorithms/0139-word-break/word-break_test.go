package word_break

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
    two []string
}

type result struct {
    one bool
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
                one: "leetcode",
                two: []string{"leet", "code"},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "applepenapple",
                two: []string{"apple", "pen"},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "catsandog",
                two: []string{"cats", "dog", "sand", "and", "cat"},
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: "cars",
                two: []string{"car", "ca", "rs"},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "aaaaaaaaaaaaaaaaaaab",
                two: []string{"a", "aa", "aaa", "aaab"},
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: "bccdbacd",
                two: []string{ "a", "b", "c", "d", "bc", "dba"},
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, wordBreak(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, wordBreak1(p.one, p.two), "输入:%v", q)
    }
}
