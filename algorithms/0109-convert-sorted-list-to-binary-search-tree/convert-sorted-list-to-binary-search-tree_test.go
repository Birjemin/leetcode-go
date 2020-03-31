package convert_sorted_list_to_binary_search_tree

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *ListNode
    two *TreeNode
}

type result struct {
    one *TreeNode
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

func makeTreeNode(ints []int) *TreeNode {
    n := len(ints)
    if n == 0 {
        return nil
    }
    root := &TreeNode{
        Val: ints[0],
    }
    queue := make([]*TreeNode, 1, n*2)
    queue[0] = root
    i := 1
    for i < n {
        node := queue[0]
        queue = queue[1:]
        if i < n && ints[i] != 0 {
            node.Left = &TreeNode{Val: ints[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < n && ints[i] != 0 {
            node.Right = &TreeNode{Val: ints[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

func Test(t *testing.T) {
    ast := assert.New(t)
    qs := []question{
        {
            p: param{
                one: makeListNode([]int{-10, -3, 1, 5, 9}),
            },
            a: result{
                one: makeTreeNode([]int{1, -3, 9, -10, 0, 5}),
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, sortedListToBST(p.one), "输入:%v", q)
    }
}
