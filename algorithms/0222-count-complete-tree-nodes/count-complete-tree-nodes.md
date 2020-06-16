## 问题
Given a complete binary tree, count the number of nodes.

Note:

Definition of a complete binary tree from Wikipedia:

In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Example:
```
Input: 
    1
   / \
  2   3
 / \  /
4  5 6

Output: 6
```

## 分析
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2^h 个节点。

## 最高分
```golang

```

## 实现
前序遍历算法最简单
```golang
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	traverse(root,&ret)
	return ret
}

func traverse(root *TreeNode, ret *int) {
	if root == nil {
		return
	}
	*ret++
	traverse(root.Left, ret)
	traverse(root.Right, ret)
}
```

## 改进
```golang
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
```

## 反思

## 参考