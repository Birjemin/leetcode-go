package bulls_and_cows

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one string
	two string
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
				one: "1807",
				two: "7810",
			},
			a: result{
				one: "1A3B",
			},
		},
		{
			p: param{
				one: "1123",
				two: "0111",
			},
			a: result{
				one: "1A1B",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, getHint(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, getHint1(p.one, p.two), "输入:%v", p)
	}
}
