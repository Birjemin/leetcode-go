## 问题
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:
```
Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
```

## 分析
给定一个n，生成n*n的螺旋矩阵，很简单，和54题类似

## 最高分
使用了递归
```golang
func generateMatrix(n int) [][]int {
    m := make([][]int, n)
    for i := 0; i < n; i++ {
        m[i] = make([]int, n)
    }
    RecSpiral(1, 0, n-1, m)
    return m
}

func RecSpiral(current, start, end int, m [][]int) {
    if start == end {
        m[start][start] = current
        return
    }
    if start > end {
        return
    }
    for i := start; i <= end; i++ {
        m[start][i] = current
        current++
    }
    for i := start+1; i <= end; i++ {
        m[i][end] = current
        current++
    }
    for i := end-1; i >= start; i-- {
        m[end][i] = current
        current++
    }
    for i := end-1; i >= start+1; i-- {
        m[i][start] = current
        current++
    }
    RecSpiral(current, start+1, end-1, m)
}
```

## 实现
在54题上面更改即可
```golang
func generateMatrix(n int) [][]int {
    // init the array
    ret, v := make([][]int, n), 1
    for i := range ret {
        ret[i] = make([]int, n)
    }

    // the number of circle
    for i := 0; i < n; i++ {
        // moving times of each circle
        row := n - i
        // first
        for a := i; a < row; a++ {
            ret[i][a] = v
            v++
        }
        // second
        for b := i + 1; b < row; b++ {
            ret[b][row-1] = v
            v++
        }

        // three
        for c := row - 2; c > i; c-- {
            ret[row-1][c] = v
            v++
        }

        // four
        for d := row - 1; d > i; d-- {
            ret[d][i] = v
            v++
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