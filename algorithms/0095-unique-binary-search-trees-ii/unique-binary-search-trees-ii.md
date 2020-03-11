## 问题
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

```
Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

## 分析
这道题目不会，需要仔细观看答案，加深理解~~

## 最高分
```golang

func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return recGenTrees(1, n)
}

func recGenTrees(s, e int) []*TreeNode {
    if e < s {
        return []*TreeNode{nil}
    }
    var result []*TreeNode
    for i := s; i <= e; i++ {
        lefts, rights := recGenTrees(s, i-1), recGenTrees(i+1, e)
        for _, l := range lefts {
            for _, r := range rights {
                result = append(result, &TreeNode{
                    Val:   i,
                    Left:  l,
                    Right: r,
                })
            }
        }
    }
    return result
}
```


## 实现
```golang

func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return recGenTrees(1, n)
}

func recGenTrees(s, e int) []*TreeNode {
    if e < s {
        return []*TreeNode{nil}
    }
    var result []*TreeNode
    for i := s; i <= e; i++ {
        lefts, rights := recGenTrees(s, i-1), recGenTrees(i+1, e)
        for _, l := range lefts {
            for _, r := range rights {
                result = append(result, &TreeNode{
                    Val:   i,
                    Left:  l,
                    Right: r,
                })
            }
        }
    }
    return result
}
```

## 改进
```golang

```

## 反思

## 参考