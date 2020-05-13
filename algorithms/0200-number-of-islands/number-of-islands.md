## 问题
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
```
Input:
11110
11010
11000
00000

Output: 1
```

Example 2:
```
Input:
11000
11000
00100
00011

Output: 3
```

## 分析
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

这道题目和数独36题、79题类似的

dfs遍历

## 最高分
```golang
func dfs(grid *[][]byte, row, col int) {
	rows := len(*grid)
	cols := len((*grid)[0])

	if row < 0 || col < 0 || row >= rows || col >= cols || (*grid)[row][col] == '0' {
		return
	}

	(*grid)[row][col] = '0'
	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)

}

func numIslands(grid [][]byte) int {

	if grid == nil {
		return 0
	}

	if len(grid) == 0 {
		return 0
	}

	islands := 0
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				islands++
				dfs(&grid, row, col)
			}
		}
	}

	return islands
}
```

## 实现
```golang
func numIslands(grid [][]byte) int {

	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])

	// init malloc memory
	check := make([]bool, height*width)

	var ret int
	// check each of board
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '0' || check[i*width+j] {
				continue
			}
			ret++
			dfs(grid, check, width, height, i, j, 0)
		}
	}
	return ret
}

func dfs(tmp [][]byte, check []bool, width, height, i, j int, from int) {
	p := i*width + j

	// if the element is not equal to '0' or had already checked
	if check[p] || tmp[i][j] == '0' {
		return
	}

	// searching around of the element
	check[p] = true

	// top
	if i != 0 && from != 1 {
		dfs(tmp, check, width, height, i-1, j, 2)
	}

	// bottom
	if i != height-1 && from != 2 {
		dfs(tmp, check, width, height, i+1, j, 1)
	}

	// left
	if j != 0 && from != 3 {
		dfs(tmp, check, width, height, i, j-1, 4)
	}

	// right
	if j != width-1 && from != 4 {
		dfs(tmp, check, width, height, i, j+1, 3)
	}
}

```

## 改进
改进一下代码
```golang
func numIslands(grid [][]byte) int {

	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])

	var ret int
	// check each of board
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '0' || grid[i][j] == '2' {
				continue
			}
			ret++
			dfs(&grid, width, height, i, j)
		}
	}
	return ret
}

func dfs(tmp *[][]byte, width, height, i, j int) {

	if i < 0 || i > height-1 || j < 0 || j > width-1 {
		return
	}

	// if the element is not equal to '0' or had already checked
	if (*tmp)[i][j] == '2' || (*tmp)[i][j] == '0' {
		return
	}

	// searching around of the element
	(*tmp)[i][j] = '2'

	// top
	dfs(tmp, width, height, i-1, j)
	// bottom
	dfs(tmp, width, height, i+1, j)
	// left
	dfs(tmp, width, height, i, j-1)
	// right
	dfs(tmp, width, height, i, j+1)
}
```

## 反思

## 参考