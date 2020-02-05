package string_to_integer_atoi


import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
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
               one: "+1",
           },
           a: result{
               one: 1,
           },
        },
        {
           p: param{
               one: "+-2",
           },
           a: result{
               one: 0,
           },
        },
        {
           p: param{
               one: "42",
           },
           a: result{
               one: 42,
           },
        },
        {
          p: param{
              one: "  -42",
          },
          a: result{
              one: -42,
          },
        },
        {
          p: param{
              one: "4193 with words",
          },
          a: result{
              one: 4193,
          },
        },
        {
          p: param{
              one: "words and 987",
          },
          a: result{
              one: 0,
          },
        },
        {
          p: param{
              one: "-91283472332",
          },
          a: result{
              one: -2147483648,
          },
        },
        {
           p: param{
               one: "  0000000000012345678",
           },
           a: result{
               one: 12345678,
           },
        },
        {
            p: param{
                one: "9223372036854775808",
            },
            a: result{
                one: 2147483647,
            },
        },
        {
            p: param{
                one: "-9223372036854775808",
            },
            a: result{
                one: -2147483648,
            },
        },
        {
            p: param{
                one: "-2147483647",
            },
            a: result{
                one: -2147483647,
            },
        },
        {
            p: param{
                one: "18446744073709551617",
            },
            a: result{
                one: 2147483647,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, myAtoi(p.one), "输入:%v", q)
        ast.Equal(a.one, myAtoi1(p.one), "输入:%v", q)
    }
}
