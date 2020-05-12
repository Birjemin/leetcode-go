package repeated_dna_sequences

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type param struct {
	one string
}

type result struct {
	one []string
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
				one: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			},
			a: result{
				one: []string{"AAAAACCCCC", "CCCCCAAAAA"},
			},
		},
		{
			p: param{
				one: "AAAAAAAAAAA",
			},
			a: result{
				one: []string{"AAAAAAAAAA"},
			},
		},
		{
			p: param{
				one: "AAAAAAAAAAAA",
			},
			a: result{
				one: []string{"AAAAAAAAAA"},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ret := findRepeatedDnaSequences(p.one)
		sort.Strings(ret)
		ast.Equal(a.one, ret, "输入:%v", p)
	}
}
