package happy_number

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
				one: 19,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 20,
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 2,
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 10,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 11,
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isHappy(p.one), "输入:%v", p)
		ast.Equal(a.one, isHappy1(p.one), "输入:%v", p)
	}
}
