package binary_tree_inorder_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var ret []int
    cal(root, &ret)
    return ret
}

func cal(t *TreeNode, ret *[]int) {
    if t == nil {
        return
    }
    // left child node
    cal(t.Left, ret)
    *ret = append(*ret, t.Val)
    // right child node
    cal(t.Right, ret)
}