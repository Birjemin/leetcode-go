package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    length := len(preorder)
    if length == 0 {
        return nil
    }
    head := &TreeNode{Val: preorder[0]}
    t := head
    cal(t, preorder, inorder, length)
    return head
}

func cal(t *TreeNode, preorder, inorder []int, length int) {
    if length == 0 {
        return
    }

    i := findIdx(inorder, t.Val)

    if i != 0 {
        t.Left = &TreeNode{Val: preorder[1]}
        cal(t.Left, preorder[1:1+i], inorder[:i], i)
    }

    if i != length-1 {
        t.Right = &TreeNode{Val: preorder[1+i]}
        cal(t.Right, preorder[1+i:], inorder[1+i:], length-1-i)
    }
}

func findIdx(inorder []int, val int) int {
    for i, v := range inorder {
        if v == val {
            return i
        }
    }
    return 0
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    i := findIdx(inorder, preorder[0])
    return &TreeNode{
        Val:   preorder[0],
        Left:  buildTree1(preorder[1:1+i], inorder[:i]),
        Right: buildTree1(preorder[1+i:], inorder[i+1:]),
    }
}
