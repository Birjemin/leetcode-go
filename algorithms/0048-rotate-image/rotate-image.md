## 问题
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:
```
Given input matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
```

Example 2:
```
Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]```

## 分析
顺时针旋转90度，观看左边变化，可以知道就是每一列数组元素变成了竖着的状态
m(i,j) = m(n-j-1,i)

## 最高分
```golang
func rotate(matrix [][]int)  {
	n := len(matrix)
	for j := 0; j < n/2; j++ {
		for i := 0; i < n/2+n%2; i++ {
			matrix[j][i], matrix[i][n-1-j], matrix[n-1-j][n-1-i], matrix[n-1-i][j] =
				matrix[n-1-i][j], matrix[j][i], matrix[i][n-1-j], matrix[n-1-j][n-1-i]
		}
	}
}
```

## 实现
先得到每次循环的圈数，然后计算每个圈需要旋转的元素个数
```golang
func rotate(matrix [][]int) {
    length := len(matrix)
    // the number of circle
    for i := 0; i < length/2; i++ {
        // moving times of each circle
        for j := i; j < length-i-1; j++ {
            // temp
            temp := matrix[i][j]
            // one
            matrix[i][j] = matrix[length-j-1][i]
            // two
            matrix[length-j-1][i] = matrix[length-i-1][length-j-1]
            // three
            matrix[length-i-1][length-j-1] = matrix[j][length-i-1]
            // four
            matrix[j][length-i-1] = temp
        }
    }
}
```

## 改进
合并旋转
```golang
func rotate(matrix [][]int) {
    length := len(matrix)
    // the number of circle
    for i := 0; i < length/2; i++ {
        // moving times of each circle
        for j := i; j < length-i-1; j++ {
            matrix[i][j], matrix[length-j-1][i], matrix[length-i-1][length-j-1], matrix[j][length-i-1] = matrix[length-j-1][i], matrix[length-i-1][length-j-1], matrix[j][length-i-1], matrix[i][j]
        }
    }
}
```

## 反思

## 参考