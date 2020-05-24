package missing_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
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
				one: []int{3, 0, 1},
			},
			a: result{
				one: 2,
			},
		},
		{
			p: param{
				one: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			a: result{
				one: 8,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, missingNumber(p.one), "输入:%v", p)
		ast.Equal(a.one, missingNumber1(p.one), "输入:%v", p)
		ast.Equal(a.one, missingNumber2(p.one), "输入:%v", p)
	}
}
