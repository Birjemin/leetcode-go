package divide_two_integers

import (
    "github.com/stretchr/testify/assert"
    "math"
    "testing"
)

type param struct {
    one int
    two int
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
                one: 10,
                two: 3,
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: 7,
                two: -3,
            },
            a: result{
                one: -2,
            },
        },
        {
            p: param{
                one: 2,
                two: 3,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: -2,
                two: 3,
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: 2,
                two: 2,
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: 2147483647,
                two: 2,
            },
            a: result{
                one: 1073741823,
            },
        },
        {
            p: param{
                one: 2147483648,
                two: 1,
            },
            a: result{
                one: math.MaxInt32,
            },
        },
        {
            p: param{
                one: -2147483648,
                two: -1,
            },
            a: result{
                one: math.MaxInt32,
            },
        },
        {
            p: param{
                one: -1,
                two: 1,
            },
            a: result{
                one: -1,
            },
        },
        {
            p: param{
                one: -2147483648,
                two: 2,
            },
            a: result{
                one: -1073741824,
            },
        },
        {
            p: param{
                one: 2147483647,
                two: 3,
            },
            a: result{
                one: 715827882,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, divide(p.one, p.two), "输入:%v", q)
    }
}
