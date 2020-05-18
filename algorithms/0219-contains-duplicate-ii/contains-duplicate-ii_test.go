package contains_duplicate_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
	two int
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
				one: []int{1, 2, 3, 1},
				two: 3,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: []int{1, 0, 1, 1},
				two: 1,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: []int{1, 2, 3, 1, 2, 3},
				two: 2,
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, containsNearbyDuplicate(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, containsNearbyDuplicate1(p.one, p.two), "输入:%v", p)
	}
}
