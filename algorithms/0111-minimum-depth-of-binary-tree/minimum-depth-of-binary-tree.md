## 问题
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree `[3,9,20,null,null,15,7]`,

```
    3
   / \
  9  20
    /  \
   15   7
```
return its minimum depth = 2.

## 分析
求解二叉树的最小深度。获取一个节点的左右子树，取小的深度即可
- 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。（得理解这句话，测试用例第5个！！）
## 最高分
解法直观、便于理解
```golang
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left, right := minDepth(root.Left), minDepth(root.Right)
    if left == 0 || right == 0 {
        return left + right + 1
    }
    return min(left, right) + 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
```

## 实现
我的实现有点复杂，考虑使用共享内存，题目没太理解，
```golang
func minDepth(root *TreeNode) int {
    var depth int
    if root == nil {
        return 0
    }
    cal(root, &depth, 1)
    return depth
}

func cal(t *TreeNode, depth *int, count int) {
    // find the end of node
    if t.Left == nil && t.Right == nil {
        if *depth == 0 || count < *depth {
            *depth = count
        }
        return
    }
    if t.Left != nil {
        cal(t.Left, depth, count+1)
    }
    if t.Right != nil {
        cal(t.Right, depth, count+1)
    }
}
```

## 改进
```golang
func minDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left, right := minDepth(root.Left), minDepth(root.Right)
    if left == 0 || right == 0 {
        return left + right + 1
    }
    if left > right {
        return right + 1
    }
    return left + 1
}
```

## 反思

## 参考