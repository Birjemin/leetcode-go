package symmetric_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
    // if root is nil
    if root == nil {
        return true
    }
    return traversal(root.Left, root.Right)
}

func traversal(l, r *TreeNode) bool {
    // end of traversal
    if l == nil && r == nil {
        return true
    }
    // the conditions of error
    if l == nil || r == nil || l.Val != r.Val || !traversal(l.Left, r.Right) {
        return false
    }
    return traversal(l.Right, r.Left)
}
