package add_digits

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
				one: 38,
			},
			a: result{
				one: 2,
			},
		},
		{
			p: param{
				one: 90,
			},
			a: result{
				one: 9,
			},
		},{
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
		ast.Equal(a.one, addDigits(p.one), "输入:%v", p)
		ast.Equal(a.one, addDigits1(p.one), "输入:%v", p)
	}
}
