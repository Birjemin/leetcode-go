package unique_binary_search_trees_ii

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return recGenTrees(1, n)
}

func recGenTrees(s, e int) []*TreeNode {
    if e < s {
        return []*TreeNode{nil}
    }
    var result []*TreeNode
    for i := s; i <= e; i++ {
        lefts, rights := recGenTrees(s, i-1), recGenTrees(i+1, e)
        for _, l := range lefts {
            for _, r := range rights {
                result = append(result, &TreeNode{
                    Val:   i,
                    Left:  l,
                    Right: r,
                })
            }
        }
    }
    return result
}
