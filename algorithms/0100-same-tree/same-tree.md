## 问题
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:
```
Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
```
Example 2:
```
Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
```
Example 3:
```
Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
```

## 分析
比较简单

## 最高分
```golang
func isSameTree(p *TreeNode, q *TreeNode) bool {
    switch {
	case p == nil && q == nil:
		return true
	case p != nil && q == nil:
		return false
	case p == nil && q != nil:
		return false
	default:
		if p.Val != q.Val {
			return false
		} else {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
}
```

## 实现
直接使用递归实现
```golang
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p != nil && q != nil {
        if p.Val != q.Val {
            return false
        }
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    } else {
        return false
    }
}
```

## 改进
无
```golang

```

## 反思

## 参考