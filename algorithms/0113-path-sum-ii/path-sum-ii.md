## 问题
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,
```
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
```

Return:
```
[
   [5,4,11,2],
   [5,8,4,5]
]
```

## 分析
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

和上一题不同的是找出所有的路径
## 最高分
```golang
func pathSum(root *TreeNode, sum int) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    if root.Left == nil && root.Right == nil {
        if sum == root.Val {
            return append(res, []int{ root.Val })
        }
        return res
    }
    
    for _, path := range pathSum(root.Left, sum - root.Val) {
        res = append(res, append([]int{ root.Val}, path... ))
    }
    for _, path := range pathSum(root.Right, sum - root.Val) {
        res = append(res, append([]int{ root.Val}, path... ))
    }
    return res
}
```

## 实现
和112解法类似。
```golang
func pathSum(root *TreeNode, sum int) [][]int {
    var ret [][]int
    var tmp []int
    cal(root, sum, tmp, &ret)
    return ret
}

func cal(root *TreeNode, sum int, tmp []int, ret *[][]int) {
    if root == nil {
        return
    }
    sum -= root.Val
    tmp = append(tmp, root.Val)
    // find the child node
    if root.Left == nil && root.Right == nil {
        if sum == 0 {
            b := make([]int, len(tmp))
            copy(b, tmp)
            *ret = append(*ret, b)
        }
        return
    }
    // left node
    cal(root.Left, sum, tmp, ret)
    // right node
    cal(root.Right, sum, tmp, ret)
}
```

## 改进
```golang

```

## 反思

## 参考