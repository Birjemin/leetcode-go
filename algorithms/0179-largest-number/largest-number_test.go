package largest_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
}

type result struct {
	one string
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
				one: []int{20, 30, 3, 2},
			},
			a: result{
				one: "330220",
			},
		},
		{
			p: param{
				one: []int{10, 2},
			},
			a: result{
				one: "210",
			},
		},
		{
			p: param{
				one: []int{20, 1},
			},
			a: result{
				one: "201",
			},
		},
		{
			p: param{
				one: []int{3, 43, 5},
			},
			a: result{
				one: "5433",
			},
		},
		{
			p: param{
				one: []int{3, 30, 34, 5, 9},
			},
			a: result{
				one: "9534330",
			},
		},
		{
			p: param{
				one: []int{121, 12},
			},
			a: result{
				one: "12121",
			},
		},
		{
			p: param{
				one: []int{1, 2, 3, 1},
			},
			a: result{
				one: "3211",
			},
		},
		{
			p: param{
				one: []int{0, 0, 0},
			},
			a: result{
				one: "0",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, largestNumber(p.one), "输入:%v", p)
		ast.Equal(a.one, largestNumber1(p.one), "输入:%v", p)
	}
}
