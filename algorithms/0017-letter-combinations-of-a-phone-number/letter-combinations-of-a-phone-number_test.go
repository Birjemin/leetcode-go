package letter_combinations_of_a_phone_number

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
}

type result struct {
    one []string
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
                one: "23",
            },
            a: result{
                one: []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, letterCombinations(p.one), "输入:%v", q)
    }
}
