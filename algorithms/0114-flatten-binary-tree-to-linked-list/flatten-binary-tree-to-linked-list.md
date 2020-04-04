## 问题
Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:
```
    1
   / \
  2   5
 / \   \
3   4   6
```

The flattened tree should look like:
```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

## 分析
给定一个二叉树，原地将它展开为链表。
- dfs遍历

## 最高分
```golang
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    _ = recFlatten(root)
}

func recFlatten(root *TreeNode) *TreeNode {
    right := root.Right
    if root.Left != nil {
        root.Right = root.Left
        root.Left = nil
        root = recFlatten(root.Right)
    }
    if right != nil {
        root.Right = right
        root = recFlatten(root.Right)
    }
    return root
}
```

## 实现
dfs遍历，得到子节点即可，然后分成4种情况进行拼接组装
```golang
func flatten(root *TreeNode) {
    dfs(root)
}

func dfs(curr *TreeNode) *TreeNode {
    if curr == nil {
        return nil
    }

    l1, l2 := dfs(curr.Left), dfs(curr.Right)

    switch {
    case l1 == nil && l2 == nil:
        return curr
    case l1 != nil && l2 != nil:
        curr.Right, l1.Right = curr.Left, curr.Right
        curr.Left = nil
        return l2
    case l1 != nil && l2 == nil:
        curr.Right, curr.Left = curr.Left, nil
        return l1
    default:
        return l2
    }
}
```

## 改进
```golang

```

## 反思

## 参考