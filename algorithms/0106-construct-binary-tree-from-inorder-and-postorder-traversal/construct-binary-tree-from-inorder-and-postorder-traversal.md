## 问题
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given
```
inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
```

Return the following binary tree:
```
    3
   / \
  9  20
    /  \
   15   7
```

## 分析
根据一棵树的中序遍历与后序遍历构造二叉树。

- 通过后序遍历找到根结点A，再通过A在中序遍历的位置找出左子树，右子树
- 在A的左子树中，找左子树的根结点（在中序中找），转步骤1
- 在A的右子树中，找右子树的根结点（在中序中找），转步骤1

## 最高分
```golang
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    j := len(postorder)-1
    i := find(postorder[j], inorder)
    left := buildTree(inorder[:i], postorder[:i])
    right := buildTree(inorder[i+1:], postorder[i:j])
    return &TreeNode{
        Val: postorder[j],
        Left: left,
        Right: right,
    }
}

func find(target int, nums []int) int {
    for i, x := range nums {
        if x == target {
            return i
        }
    }
    return -1
}
```

## 实现
参照105的实现
```golang
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

```

## 改进
```golang
func buildTree(inorder []int, postorder []int) *TreeNode {
    length := len(postorder)
    if length == 0 {
        return nil
    }
    i := findIdx(inorder, postorder[length-1])
    return &TreeNode{
        Val:   postorder[length-1],
        Left:  buildTree(inorder[:i], postorder[:i]),
        Right: buildTree(inorder[i+1:], postorder[i:length-1]),
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
```

## 反思

## 参考