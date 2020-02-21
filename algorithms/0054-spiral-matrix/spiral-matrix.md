## 问题
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:
```
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
```
Example 2:
```
Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

## 分析
螺旋输入一个矩阵数组的元素，和48题类似的，也是找规律的题目

## 最高分
```golang
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var res []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	point := []int{0, 0}

	for {
		// move to right
		for i := point[1]; i <= right; i++ {
			res = append(res, matrix[point[0]][i])
		}
		if point[0] >= bottom {
			break
		}
		top, point[0], point[1] = top+1, point[0]+1, right

		// move to bottom
		for i := point[0]; i <= bottom; i++ {
			res = append(res, matrix[i][point[1]])
		}
		if point[1] <= left {
			break
		}
		right, point[0], point[1] = right-1, bottom, point[1]-1

		// move to left
		for i := point[1]; i >= left; i-- {
			res = append(res, matrix[point[0]][i])
		}
		if point[0] <= top {
			break
		}
		bottom, point[0], point[1] = bottom-1, point[0]-1, left

		// move to top
		for i := point[0]; i >= top; i-- {
			res = append(res, matrix[i][point[1]])
		}
		if point[1] >= right {
			break
		}
		left, point[0], point[1] = left+1, top, point[1]+1
	}
	return res
}
```

## 实现
找规律的题目，比如3x3的先求
- 第一行(0,0),(0,1),(0,2)
- 然后求(1,2),(2,2)
- 求(2,1)
- 求(2,0),(1,0)
- 求(1,1)
```golang
func spiralOrder(matrix [][]int) []int {
    col := len(matrix)
    if col == 0 {
        return []int{}
    }
    row := len(matrix[0])

    var ret []int

    var min int
    if col > row {
        min = row
    } else {
        min = col
    }
    // the number of circle
    i := 0
    for ; i < (min+1)/2; i++ {
        // moving times of each circle
        a := i
        for ; a < row-i; a++ {
            ret = append(ret, matrix[i][a])
        }
        a--
        b := i+1
        for ; b < col-i; b++ {
            ret = append(ret, matrix[b][a])
        }
        b--
        c := a-1
        if col-2*i == 1 || row-2*i == 1 {
            break
        }
        for ; c > i; c-- {
            ret = append(ret, matrix[b][c])
        }
        d := b
        for ; d > i; d-- {
            ret = append(ret, matrix[d][c])
        }
    }

    return ret
}

```

## 改进
调整代码，合并逻辑
```golang
func spiralOrder(matrix [][]int) []int {
    col := len(matrix)
    if col == 0 {
        return []int{}
    }
    row := len(matrix[0])
    min := (min(row, col) + 1) / 2
    var ret []int
    // the number of circle
    for i := 0; i < min; i++ {
        // moving times of each circle
        row1, col1 := row-i, col-i
        // first

        for a := i; a < row1; a++ {
            ret = append(ret, matrix[i][a])
        }
        // second
        for b := i + 1; b < col1; b++ {
            ret = append(ret, matrix[b][row1-1])
        }

        // break
        if col1-i == 1 || row1-i == 1 {
            break
        }

        // three
        for c := row1 - 2; c > i; c-- {
            ret = append(ret, matrix[col1-1][c])
        }

        // four
        for d := col1 - 1; d > i; d-- {
            ret = append(ret, matrix[d][i])
        }
    }
    return ret
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

```

## 反思

## 参考