package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
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
				one: 123,
			},
			a: result{
				one: 321,
			},
		},
		{
			p: param{
				one: -123,
			},
			a: result{
				one: -321,
			},
		},
		{
			p: param{
				one: 120,
			},
			a: result{
				one: 21,
			},
		},
		{
			p: param{
				one: 2147483647,
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: -2147483648,
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: 0,
			},
			a: result{
				one: 0,
			},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverse1(p.one), "输入:%v", q)
		ast.Equal(a.one, reverse2(p.one), "输入:%v", q)
	}
}