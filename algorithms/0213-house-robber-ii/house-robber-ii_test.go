package house_robber_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
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
				one: []int{2, 3, 2},
			},
			a: result{
				one: 3,
			},
		},
		{
			p: param{
				one: []int{1, 2, 3, 1},
			},
			a: result{
				one: 4,
			},
		},
		{
			p: param{
				one: []int{1},
			},
			a: result{
				one: 1,
			},
		},
		{
			p: param{
				one: []int{1, 2},
			},
			a: result{
				one: 2,
			},
		},
		{
			p: param{
				one: []int{1, 2, 1, 1},
			},
			a: result{
				one: 3,
			},
		},
		{
			p: param{
				one: []int{2, 1, 1, 2},
			},
			a: result{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, rob(p.one), "输入:%v", p)
		ast.Equal(a.one, rob1(p.one), "输入:%v", p)
	}
}
