package binary_tree_paths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var ret []string
	dfs(root, "", &ret)
	return ret
}

func dfs(t *TreeNode, str string, ret *[]string) {
	if t == nil {
		return
	}
	str += strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		*ret = append(*ret, str)
		return
	}
	str += "->"
	dfs(t.Left, str, ret)
	dfs(t.Right, str, ret)
}
