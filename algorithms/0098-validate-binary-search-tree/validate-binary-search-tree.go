package validate_binary_search_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var ret []int
    return cal(root, &ret)
}

func cal(t *TreeNode, ret *[]int) bool {
    if t == nil {
        return true
    }
    // left child node
    if !cal(t.Left, ret) {
        return false
    }
    if *ret == nil {
        *ret = append(*ret, t.Val)
    } else if (*ret)[0] >= t.Val {
        return false
    } else {
        (*ret)[0] = t.Val
    }
    // right child node
    return cal(t.Right, ret)
}

func isValidBST1(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return cal1(root, nil, nil)
}

func cal1(n, min, max *TreeNode) bool {
    if n == nil {
        return true
    }
    if min != nil && min.Val >= n.Val {
        return false
    }
    if max != nil && max.Val <= n.Val {
        return false
    }
    return cal1(n.Left, min, n) && cal1(n.Right, n, max)
}
