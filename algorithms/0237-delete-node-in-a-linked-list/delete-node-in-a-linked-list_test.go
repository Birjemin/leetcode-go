package delete_node_in_a_linked_list

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeListNode(params []int, target int) (*ListNode, *ListNode) {
	res := &ListNode{
		Val: params[0],
	}
	var flag *ListNode
	temp := res
	for i, v := range params {
		if i != 0 {
			temp.Next = &ListNode{Val: v}
			temp = temp.Next
			if v == target {
				flag = temp
			}
		}
	}
	return res, flag
}

func showListNode(res *ListNode) string {
	str := ""
	for res != nil {
		str += cast.ToString(res.Val) + "-"
		res = res.Next
	}
	return str
}

func Test(t *testing.T) {
	ast := assert.New(t)

		list, target := makeListNode([]int{4,5,1,9}, 5)
		deleteNode1(target)
		ast.Equal("4-1-9-", showListNode(list), "输入:%v")
}
