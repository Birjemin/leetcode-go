## 问题
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given
```
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
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
根据先序和中序结果创建树

- 通过先序遍历找到根结点A，再通过A在中序遍历的位置找出左子树，右子树
- 在A的左子树中，找左子树的根结点（在先序中找），转步骤1
- 在A的右子树中，找右子树的根结点（在先序中找），转步骤1

## 最高分
正常的逻辑运算
```golang
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    i := find(preorder[0], inorder)
    left := buildTree(preorder[1:i+1], inorder[:i])
    right := buildTree(preorder[i+1:], inorder[i+1:])
    return &TreeNode{
        Val: preorder[0],
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
简单的实现

```golang
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    head := &TreeNode{Val:preorder[0]}
    t := head
    cal1(t, preorder, inorder)
    return head
}

func cal1(t *TreeNode, preorder, inorder []int) {
    length := len(preorder)
    if length == 0 {
        return
    }
    for i := 0; i < length; i++ {
        if inorder[i] == t.Val {
            if i != 0 {
                t.Left = &TreeNode{Val:preorder[1]}
                cal1(t.Left, preorder[1:1+i], inorder[:i])
            }
            if i != length-1 {
                t.Right = &TreeNode{Val:preorder[1+i]}
                cal1(t.Right, preorder[1+i:], inorder[i+1:])
            }
        }
    }
}
```

## 改进
合并函数改进
```golang
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    head := &TreeNode{Val: preorder[0]}
    t := head
    cal(t, preorder, inorder)
    return head
}

func cal(t *TreeNode, preorder, inorder []int) {
    length := len(preorder)
    if length == 0 {
        return
    }
    
    i := findIdx(inorder, t.Val)

    if i != 0 {
        t.Left = &TreeNode{Val: preorder[1]}
        cal(t.Left, preorder[1:1+i], inorder[:i])
    }

    if i != length-1 {
        t.Right = &TreeNode{Val: preorder[1+i]}
        cal(t.Right, preorder[1+i:], inorder[i+1:])
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
激进的做法
```
func findIdx(inorder []int, val int) int {
    for i, v := range inorder {
        if v == val {
            return i
        }
    }
    return 0
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    i := findIdx(inorder, preorder[0])
    return &TreeNode{
        Val:   preorder[0],
        Left:  buildTree(preorder[1:1+i], inorder[:i]),
        Right: buildTree(preorder[1+i:], inorder[i+1:]),
    }
}
```

## 反思

## 参考