package path_sum

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    // find the child node
    if root.Left == nil && root.Right == nil && sum == root.Val {
        return true
    }
    // left node
    if root.Left != nil && hasPathSum(root.Left, sum-root.Val) {
        return true
    }
    // right node
    if root.Right != nil && hasPathSum(root.Right, sum-root.Val) {
        return true
    }
    return false
}

func hasPathSum1(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    // find the child node
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    // left node
    return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
