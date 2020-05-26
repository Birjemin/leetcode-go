package word_pattern

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
				one: "abba",
				two: "dog cat cat dog",
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: "abba",
				two: "dog cat cat fish",
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: "aaaa",
				two: "dog cat cat dog",
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: "abba",
				two: "dog dog dog dog",
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: "aba",
				two: "cat cat cat dog",
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, wordPattern(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, wordPattern1(p.one, p.two), "输入:%v", p)
	}
}
