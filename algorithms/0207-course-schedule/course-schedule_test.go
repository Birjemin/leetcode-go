package course_schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
	two [][]int
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
				one: 2,
				two: [][]int{{1, 0}},
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 4,
				two: [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}},
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 2,
				two: [][]int{{1, 0}, {0, 1}},
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 2,
				two: [][]int{},
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 4,
				two: [][]int{{0, 1}, {3, 1}, {1, 3}, {3, 2}},
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 3,
				two: [][]int{{1, 0}, {2, 0}, {0, 2}},
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 3,
				two: [][]int{{1, 0}, {0, 2}, {2, 1}},
			},
			a: result{
				one: false,
			},
		},
		{
			p: param{
				one: 3,
				two: [][]int{{0, 1}, {0, 2}, {1, 2}},
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 5,
				two: [][]int{{1, 0}, {4, 2}, {2, 3}, {3, 1}},
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: 7,
				two: [][]int{{1, 2}, {1, 3}, {1, 6}, {2, 4}, {4, 5}, {6, 5}},
			},
			a: result{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		// ast.Equal(a.one, canFinish(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, canFinish1(p.one, p.two), "输入:%v", p)
	}
}
