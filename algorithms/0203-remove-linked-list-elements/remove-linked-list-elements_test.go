package remove_linked_list_elements

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *ListNode
	two int
}

type result struct {
	one *ListNode
}

type question struct {
	p param
	a result
}

func makeListNode(params []int) *ListNode {
	if len(params) == 0 {
		return nil
	}
	res := &ListNode{
		Val: params[0],
	}
	temp := res
	for i, v := range params {
		if i != 0 {
			temp.Next = &ListNode{Val: v}
			temp = temp.Next
		}
	}
	return res
}

func showListNode(res *ListNode) {
	str := ""
	for res != nil {
		str += cast.ToString(res.Val) + " -> "
		res = res.Next
	}
	str += "\n"
	fmt.Print(str)
}

func Test(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: makeListNode([]int{1, 2, 6, 3, 4, 5, 6}),
				two: 6,
			},
			a: result{
				one: makeListNode([]int{1, 2, 3, 4, 5}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{1}),
				two: 1,
			},
			a: result{
				one: makeListNode([]int{}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{1, 2, 3}),
				two: 3,
			},
			a: result{
				one: makeListNode([]int{1, 2}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{1, 2, 3}),
				two: 1,
			},
			a: result{
				one: makeListNode([]int{2, 3}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{1, 1, 2}),
				two: 2,
			},
			a: result{
				one: makeListNode([]int{1, 1}),
			},
		},

		{
			p: param{
				one: makeListNode([]int{1, 2, 2, 1}),
				two: 1,
			},
			a: result{
				one: makeListNode([]int{2, 2}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{1, 1, 1}),
				two: 1,
			},
			a: result{
				one: makeListNode([]int{}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{1, 2, 2, 1}),
				two: 2,
			},
			a: result{
				one: makeListNode([]int{1, 1}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, removeElements(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, removeElements1(p.one, p.two), "输入:%v", p)
	}
}
