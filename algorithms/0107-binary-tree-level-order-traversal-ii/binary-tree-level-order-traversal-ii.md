## 问题
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:

Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```

return its bottom-up level order traversal as:
```
[
  [15,7],
  [9,20],
  [3]
]
```

## 分析
和104类似，对不同深度的值进行汇总

## 最高分
有点奇妙，先遍dfs算法遍历处理每一个节点，直接保存到二维数组中（这里的二维数组是顺序的），然后转成逆序
```golang
func dfs(s *[][]int, level int, root *TreeNode)  {
    if root == nil {
        return
    }
    if len(*s) == level {
        *s = append(*s, []int{})
    }
    (*s)[level] = append((*s)[level], root.Val)
    for _, v := range []*TreeNode{root.Left, root.Right} {
        dfs(s, level + 1, v)
    }
}

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    var s [][]int
    dfs(&s, 0, root)
    for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return s
}
```


## 实现
在104的基础上改变，将每个层次的value保存下来即可，借助了map结构
```golang
func levelOrderBottom(root *TreeNode) [][]int {
    tmp := make(map[int][]int)
    dfs(root, &tmp, 0)
    return convertArrInt(tmp)
}

// convert map[int][]int to [][]int
func convertArrInt(tmp map[int][]int) [][]int {
    length := len(tmp)
    ret := make([][]int, length)
    for i, v := range tmp {
        ret[length-1-i] = v
    }
    return ret
}

// dsf
func dfs(t *TreeNode, tmp *map[int][]int, depth int) {
    if t == nil {
        return
    }
    // record the value
    (*tmp)[depth] = append((*tmp)[depth], t.Val)
    dfs(t.Left, tmp, depth+1)
    dfs(t.Right, tmp, depth+1)
}
```

## 改进
不借助map，和最高分类似
```golang

```

## 反思

## 参考