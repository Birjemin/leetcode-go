## 问题
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

* The left subtree of a node contains only nodes with keys less than the node's key.
* The right subtree of a node contains only nodes with keys greater than the node's key.
* Both the left and right subtrees must also be binary search trees.
 

Example 1:
```
    2
   / \
  1   3

Input: [2,1,3]
Output: true
```

Example 2:
```
    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

## 分析
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
没啥思路，看一下网友们的解答

- 中序遍历，然后判断每一项是否大于前一项即可

## 最高分
这个解法就很奇妙~~~
需要研究一下
```golang
func isValidBST(root *TreeNode) bool {
    return RecValidate(root, nil, nil)
}

func RecValidate(n, min, max *TreeNode) bool {
    if n == nil {
        return true
    }
    if min != nil && n.Val <= min.Val {
        return false
    }
    if max != nil && n.Val >= max.Val {
        return false
    }
    return RecValidate(n.Left, min, n) && RecValidate(n.Right, n, max)
}
```


## 实现
中序遍历的变种
```golang
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var ret []int
    return cal(root, &ret)
}

func cal(t *TreeNode, ret *[]int) bool {
    if t == nil {
        return true
    }
    // left child node
    if !cal(t.Left, ret) {
        return false
    }
    if *ret == nil {
        *ret = append(*ret, t.Val)
    } else if (*ret)[0] >= t.Val {
        return false
    } else {
        (*ret)[0] = t.Val
    }
    // right child node
    return cal(t.Right, ret)
}
```

## 改进
递归遍历
```golang
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return cal1(root, nil, nil)
}

func cal1(n, min, max *TreeNode) bool {
    if n == nil {
        return true
    }
    if min != nil && min.Val >= n.Val {
        return false
    }
    if max != nil && max.Val <= n.Val {
        return false
    }
    return cal1(n.Left, min, n) && cal1(n.Right, n, max)
}
```

## 反思

## 参考