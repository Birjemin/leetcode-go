package maximum_depth_of_binary_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    var depth int
    dfs(root, &depth, 0)
    return depth
}

func dfs(t *TreeNode, depth *int, count int) {
    if t == nil {
        if count > *depth {
            *depth = count
        }
        return
    }
    dfs(t.Left, depth, count+1)
    dfs(t.Right, depth, count+1)
}
