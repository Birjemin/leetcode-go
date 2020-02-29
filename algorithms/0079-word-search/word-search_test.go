package word_search

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one [][]byte
    two string
}

type result struct {
    one bool
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
                one: [][]byte{
                    {'A', 'B', 'C', 'E'},
                    {'S', 'F', 'C', 'S'},
                    {'A', 'D', 'E', 'E'},
                },
                two: "ABCCED",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'A', 'B', 'C', 'E'},
                    {'S', 'F', 'C', 'S'},
                    {'A', 'D', 'E', 'E'},
                },
                two: "SEE",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'A', 'B', 'C', 'E'},
                    {'S', 'F', 'C', 'S'},
                    {'A', 'D', 'E', 'E'},
                },
                two: "ABCB",
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'A', 'B', 'C', 'E'},
                    {'S', 'F', 'C', 'S'},
                    {'A', 'D', 'E', 'E'},
                },
                two: "ABCCFB",
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'A', 'B'},
                    {'C', 'D'},
                },
                two: "DBAC",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'a', 'a', 'a', 'a'},
                    {'a', 'a', 'a', 'a'},
                    {'a', 'a', 'a', 'a'},
                },
                two: "aaaaaaaaaaaaa",
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'C', 'A', 'A'},
                    {'A', 'A', 'A'},
                    {'B', 'C', 'D'},
                },
                two: "AAB",
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'A', 'B', 'C', 'E'},
                    {'S', 'F', 'E', 'S'},
                    {'A', 'D', 'E', 'E'},
                },
                two: "ABCESEEEFS",
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, exist(p.one, p.two), "输入:%v", q)
    }
}
