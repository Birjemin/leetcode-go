## 问题
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:
```
a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
```
 

Example 1:

Given the following tree `[3,9,20,null,null,15,7]`:
```
    3
   / \
  9  20
    /  \
   15   7
```
Return true.

Example 2:

Given the following tree `[1,2,2,3,3,null,null,4,4]`:
```
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```
Return false.

## 分析
给定一个二叉树，判断它是否是高度平衡的二叉树。一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

## 最高分
这个思路试了半天，没有试出来，只想到了两个参数的方式
```golang
func isBalanced(root *TreeNode) bool {
    return getDepth(root) >= 0
}

func getDepth(n *TreeNode) int {
    if n == nil {
        return 0
    }
    leftDepth, rightDepth := getDepth(n.Left), getDepth(n.Right)
    if leftDepth < 0 || rightDepth < 0 || leftDepth-rightDepth < -1 || leftDepth-rightDepth > 1 {
        return -1
    }
    if leftDepth > rightDepth {
        return leftDepth + 1
    }
    return rightDepth + 1
}
```

## 实现
有两个变量，想不到实现的好办法，返回两个值
```golang
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    ret, _ := depth(root, 0)
    return ret
}

func depth(t *TreeNode, count int) (bool, int) {
    if t == nil {
        return true, count
    }
    flag, left := depth(t.Left, count+1)
    // find it!, break
    if !flag {
        return false, 0
    }
    flag, right := depth(t.Right, count+1)
    // find it!, break
    if !flag || left-right > 1 || right-left > 1 {
        return false, 0
    }
    // get the max value
    if left < right {
        return true, right
    } else {
        return true, left
    }
}
```

## 改进
```golang

```

## 反思

## 参考