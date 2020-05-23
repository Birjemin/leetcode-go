## 问题
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]

![img](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

Example 1:
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
```

Example 2:
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
```

Note:

- All of the nodes' values will be unique.
- p and q are different and both values will exist in the BST.

## 分析
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

这道题目标注的是简易，但是想了很久没有思路。。。缺乏灵感，也可能是想复杂了。

PS:二叉搜索树！！是有序的！！！卧槽，我竟然忽略了这个特性。。。。

## 最高分
```golang
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	val1, val2 := p.Val, q.Val
	cur := root
	for {
		switch {
		case val1 < cur.Val && val2 < cur.Val:
			cur = cur.Left
		case val1 > cur.Val && val2 > cur.Val:
			cur = cur.Right
		default:
			return cur
		}
	}
	return root
}
```

## 实现
二叉搜索树！！！使用递归解决即可
```golang
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
```

## 改进
```golang

```

## 反思

## 参考