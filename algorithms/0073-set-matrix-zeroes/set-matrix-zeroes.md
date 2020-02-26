## 问题
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:
```
Input: 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output: 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
```

Example 2:
```
Input: 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output: 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
```

Follow up:

- A straight forward solution using O(mn) space is probably a bad idea.
- A simple improvement uses O(m + n) space, but still not the best solution.
- Could you devise a constant space solution?

## 分析
判断每一个横竖行有没有0，有就更新，没有就不更新

## 最高分
```golang
func setZeroes(mat [][]int) {
	m, n := len(mat), len(mat[0])
	col0 := 1

	for i := 0; i < m; i++ {
		if mat[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

    for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
		if col0 == 0 {
			mat[i][0] = 0
		}
	}

}
```


## 实现
- 先存储第一行第一列是否有0
- 将其余列有0的行和列第一个元素赋值为0
- 扫描竖行和横行
- 将第一行和第一列赋值
```golang
func setZeroes(matrix [][]int) {
    height := len(matrix)
    if height == 0 {
        return
    }
    width := len(matrix[0])
    var h, w bool

    // get status of first row and first col
    for i := 0; i < width; i++ {
        if matrix[0][i] == 0 {
            w = true
        }
    }
    for i := 0; i < height; i++ {
        if matrix[i][0] == 0 {
            h = true
        }
    }

    // move
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }

    // set col
    for i := 1; i < width; i++ {
        if matrix[0][i] == 0 {
            for j := 1; j < height; j++ {
                matrix[j][i] = 0
            }
        }
    }

    // set row
    for i := 1; i < height; i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < width; j++ {
                matrix[i][j] = 0
            }
        }
    }

    // set first col
    if h {
        for i := 0; i < height; i++ {
            matrix[i][0] = 0
        }
    }

    // set first row
    if w {
        for i := 0; i < width; i++ {
            matrix[0][i] = 0
        }
    }

    return
}
```

## 改进
改进上面的算法，用了太多for循环
```golang

```

## 反思

## 参考