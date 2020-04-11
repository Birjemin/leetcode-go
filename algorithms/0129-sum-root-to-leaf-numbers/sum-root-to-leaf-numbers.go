package sum_root_to_leaf_numbers

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
    var ret int
    dfs(root, 0, &ret)
    return ret
}

func dfs(root *TreeNode, prev int, ret *int) {
    if root == nil {
        return
    }

    prev = prev*10 + root.Val

    if root.Left == nil && root.Right == nil {
        *ret += prev
        return
    }

    dfs(root.Left, prev, ret)
    dfs(root.Right, prev, ret)
}
