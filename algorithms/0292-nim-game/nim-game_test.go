package nim_game

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
				one: 4,
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 5,
			},
			a: result{
				one: true,
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
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, canWinNim(p.one), "输入:%v", p)
		// ast.Equal(a.one, canWinNim1(p.one), "输入:%v", p)
	}
}
