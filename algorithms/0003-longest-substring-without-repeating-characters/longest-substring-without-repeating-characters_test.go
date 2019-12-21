package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one string
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
				one: "abcabcbb",
			},
			a: result{
				one: 3,
			},
		},
		{
			p: param{
				one: "abba",
			},
			a: result{
				one: 2,
			},
		},
		{
			p: param{
				one: "bbbbb",
			},
			a: result{
				one: 1,
			},
		},
		{
			p: param{
				one: "pwwkew",
			},
			a: result{
				one: 3,
			},
		},
		{
			p: param{
				one: "",
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: "abcdef",
			},
			a: result{
				one: 6,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstring(p.one), "输入:%v", p)
		ast.Equal(a.one, lengthOfLongestSubstring1(p.one), "输入:%v", p)
	}
}
