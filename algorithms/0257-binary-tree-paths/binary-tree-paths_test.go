package binary_tree_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *TreeNode
}

type result struct {
	one []string
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
				one: makeTreeNode([]int{1, 2, 3, 0, 5}),
			},
			a: result{
				one: []string{
					"1->2->5",
					"1->3",
				},
			},
		},
		{
			p: param{
				one: makeTreeNode([]int{1, 2, 3, 4, 5, 6}),
			},
			a: result{
				one: []string{
					"1->2->4",
					"1->2->5",
					"1->3->6",
				},
			},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, binaryTreePaths(p.one), "输入:%v", q)
	}
}
