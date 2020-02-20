package group_anagrams

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []string
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
                one: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
            },
            a: result{
                one: [][]string{
                    {"eat", "tea", "ate"},
                    {"tan", "nat"},
                    {"bat"},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, groupAnagrams(p.one), "输入:%v", q)
        ast.Equal(a.one, groupAnagrams1(p.one), "输入:%v", q)
    }
}
