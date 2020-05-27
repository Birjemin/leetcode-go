package bitwise_and_of_numbers_range

import (
	"github.com/stretchr/testify/assert"
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
				one: 5,
				two: 7,
			},
			a: result{
				one: 4,
			},
		},
		{
			p: param{
				one: 0,
				two: 1,
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: 28,
				two: 31,
			},
			a: result{
				one: 28,
			},
		},
		{
			p: param{
				one: 1,
				two: 1,
			},
			a: result{
				one: 1,
			},
		},
		{
			p: param{
				one: 5,
				two: 5,
			},
			a: result{
				one: 5,
			},
		},
		{
			p: param{
				one: 20,
				two: 22,
			},
			a: result{
				one: 20,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, rangeBitwiseAnd(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, rangeBitwiseAnd1(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, rangeBitwiseAnd2(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, rangeBitwiseAnd3(p.one, p.two), "输入:%v", p)
	}
}
