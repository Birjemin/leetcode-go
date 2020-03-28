package binary_tree_level_order_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    var ret [][]int
    if root == nil {
        return [][]int{}
    }
    cal([]*TreeNode{root}, &ret)
    return ret
}

func cal(tmp []*TreeNode, ret *[][]int) {
    if len(tmp) == 0 {
        return
    }
    var t []*TreeNode
    var r []int
    var flag bool
    for _, v := range tmp {
        r = append(r, v.Val)
        if v.Left != nil {
            t = append(t, v.Left)
            flag = true
        }
        if v.Right != nil {
            t = append(t, v.Right)
            flag = true
        }
    }
    *ret = append(*ret, r)
    if !flag {
        return
    }
    cal(t, ret)
}

func levelOrder1(root *TreeNode) [][]int {
    var ret [][]int
    if root == nil {
        return [][]int{}
    }
    cal1(root, 0, &ret)
    return ret
}

func cal1(tmp *TreeNode, depth int, ret *[][]int) {
    if tmp == nil {
        return
    }
    if len(*ret) <= depth {
        *ret = append(*ret, []int{})
    }

    (*ret)[depth] = append((*ret)[depth], tmp.Val)
    if tmp.Left != nil {
        cal1(tmp.Left, depth+1, ret)
    }
    if tmp.Right != nil {
        cal1(tmp.Right, depth+1, ret)
    }
}
