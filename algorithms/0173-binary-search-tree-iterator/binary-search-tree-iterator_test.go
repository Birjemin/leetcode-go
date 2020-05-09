package binary_search_tree_iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	root := makeTreeNode([]int{7, 3, 15, 0, 0, 9, 20})
	iterator := Constructor(root)

	ast.Equal(iterator.Next(), 3, "输入:%v")
	ast.Equal(iterator.Next(), 7, "输入:%v")
	ast.Equal(iterator.HasNext(), true, "输入:%v")
	ast.Equal(iterator.Next(), 9, "输入:%v")
	ast.Equal(iterator.HasNext(), true, "输入:%v")
	ast.Equal(iterator.Next(), 15, "输入:%v")
	ast.Equal(iterator.HasNext(), true, "输入:%v")
	ast.Equal(iterator.Next(), 20, "输入:%v")
	ast.Equal(iterator.HasNext(), false, "输入:%v")

}
