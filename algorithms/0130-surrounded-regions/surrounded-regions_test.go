package surrounded_regions

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one [][]byte
}

type result struct {
    one [][]byte
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
                    {'X', 'X', 'X', 'X'},
                    {'X', 'O', 'O', 'X'},
                    {'X', 'X', 'O', 'X'},
                    {'X', 'O', 'X', 'X'},
                },
            },
            a: result{
                one: [][]byte{
                    {'X', 'X', 'X', 'X'},
                    {'X', 'X', 'X', 'X'},
                    {'X', 'X', 'X', 'X'},
                    {'X', 'O', 'X', 'X'},
                },
            },
        },
        {
            p: param{
                one: [][]byte{
                    {'X','O','X','O','X','O'},
                    {'O','X','O','X','O','X'},
                    {'X','O','X','O','X','O'},
                    {'O','X','O','X','O','X'},
                },
            },
            a: result{
                one: [][]byte{
                    {'X','O','X','O','X','O'},
                    {'O','X','X','X','X','X'},
                    {'X','X','X','X','X','O'},
                    {'O','X','O','X','O','X'},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        solve(p.one)
        ast.Equal(a.one, p.one, "输入:%v", q)
    }
}
