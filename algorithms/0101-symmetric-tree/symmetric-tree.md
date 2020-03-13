## 问题
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

But the following [1,2,2,null,3,null,3] is not:
```
    1
   / \
  2   2
   \   \
   3    3
 
```
Note:
Bonus points if you could solve it both recursively and iteratively.

## 分析
对称树
- 将根节点的左支推入栈，右支出栈，进行判断（无法分辨[1, 2, 2, 2, null, 2]）

## 最高分
类似的原理
```golang
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true 
    }

    return recur(root.Left, root.Right)
}

func recur(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }

    if left == nil || right == nil {
        return false
    }

    return left.Val == right.Val && recur(left.Left, right.Right) && recur(left.Right, right.Left)
```


## 实现
想清楚之后会很简单，基本上就是同时两个指针走左子树，一个走右子树，然后判断两个大小
```golang
func isSymmetric(root *TreeNode) bool {
    // if root is nil
    if root == nil {
        return true
    }
    return traversal(root.Left, root.Right)
}

func traversal(l, r *TreeNode) bool {
    // end of traversal
    if l == nil && r == nil {
        return true
    }
    // the conditions of error
    if l == nil || r == nil || l.Val != r.Val || !traversal(l.Left, r.Right) {
        return false
    }
    return traversal(l.Right, r.Left)
}
```

## 改进
```golang

```

## 反思

## 参考