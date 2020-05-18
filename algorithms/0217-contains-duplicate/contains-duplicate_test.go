package contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
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
				one: []int{1, 2, 3, 1},
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: []int{1, 2, 3, 4},
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: []int{1, 2, 3, 11, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			a: result{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, containsDuplicate(p.one), "输入:%v", p)
		ast.Equal(a.one, containsDuplicate1(p.one), "输入:%v", p)
	}
}
