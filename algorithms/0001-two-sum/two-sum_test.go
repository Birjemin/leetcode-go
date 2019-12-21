package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
	two int
}

type result struct {
	one []int
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
				one: []int{3, 2, 4},
				two: 6,
			},
			a: result{
				one: []int{1, 2},
			},
		},
		{
			p: param{
				one: []int{3, 2, 4},
				two: 9,
			},
			a: result{
				one: nil,
			},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum1(p.one, p.two), "输入:%v", q)
		ast.Equal(a.one, twoSum2(p.one, p.two), "输入:%v", q)
		ast.Equal(a.one, twoSum3(p.one, p.two), "输入:%v", q)
		ast.Equal(a.one, twoSum4(p.one, p.two), "输入:%v", q)
	}
}
