package add_binary

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
    two string
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
               one: "11",
               two: "1",
           },
           a: result{
               one: "100",
           },
        },
        {
          p: param{
              one: "1",
              two: "0",
          },
          a: result{
              one: "1",
          },
        },
        {
          p: param{
              one: "",
              two: "",
          },
          a: result{
              one: "0",
          },
        },
        {
          p: param{
              one: "1",
              two: "",
          },
          a: result{
              one: "1",
          },
        },
        {
           p: param{
               one: "1010",
               two: "1011",
           },
           a: result{
               one: "10101",
           },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, addBinary(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, addBinary1(p.one, p.two), "输入:%v", q)
    }
}