## 问题
Given a binary tree, return the inorder traversal of its nodes' values.

Example:
```
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
```
Follow up: Recursive solution is trivial, could you do it iteratively?

## 分析
中序遍历

输入的数组怎么变成二叉树的？？很疑惑
数组第一项为根节点，第二项为左节点，如果为null，则第三项为有节点，反之第三项为第二项的左节点（左右移动）

```
关于测试：由于golang中int不能为null，所以测试用例中当为0时就认为是null（详情看测试文件）

```

## 最高分
```golang
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    result := inorderTraversal(root.Left)
    result = append(result, root.Val)
    return append(result, inorderTraversal(root.Right)...)
}
```

## 实现
左右子树计算就好了
```golang
func inorderTraversal(root *TreeNode) []int {
    var ret []int
    cal(root, &ret)
    return ret
}

func cal(t *TreeNode, ret *[]int) {
    if t == nil {
        return
    }
    // left child node
    cal(t.Left, ret)
    *ret = append(*ret, t.Val)
    // right child node
    cal(t.Right, ret)
}
```

## 改进
```golang

```

## 反思

## 参考