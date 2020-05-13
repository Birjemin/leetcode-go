package number_of_islands

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

func numIslands1(grid [][]byte) int {

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
			dfs1(&grid, width, height, i, j)
		}
	}
	return ret
}

func dfs1(tmp *[][]byte, width, height, i, j int) {

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
	dfs1(tmp, width, height, i-1, j)
	// bottom
	dfs1(tmp, width, height, i+1, j)
	// left
	dfs1(tmp, width, height, i, j-1)
	// right
	dfs1(tmp, width, height, i, j+1)
}
