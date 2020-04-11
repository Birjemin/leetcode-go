## 问题
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

Note: A leaf is a node with no children.

Example:
```
Input: [1,2,3]
    1
   / \
  2   3
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
```

Example 2:
```
Input: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.
```

## 分析
计算从根到叶子节点生成的所有数字之和。
dfs算法即可

## 最高分
```golang
func sumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(n *TreeNode, sum int) int {
    if n == nil {
        return 0
    }
        
    sum = sum*10 + n.Val
    
    if n.Left == nil && n.Right == nil {
        return sum
    }
    
    l := 0
    if n.Left != nil {
       l = dfs(n.Left, sum) 
    }
    
    r := 0
    if n.Right != nil {
        r = dfs(n.Right, sum)
    }
    
    return l+r
}
```

## 实现
最直接的方式
```golang
func sumNumbers(root *TreeNode) int {
    var ret int
    dfs(root, 0, &ret)
    return ret
}

func dfs(root *TreeNode, prev int, ret *int) {
    if root == nil {
        return
    }
    
    prev = prev*10 + root.Val
    
    if root.Left == nil && root.Right == nil {
        *ret += prev
        return
    }
    
    dfs(root.Left, prev, ret)
    dfs(root.Right, prev, ret)
}
```

## 改进
优化代码
```golang

```

## 反思

## 参考