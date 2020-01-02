package count_and_say

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
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
               one: 1,
           },
           a: result{
               one: "1",
           },
        },
        {
          p: param{
              one: 2,
          },
          a: result{
              one: "11",
          },
        },
        {
          p: param{
              one: 4,
          },
          a: result{
              one: "1211",
          },
        },
        {
           p: param{
               one: 7,
           },
           a: result{
               one: "13112221",
           },
        },
        {
           p: param{
               one: 10,
           },
           a: result{
               one: "13211311123113112211",
           },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, countAndSay(p.one), "输入:%v", q)
    }
}