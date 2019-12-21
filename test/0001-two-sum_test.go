package test

import (
	"github.com/birjemin/leetcode-go/code"
	"github.com/stretchr/testify/assert"
	"testing"
)

type params struct {
	one []int
	two int
}

type question struct {
	p params
	a []int
}

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: params{
				one: []int{3, 2, 4},
				two: 6,
			},
			a: []int{1, 2},
		},
		{
			p: params{
				one: []int{3, 2, 4},
				two: 8,
			},
			a: nil,
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a, code.TwoSum1(p.one, p.two), "输入:%v", q)
		ast.Equal(a, code.TwoSum2(p.one, p.two), "输入:%v", q)
		ast.Equal(a, code.TwoSum3(p.one, p.two), "输入:%v", q)
		ast.Equal(a, code.TwoSum4(p.one, p.two), "输入:%v", q)
	}
}
