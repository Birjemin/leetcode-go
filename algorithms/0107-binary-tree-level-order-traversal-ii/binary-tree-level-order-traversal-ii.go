package binary_tree_level_order_traversal_ii

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
    tmp := make(map[int][]int)
    dfs(root, &tmp, 0)
    return convertArrInt(tmp)
}

// convert map[int][]int to [][]int
func convertArrInt(tmp map[int][]int) [][]int {
    length := len(tmp)
    ret := make([][]int, length)
    for i, v := range tmp {
        ret[length-1-i] = v
    }
    return ret
}

// dsf
func dfs(t *TreeNode, tmp *map[int][]int, depth int) {
    if t == nil {
        return
    }
    // record the value
    (*tmp)[depth] = append((*tmp)[depth], t.Val)
    dfs(t.Left, tmp, depth+1)
    dfs(t.Right, tmp, depth+1)
}

func dfs1(s *[][]int, level int, root *TreeNode)  {
    if root == nil {
        return
    }
    if len(*s) == level {
        *s = append(*s, []int{})
    }
    (*s)[level] = append((*s)[level], root.Val)
    for _, v := range []*TreeNode{root.Left, root.Right} {
        dfs1(s, level + 1, v)
    }
}

func levelOrderBottom1(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    var s [][]int
    dfs1(&s, 0, root)
    for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return s
}
