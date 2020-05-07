package fraction_to_recurring_decimal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
	two int
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
				one: 1,
				two: 2,
			},
			a: result{
				one: "0.5",
			},
		},
		{
			p: param{
				one: 2,
				two: 1,
			},
			a: result{
				one: "2",
			},
		},
		{
			p: param{
				one: 1,
				two: 8,
			},
			a: result{
				one: "0.125",
			},
		},
		{
			p: param{
				one: 1,
				two: 8,
			},
			a: result{
				one: "0.125",
			},
		},
		{
			p: param{
				one: 123,
				two: 990,
			},
			a: result{
				one: "0.1(24)",
			},
		},
		{
			p: param{
				one: 12,
				two: 99,
			},
			a: result{
				one: "0.(12)",
			},
		},
		{
			p: param{
				one: 2,
				two: 3,
			},
			a: result{
				one: "0.(6)",
			},
		},
		{
			p: param{
				one: 4,
				two: 333,
			},
			a: result{
				one: "0.(012)",
			},
		},
		{
			p: param{
				one: 22,
				two: 7,
			},
			a: result{
				one: "3.(142857)",
			},
		},
		{
			p: param{
				one: 500,
				two: 10,
			},
			a: result{
				one: "50",
			},
		},
		{
			p: param{
				one: -50,
				two: 8,
			},
			a: result{
				one: "-6.25",
			},
		},
		{
			p: param{
				one: -22,
				two: -7,
			},
			a: result{
				one: "3.(142857)",
			},
		},
		{
			p: param{
				one: 0,
				two: -5,
			},
			a: result{
				one: "0",
			},
		},}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, fractionToDecimal(p.one, p.two), "输入:%v", q)
	}
}
