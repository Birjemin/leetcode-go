package ugly_number

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
				one: 0,
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 6,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 8,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 14,
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: -8,
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isUgly(p.one), "输入:%v", p)
		ast.Equal(a.one, isUgly1(p.one), "输入:%v", p)
		ast.Equal(a.one, isUgly2(p.one), "输入:%v", p)
	}
}
