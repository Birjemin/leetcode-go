package course_schedule_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
	two [][]int
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
				one: 1,
				two: [][]int{},
			},
			a: result{
				one: []int{0},
			},
		},
		{
			p: param{
				one: 3,
				two: [][]int{{2, 1}, {1, 0}},
			},
			a: result{
				one: []int{0, 1, 2},
			},
		},
		{
			p: param{
				one: 2,
				two: [][]int{{1, 0}},
			},
			a: result{
				one: []int{0, 1},
			},
		},
		{
			p: param{
				one: 4,
				two: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			a: result{
				one: []int{0, 1, 2, 3},
			},
		},
		{
			p: param{
				one: 3,
				two: [][]int{{1, 0}, {1, 2}, {0, 1}},
			},
			a: result{
				one: []int{},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findOrder(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, findOrder1(p.one, p.two), "输入:%v", p)
	}
}
