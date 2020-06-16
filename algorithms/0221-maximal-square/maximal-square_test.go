package maximal_square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one [][]byte
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
				one: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			a: result{
				one: 4,
			},
		},
		{
			p: param{
				one: [][]byte{},
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: [][]byte{
					{'0', '1'},
				},
			},
			a: result{
				one: 1,
			},
		},
		{
			p: param{
				one: [][]byte{
					{'0', '0', '0'},
					{'0', '0', '0'},
					{'1', '1', '1'},
				},
			},
			a: result{
				one: 1,
			},
		},
		{
			p: param{
				one: [][]byte{
					{'0', '0', '1', '0'},
					{'0', '1', '1', '1'},
					{'1', '1', '1', '1'},
					{'1', '1', '1', '1'},
				},
			},
			a: result{
				one: 9,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, maximalSquare(p.one), "输入:%v", p)
		ast.Equal(a.one, maximalSquare1(p.one), "输入:%v", p)
	}
}
