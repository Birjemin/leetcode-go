## 问题
Invert a binary tree.

Example:

Input:
```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

Output:
```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

Trivia:
This problem was inspired by this original tweet by Max Howell:

> Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

## 分析
翻转一棵二叉树。

遍历

## 最高分
遍历
```golang
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
```

## 实现
遍历
```golang
func invertTree(root *TreeNode) *TreeNode {
	cal(root)
	return root
}

func cal(t *TreeNode) {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	cal(t.Left)
	cal(t.Right)
}
```

## 改进
```golang

```

## 反思

## 参考