## 问题
Given a binary tree, return the preorder traversal of its nodes' values.

Example:
```
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
```

Follow up: Recursive solution is trivial, could you do it iteratively?

## 分析
二叉树的前序遍历(不可以递归，使用迭代)
dfs（使用栈）
bfs（使用队列，略麻烦，需要记录每次的数据插入位置）

## 最高分
```golang
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    preorder := make([]int, 0)
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        head := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        preorder = append(preorder, head.Val)
        if head.Right != nil {
            stack = append(stack, head.Right)    
        }
        if head.Left != nil {
            stack = append(stack, head.Left)    
        }        
    }
    return preorder        
}
```

## 实现
```golang
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    var ret []int

    for stack, length := []*TreeNode{root}, 1; length != 0; {

        // drop one element
        root, stack = stack[length-1], stack[:length-1]
        length--

        ret = append(ret, root.Val)

        if root.Right != nil {
            stack = append(stack, root.Right)
            length++
        }
        if root.Left != nil {
            stack = append(stack, root.Left)
            length++
        }
    }
    return ret
}
```

## 改进
```golang

```

## 反思

## 参考