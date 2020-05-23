package lowest_common_ancestor_of_a_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one   *TreeNode
	two   *TreeNode
	three *TreeNode
}

type result struct {
	one *TreeNode
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

	for i := 1; i < n; {
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
				one:   makeTreeNode([]int{6, 2, 8, 1, 4, 7, 9, 0, 0, 3, 5}),
				two:   makeTreeNode([]int{2}),
				three: makeTreeNode([]int{8}),
			},
			a: result{
				one: makeTreeNode([]int{6}),
			},
		},
		{
			p: param{
				one:   makeTreeNode([]int{6, 2, 8, 1, 4, 7, 9, 0, 0, 3, 5}),
				two:   makeTreeNode([]int{2}),
				three: makeTreeNode([]int{4}),
			},
			a: result{
				one: makeTreeNode([]int{2}),
			},
		},
		{
			p: param{
				one:   makeTreeNode([]int{6, 2, 8, 1, 4, 7, 9, 0, 0, 3, 5}),
				two:   makeTreeNode([]int{1}),
				three: makeTreeNode([]int{5}),
			},
			a: result{
				one: makeTreeNode([]int{2}),
			},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one.Val, lowestCommonAncestor(p.one, p.two, p.three).Val, "输入:%v", q)
	}
}
