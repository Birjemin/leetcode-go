package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
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
				one: []int{0, 1, 0, 3, 12},
			},
			a: result{
				one: []int{1, 3, 12, 0, 0},
			},
		},
		{
			p: param{
				one: []int{0, 0, 1},
			},
			a: result{
				one: []int{1, 0, 0},
			},
		},
		{
			p: param{
				one: []int{1, 0, 1},
			},
			a: result{
				one: []int{1, 1, 0},
			},
		},
		{
			p: param{
				one: []int{1, 0},
			},
			a: result{
				one: []int{1, 0},
			},
		},
		{
			p: param{
				one: []int{1, 2},
			},
			a: result{
				one: []int{1, 2},
			},
		},
		{
			p: param{
				one: []int{0, 1, 2, 0, 3, 0, 4, 0},
			},
			a: result{
				one: []int{1, 2, 3, 4, 0, 0, 0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		moveZeroes2(p.one)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
