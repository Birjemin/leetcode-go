package palindrome_partitioning

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
}

type result struct {
    one [][]string
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
                one: "aab",
            },
            a: result{
                one: [][]string{
                    {"a", "a", "b"},
                    {"aa", "b"},
                },
            },
        },
        {
            p: param{
                one: "aa",
            },
            a: result{
                one: [][]string{
                    {"a", "a"},
                    {"aa"},
                },
            },
        },
        {
            p: param{
                one: "aabaa",
            },
            a: result{
                one: [][]string{
                    {"a", "a", "b", "a", "a"},
                    {"a", "a", "b", "aa"},
                    {"a", "aba", "a"},
                    {"aa", "b", "a", "a"},
                    {"aa", "b", "aa"},
                    {"aabaa"},
                },
            },
        },
        {
            p: param{
                one: "cbbbcc",
            },
            a: result{
                one: [][]string{
                    {"c", "b", "b", "b", "c", "c"},
                    {"c", "b", "b", "b", "cc"},
                    {"c", "b", "bb", "c", "c"},
                    {"c", "b", "bb", "cc"},
                    {"c", "bb", "b", "c", "c"},
                    {"c", "bb", "b", "cc"},
                    {"c", "bbb", "c", "c"},
                    {"c", "bbb", "cc"},
                    {"cbbbc", "c"},
                },
            },
        },
        {
            p: param{
                one: "cbbbcc",
            },
            a: result{
                one: [][]string{
                    {"c", "b", "b", "b", "c", "c"},
                    {"c", "b", "b", "b", "cc"},
                    {"c", "b", "bb", "c", "c"},
                    {"c", "b", "bb", "cc"},
                    {"c", "bb", "b", "c", "c"},
                    {"c", "bb", "b", "cc"},
                    {"c", "bbb", "c", "c"},
                    {"c", "bbb", "cc"},
                    {"cbbbc", "c"},
                },
            },
        },
        {
            p: param{
                one: "ababba",
            },
            a: result{
                one: [][]string{
                    {"a", "b", "a", "b", "b", "a"},
                    {"a", "b", "a", "bb", "a"},
                    {"a", "b", "abba"},
                    {"a", "bab", "b", "a"},
                    {"aba", "b", "b", "a"},
                    {"aba", "bb", "a"},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, partition(p.one), "输入:%v", q)
    }
}
