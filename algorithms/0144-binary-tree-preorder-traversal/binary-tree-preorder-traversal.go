package binary_tree_preorder_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    var ret []int

    for stack, length := []*TreeNode{root}, 1; length != 0; {

        // drop one element
        root, stack = stack[length-1], stack[:length-1]
        length--

        ret = append(ret, root.Val)

        if root.Right != nil {
            stack = append(stack, root.Right)
            length++
        }
        if root.Left != nil {
            stack = append(stack, root.Left)
            length++
        }
    }
    return ret
}
