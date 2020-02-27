## 问题
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.

Example 1:
```
Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
```

Example 2:
```
Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
```

## 分析
在一个有序的矩阵中查找指定的target，存在就返回true，反之，返回false
- 找到所在行
- 使用二分法

## 最高分
简练的方式
```golang
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    row := sort.Search(len(matrix), func(i int) bool {
        return matrix[i][0] >= target
    })
    if row < len(matrix) && matrix[row][0] == target {
        return true
    }
    if row == 0 {
        return false
    }
    row = row - 1
    col := sort.Search(len(matrix[row]), func(i int) bool {
        return matrix[row][i] >= target
    })
    if col < len(matrix[row]) && matrix[row][col] == target {
        return true
    }
    return false
}
```

## 实现
```golang
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    width := len(matrix[0])
    if width == 0 {
        return false
    }
    var row int
    // find row
    for i, v := range matrix {
        if v[width-1] == target {
            return true
        } else if v[width-1] > target {
            row = i
            break
        }
    }

    // split
    left, mid := 0, width/2
    for ; mid != left; mid = (width-left)/2 + left {
        if matrix[row][mid] == target {
            return true
        } else if matrix[row][mid] > target {
            width = mid
        } else {
            left = mid
        }
    }
    if matrix[row][mid] == target {
        return true
    }
    return false
}
```

## 改进
```golang

```

## 反思

## 参考