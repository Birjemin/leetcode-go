package reverse_words_in_a_string

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
                one: "the sky is blue",
            },
            a: result{
                one: "blue is sky the",
            },
        },
        {
            p: param{
                one: "  hello world!  ",
            },
            a: result{
                one: "world! hello",
            },
        },
        {
            p: param{
                one: "a good   example",
            },
            a: result{
                one: "example good a",
            },
        },
        {
            p: param{
                one: "   a   b  c d   e  ",
            },
            a: result{
                one: "e d c b a",
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        // ast.Equal(a.one, reverseWords(p.one), "输入:%v", q)
        // ast.Equal(a.one, reverseWords1(p.one), "输入:%v", q)
        ast.Equal(a.one, reverseWords2(p.one), "输入:%v", q)
    }
}

