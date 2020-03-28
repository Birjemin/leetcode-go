package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    var ret [][]int
    var b bool
    if root == nil {
        return [][]int{}
    }
    cal([]*TreeNode{root}, &ret, b)
    return ret
}

func cal(tmp []*TreeNode, ret *[][]int, b bool) {
    if len(tmp) == 0 {
        return
    }
    var t []*TreeNode
    var r []int
    var flag bool
    for _, v := range tmp {
        if b {
            r = append([]int{v.Val}, r...)
        } else {
            r = append(r, v.Val)
        }
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
    cal(t, ret, !b)
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
    var ret [][]int
    if root == nil {
        return [][]int{}
    }
    var b bool
    cal1(root, 0, b, &ret)
    return ret
}

func cal1(tmp *TreeNode, depth int, b bool, ret *[][]int) {
    if tmp == nil {
        return
    }
    if len(*ret) <= depth {
        *ret = append(*ret, []int{})
    }

    if b {
        (*ret)[depth] = append([]int{tmp.Val}, (*ret)[depth]...)
    } else {
        (*ret)[depth] = append((*ret)[depth], tmp.Val)
    }

    if tmp.Left != nil {
        cal1(tmp.Left, depth+1, !b, ret)
    }
    if tmp.Right != nil {
        cal1(tmp.Right, depth+1, !b, ret)
    }
}
