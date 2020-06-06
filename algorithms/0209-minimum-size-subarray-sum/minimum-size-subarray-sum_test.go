package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
	two []int
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
				one: 7,
				two: []int{2, 3, 1, 2, 4, 3},
			},
			a: result{
				one: 2,
			},
		},
		{
			p: param{
				one: 213,
				two: []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			},
			a: result{
				one: 8,
			},
		},
		{
			p: param{
				one: 11,
				two: []int{1, 2, 3, 4, 5},
			},
			a: result{
				one: 3,
			},
		},
		{
			p: param{
				one: 3,
				two: []int{1, 1},
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: 3,
				two: []int{1, 2, 3},
			},
			a: result{
				one: 1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, minSubArrayLen(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, minSubArrayLen1(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, minSubArrayLen2(p.one, p.two), "输入:%v", p)
	}
}
