package first_bad_version


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
	two int
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
				one: 5,
				two: 4,
			},
			a: result{
				one: 4,
			},
		},
	}


	for _, q := range qs {
		a, p := q.a, q.p
		setBadVersion(p.two)
		ast.Equal(a.one, firstBadVersion(p.one), "输入:%v", p)
		ast.Equal(a.one, firstBadVersion1(p.one), "输入:%v", p)
	}
}
