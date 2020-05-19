package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	cal(root)
	return root
}

func cal(t *TreeNode) {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	cal(t.Left)
	cal(t.Right)
}