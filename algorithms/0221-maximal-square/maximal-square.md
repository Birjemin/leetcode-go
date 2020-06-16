## 问题
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

For example, given the following matrix:
```
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
Return 4.
```

## 分析
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

## 最高分
```golang

```

## 实现
最简单的方式实现
```golang
func maximalSquare(matrix [][]byte) int {
	length, width := len(matrix), len(matrix[0])
	maxLen := min(length, width)

	for maxLen > 0 {
		for i := 0; i < length-maxLen+1; i++ {
			for j := 0; j < width-maxLen+1; j++ {
				if isSquare(matrix, i, j, maxLen) {
					return maxLen * maxLen
				}
			}
		}
		maxLen--
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func isSquare(matrix [][]byte, x, y, border int) bool {
	for i := x; i < border+x; i++ {
		for j := y; j < y+border; j++ {
			if matrix[i][j] != 1 {
				return false
			}
		}
	}
	return true
}
```

## 改进
根据官方题解，给出解法
```
         [ 0,  matrix[i][j]=0
dp[i][j]=|
         [ min{dp[i−1][j−1],  dp[i−1][j],  dp[i][j−1]}+1,  matrix[i][j]=1
```

```golang
func maximalSquare(matrix [][]byte) int {
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])
	if width == 0 {
		return 0
	}

	dp := make([][]int, height)
	var ret int
	for i := range dp {
		dp[i] = make([]int, width)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ret = 1
		}
	}

	for j := 1; j < width; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ret = 1
		}
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				ret = max(ret, dp[i][j])
			}
		}
	}
	return ret * ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

## 反思

## 参考