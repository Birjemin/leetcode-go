package word_ladder

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one   string
    two   string
    three []string
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
                one:   "hit",
                two:   "cog",
                three: []string{"hot", "dot", "dog", "lot", "log", "cog"},
            },
            a: result{
                one: 5,
            },
        },
        {
            p: param{
                one:   "hit",
                two:   "cog",
                three: []string{"hot", "dot", "dog", "lot", "log"},
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one:   "hot",
                two:   "dog",
                three: []string{"hot", "dog", "dot"},
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one:   "a",
                two:   "c",
                three: []string{"a", "b", "c", "d", "e"},
            },
            a: result{
                one: 2,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, ladderLength(p.one, p.two, p.three), "输入:%v", q)
    }
}
