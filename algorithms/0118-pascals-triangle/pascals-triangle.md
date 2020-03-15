## 问题
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
```
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
```
## 分析
打印杨辉三角~~
- dp算法

## 最高分
循环计算
```golang
func generate(numRows int) [][]int {
    rets := make([][]int, numRows)
    if numRows == 0 {
        return rets
    }
    rets[0] = []int{1}
    for i := 1; i < numRows; i++ {
        rets[i] = make([]int, i+1)
        rets[i][0] = 1
        rets[i][i] = 1
        for j := 1; j < i; j++ {
            rets[i][j] = rets[i-1][j-1] + rets[i-1][j]
        }
    }
    return rets
}
```


## 实现
dp算法，保存n-1的数据
```golang
func generate(numRows int) [][]int {
    if numRows == 0 {
        return [][]int{}
    }
    ret := make([][]int, numRows)
    dp(&ret, numRows)

    return ret
}

func dp(r *[][]int, n int) {
    if n == 1 {
        (*r)[0] = []int{1}
        return
    }
    // cal the collection of n-1
    dp(r, n-1)

    (*r)[n-1] = []int{1}
    for i := 0; i < n-2; i++ {
        (*r)[n-1] = append((*r)[n-1], (*r)[n-2][i] + (*r)[n-2][i+1])
    }
    (*r)[n-1] = append((*r)[n-1], 1)
}
```

## 改进
```golang

```

## 反思

## 参考