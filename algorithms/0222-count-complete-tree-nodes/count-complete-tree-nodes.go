package count_complete_tree_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	traverse(root,&ret)
	return ret
}

func traverse(root *TreeNode, ret *int) {
	if root == nil {
		return
	}
	*ret++
	traverse(root.Left, ret)
	traverse(root.Right, ret)
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes1(root.Left) + countNodes1(root.Right)
}
