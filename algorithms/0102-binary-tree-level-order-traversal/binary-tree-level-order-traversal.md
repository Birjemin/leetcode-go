## 问题
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```

return its level order traversal as:
```
[
  [3],
  [9,20],
  [15,7]
]
```

## 分析
二叉树的层次遍历，可以使用bfs算法，bfs算法也可

## 最高分
```golang
func levelOrder1(root *TreeNode) [][]int {
    levels := [][]int{}
    dfsLevel(root, -1, &levels)
    return levels
}

func dfsLevel(node *TreeNode, level int, res *[][]int) {
    if node == nil {
        return
    }
    currLevel := level + 1
    for len(*res) <= currLevel {
        *res = append(*res, []int{})
    }
    (*res)[currLevel] = append((*res)[currLevel], node.Val)
    dfsLevel(node.Left, currLevel, res)
    dfsLevel(node.Right, currLevel, res)
}
```

## 实现
bfs算法即可

记录每次横向遍历的值和左右孩子树，然后在递归中循环孩子树即可
```golang
func levelOrder(root *TreeNode) [][]int {
    var ret [][]int
    if root == nil {
        return [][]int{}
    }
    cal([]*TreeNode{root}, &ret)
    return ret
}

func cal(tmp []*TreeNode, ret *[][]int) {
    if len(tmp) == 0 {
        return
    }
    var t []*TreeNode
    var r []int
    var flag bool
    for _, v := range tmp {
        r = append(r, v.Val)
        if v.Left != nil {
            t = append(t, v.Left)
            flag = true
        }
        if v.Right != nil {
            t = append(t, v.Right)
            flag = true
        }
    }
    *ret = append(*ret, r)
    if !flag {
        return
    }
    cal(t, ret)
}
```

## 改进
dfs算法

回溯的特性，记录每次深度
```golang
func levelOrder(root *TreeNode) [][]int {
    var ret [][]int
    if root == nil {
        return [][]int{}
    }
    cal(root, 0, &ret)
    return ret
}

func cal(tmp *TreeNode, depth int, ret *[][]int) {
    if tmp == nil {
        return
    }
    if len(*ret) <= depth {
        *ret = append(*ret, []int{})
    }

    (*ret)[depth] = append((*ret)[depth], tmp.Val)
    if tmp.Left != nil {
        cal(tmp.Left, depth+1, ret)
    }
    if tmp.Right != nil {
        cal(tmp.Right, depth+1, ret)
    }
}
```

## 反思

## 参考