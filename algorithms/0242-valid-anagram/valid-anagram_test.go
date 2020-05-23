package valid_anagram

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
				one: "anagram",
				two: "nagaram",
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: "rat",
				two: "car",
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: "ab",
				two: "b",
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isAnagram(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, isAnagram1(p.one, p.two), "输入:%v", p)
	}
}

