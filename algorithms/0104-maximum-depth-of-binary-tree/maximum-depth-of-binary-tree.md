## 问题
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7

```
return its depth = 3.

## 分析
dfs遍历，计算二叉树深度

## 最高分
```golang
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```


## 实现
dfs算法遍历即可
```golang
func maxDepth(root *TreeNode) int {
    var depth int
    dfs(root, &depth, 0)
    return depth
}

func dfs(t *TreeNode, depth *int, count int) {
    if t == nil {
        if count > *depth {
            *depth = count
        }
        return
    }
    dfs(t.Left, depth, count+1)
    dfs(t.Right, depth, count+1)
}
```

## 改进
```golang

```

## 反思

## 参考