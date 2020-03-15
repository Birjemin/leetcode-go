package minimum_depth_of_binary_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
    var depth int
    if root == nil {
        return 0
    }
    cal(root, &depth, 1)
    return depth
}

func cal(t *TreeNode, depth *int, count int) {
    // find the end of node
    if t.Left == nil && t.Right == nil {
        // get the min depth of node
        if *depth == 0 || count < *depth {
            *depth = count
        }
        return
    }
    if t.Left != nil {
        cal(t.Left, depth, count+1)
    }
    if t.Right != nil {
        cal(t.Right, depth, count+1)
    }
}

func minDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left, right := minDepth(root.Left), minDepth(root.Right)
    if left == 0 || right == 0 {
        return left + right + 1
    }
    if left > right {
        return right + 1
    }
    return left + 1
}
