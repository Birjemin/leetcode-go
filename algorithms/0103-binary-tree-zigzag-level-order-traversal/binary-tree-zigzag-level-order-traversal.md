## 问题
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```

return its zigzag level order traversal as:
```
[
  [3],
  [20,9],
  [15,7]
]
```

## 分析
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
- 和102类似，可以使用bfs和dfs算法
## 最高分
```golang

```


## 实现
bfs算法

```golang
func zigzagLevelOrder(root *TreeNode) [][]int {
    var ret [][]int
    var b bool
    if root == nil {
        return [][]int{}
    }
    cal([]*TreeNode{root}, &ret, b)
    return ret
}

func cal(tmp []*TreeNode, ret *[][]int, b bool) {
    if len(tmp) == 0 {
        return
    }
    var t []*TreeNode
    var r []int
    var flag bool
    for _, v := range tmp {
        if b {
            r = append([]int{v.Val}, r...)
        } else {
            r = append(r, v.Val)
        }
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
    cal(t, ret, !b)
}
```

## 改进
dfs算法实现

```golang
func zigzagLevelOrder(root *TreeNode) [][]int {
    var ret [][]int
    if root == nil {
        return [][]int{}
    }
    var b bool
    cal(root, 0, b, &ret)
    return ret
}

func cal(tmp *TreeNode, depth int, b bool, ret *[][]int) {
    if tmp == nil {
        return
    }
    if len(*ret) <= depth {
        *ret = append(*ret, []int{})
    }

    if b {
        (*ret)[depth] = append([]int{tmp.Val}, (*ret)[depth]...)
    } else {
        (*ret)[depth] = append((*ret)[depth], tmp.Val)
    }

    if tmp.Left != nil {
        cal(tmp.Left, depth+1, !b, ret)
    }
    if tmp.Right != nil {
        cal(tmp.Right, depth+1, !b, ret)
    }
}

```

## 反思

## 参考