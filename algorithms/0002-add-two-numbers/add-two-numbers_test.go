package add_two_numbers

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *ListNode
	two *ListNode
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
		} else {
			continue
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
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: result{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{2, 7, 7, 8, 5}),
				two: makeListNode([]int{8, 2, 2, 1, 4}),
			},
			a: result{
				one: makeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		{
			p: param{
				one: makeListNode([]int{0}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: result{
				one: makeListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers1(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, addTwoNumbers2(p.one, p.two), "输入:%v", p)
	}
}
