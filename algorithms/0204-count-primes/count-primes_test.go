package count_primes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
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
				one: 10,
			},
			a: result{
				one: 4,
			},
		},
		{
			p: param{
				one: 3,
			},
			a: result{
				one: 1,
			},
		},
		{
			p: param{
				one: 2,
			},
			a: result{
				one: 0,
			},
		},
		{
			p: param{
				one: 1,
			},
			a: result{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, countPrimes(p.one), "输入:%v", p)
		ast.Equal(a.one, countPrimes1(p.one), "输入:%v", p)
	}
}
