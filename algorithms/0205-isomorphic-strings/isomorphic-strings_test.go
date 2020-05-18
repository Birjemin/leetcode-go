package isomorphic_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one string
	two string
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
				one: "egg",
				two: "add",
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: "foo",
				two: "bar",
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: "paper",
				two: "title",
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: "a",
				two: "a",
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: "ab",
				two: "aa",
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isIsomorphic(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, isIsomorphic1(p.one, p.two), "输入:%v", p)
	}
}
