package binary_tree_right_side_view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *TreeNode
}

type result struct {
	one []int
}

type question struct {
	p param
	a result
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
				one: makeTreeNode([]int{1, 2, 3, 0, 5, 0, 4}),
			},
			a: result{
				one: []int{1, 3, 4},
			},
		},
		{
			p: param{
				one: makeTreeNode([]int{}),
			},
			a: result{
				one: []int{},
			},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, rightSideView(p.one), "输入:%v", q)
		ast.Equal(a.one, rightSideView1(p.one), "输入:%v", q)
		ast.Equal(a.one, rightSideView2(p.one), "输入:%v", q)
	}
}
