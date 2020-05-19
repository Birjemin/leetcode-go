package power_of_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
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
				one: 1,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 16,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 218,
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isPowerOfTwo(p.one), "输入:%v", p)
		ast.Equal(a.one, isPowerOfTwo1(p.one), "输入:%v", p)
	}
}
