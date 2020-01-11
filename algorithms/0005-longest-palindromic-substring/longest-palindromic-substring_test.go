package longest_palindromic_substring

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
              one: "babad",
          },
          a: result{
              one: "bab",
          },
        },
        {
         p: param{
             one: "cbbd",
         },
         a: result{
             one: "bb",
         },
        },
        {
           p: param{
               one: "abcdeeedcba",
           },
           a: result{
               one: "abcdeeedcba",
           },
        },
        {
          p: param{
              one: "a",
          },
          a: result{
              one: "a",
          },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        //ast.Equal(a.one, longestPalindrome1(p.one), "输入:%v", q)
        ast.Equal(a.one, longestPalindrome(p.one), "输入:%v", q)
    }
}

