## 问题
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,
```
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
```
return true, as there exist a root-to-leaf path `5->4->11->2` which sum is 22.

## 分析
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

## 最高分
比我的简洁很多
```golang
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}
```

## 实现
就是判断到达叶子节点时（left和right为nil）是否为sum，不为sum则为false
```golang
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    // find the child node
    if root.Left == nil && root.Right == nil && sum == root.Val {
        return true
    }
    // left node
    if root.Left != nil && hasPathSum(root.Left, sum-root.Val) {
        return true
    }
    // right node
    if root.Right != nil && hasPathSum(root.Right, sum-root.Val) {
        return true
    }
    return false
}
```

## 改进
```golang
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    // find the child node
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    // left node
    return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
```

## 反思

## 参考