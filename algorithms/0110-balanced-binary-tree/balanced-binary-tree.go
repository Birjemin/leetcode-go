package balanced_binary_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    ret, _ := depth(root, 0)
    return ret
}

func depth(t *TreeNode, count int) (bool, int) {
    if t == nil {
        return true, count
    }
    flag, left := depth(t.Left, count+1)
    // find it!, break
    if !flag {
        return false, 0
    }
    flag, right := depth(t.Right, count+1)
    // find it!, break
    if !flag || left-right > 1 || right-left > 1 {
        return false, 0
    }
    // get the max value
    if left < right {
        return true, right
    } else {
        return true, left
    }
}
