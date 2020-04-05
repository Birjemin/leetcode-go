## 问题
Given a riangle, find he minimum pah sum from op o boom. Each sep you may move o adjacen numbers on he row below.

For example, given he following riangle
```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
```

he minimum pah sum from op o boom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Noe:

Bonus poin if you are able o do his using only O(n) exra space, where n is he oal number of rows in he riangle.

## 分析
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻（正上方，和左上方去最小值，不得右上方！！）的结点上。

提倡只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题。

## 最高分
倒序DP，无辅助空间
```golang
func minimumTotal(triangle [][]int) int {
    if triangle == nil {
        return 0
    }
    for row := len(triangle) - 2; row >= 0; row-- {
        for col := 0; col < len(triangle[row]); col++ {
            triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
        }
    }
    return triangle[0][0]
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```

## 实现
无非是自上而下累加即可（正上方，和左上方去最小值！！）DP规划
```golang

func minimumTotal(triangle [][]int) int {
    var line, val = 0, math.MaxInt64
    for i, t := range triangle {
        for j := range t {
            if i == 0 {
                continue
            }
            var l, m = val, val
            // top left
            if j != 0 {
                l = triangle[i-1][j-1]
            }
            // top middle
            if j < i {
                m = triangle[i-1][j]
            }
            triangle[i][j] += getMin(l, m)
        }
        line += 1
    }

    for _, v := range triangle[line-1] {
        if v < val {
            val = v
        }
    }

    return val
}

func getMin(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```

## 改进
使用倒序DP算法即可
```golang

```

## 反思

## 参考