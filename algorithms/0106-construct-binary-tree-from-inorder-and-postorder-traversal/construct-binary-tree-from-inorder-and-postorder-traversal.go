package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    length := len(postorder)
    if length == 0 {
        return nil
    }
    head := &TreeNode{Val: postorder[length-1]}
    t := head
    cal(t, inorder, postorder, length)
    return head
}

func cal(t *TreeNode, inorder, postorder []int, length int) {
    if length == 0 {
        return
    }
    i := findIdx(inorder, t.Val)

    if i != 0 {
        t.Left = &TreeNode{Val: postorder[i-1]}
        cal(t.Left, inorder[:i], postorder[:i], i)
    }
    if i != length-1 {
        t.Right = &TreeNode{Val: postorder[length-2]}
        cal(t.Right, inorder[1+i:], postorder[i:length-1], length-1-i)
    }
}

func findIdx(inorder []int, val int) int {
    for i, v := range inorder {
        if v == val {
            return i
        }
    }
    return -1
}

func buildTree1(inorder []int, postorder []int) *TreeNode {
    length := len(postorder)
    if length == 0 {
        return nil
    }
    i := findIdx(inorder, postorder[length-1])
    return &TreeNode{
        Val:   postorder[length-1],
        Left:  buildTree1(inorder[:i], postorder[:i]),
        Right: buildTree1(inorder[i+1:], postorder[i:length-1]),
    }
}
