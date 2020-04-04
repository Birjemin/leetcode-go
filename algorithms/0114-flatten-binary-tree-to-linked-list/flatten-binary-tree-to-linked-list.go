package flatten_binary_tree_to_linked_list

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode) {
    dfs(root)
}

func dfs(curr *TreeNode) *TreeNode {
    if curr == nil {
        return nil
    }

    l1, l2 := dfs(curr.Left), dfs(curr.Right)

    switch {
    case l1 == nil && l2 == nil:
        return curr
    case l1 != nil && l2 != nil:
        curr.Right, l1.Right = curr.Left, curr.Right
        curr.Left = nil
        return l2
    case l1 != nil && l2 == nil:
        curr.Right, curr.Left = curr.Left, nil
        return l1
    default:
        return l2
    }
}
