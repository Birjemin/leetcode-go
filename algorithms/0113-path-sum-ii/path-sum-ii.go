package path_sum_ii

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
    var ret [][]int
    var tmp []int
    cal(root, sum, tmp, &ret)
    return ret
}

func cal(root *TreeNode, sum int, tmp []int, ret *[][]int) {
    if root == nil {
        return
    }
    sum -= root.Val
    tmp = append(tmp, root.Val)
    // find the child node
    if root.Left == nil && root.Right == nil {
        if sum == 0 {
            b := make([]int, len(tmp))
            copy(b, tmp)
            *ret = append(*ret, b)
        }
        return
    }
    // left node
    cal(root.Left, sum, tmp, ret)
    // right node
    cal(root.Right, sum, tmp, ret)
}
