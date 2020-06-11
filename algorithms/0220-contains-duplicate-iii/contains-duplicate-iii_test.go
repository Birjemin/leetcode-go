package contains_duplicate_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
	two int
	three int
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
				three: 0,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: []int{1, 0, 1, 1},
				two: 1,
				three: 2,
			},
			a: result{
				one: true,
			},
		},
		{
			p: param{
				one: []int{1, 5, 9, 1, 5, 9},
				two: 2,
				three: 3,
			},
			a: result{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, containsNearbyAlmostDuplicate(p.one, p.two, p.three), "输入:%v", p)
		ast.Equal(a.one, containsNearbyAlmostDuplicate1(p.one, p.two, p.three), "输入:%v", p)
	}
}

