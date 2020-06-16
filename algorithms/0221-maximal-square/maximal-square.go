package maximal_square

func maximalSquare(matrix [][]byte) int {
	length := len(matrix)
	if length == 0 {
		return 0
	}
	width := len(matrix[0])
	if width == 0 {
		return 0
	}

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

func isSquare(matrix [][]byte, x, y, border int) bool {
	for i := x; i < border+x; i++ {
		for j := y; j < y+border; j++ {
			if matrix[i][j] != '1' {
				return false
			}
		}
	}
	return true
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

func maximalSquare1(matrix [][]byte) int {
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])
	if width == 0 {
		return 0
	}

	dp, ret := make([][]int, height), 0

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
